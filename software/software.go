package software

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	logxi "github.com/mgutz/logxi/v1"
	yaml "gopkg.in/yaml.v2"
)

//Rule describes how to obtain a result
type Rule struct {
	Name       string            `yaml:"name"`
	Method     string            `yaml:"method"`
	Parameters map[string]string `yaml:"parameters"`
}

//Platform is used to seperate rules.
//A platform is something like mac, windows, linux etc.
type Platform struct {
	Name  string `yaml:"name"`
	Rules []Rule `yaml:"rules"`
}

//Software defines the framework for update checking
type Software struct {
	Name        string     `yaml:"name"`
	DisplayName string     `yaml:"displayname"`
	Description string     `yaml:"description"`
	Platforms   []Platform `yaml:"platforms"`
}

//YamlFile represents the structure imported from Yaml
type YamlFile struct {
	Version  int        `yaml:"version"`
	Software []Software `yaml:"software"`
}

//SoftwareFromYAMLBytes returns Software from byte array
func SoftwareFromYAMLBytes(in []byte) ([]Software, error) {
	y := YamlFile{}
	err := yaml.Unmarshal(in, &y)
	if err != nil {
		return nil, err
	}
	return y.Software, nil
}

//SoftwareFromYAMLFile returns Software from a yaml file
func SoftwareFromYAMLFile(filename string) ([]Software, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	s, err := SoftwareFromYAMLBytes(b)
	return s, err
}

//GetRuleResults runs through rules and populates a map with results
func GetRuleResults(rules []Rule) (map[string]string, error) {
	results := map[string]string{}
	for _, rule := range rules {
		switch rule.Method {
		case "scrape":
			url, ok := rule.Parameters["url"]
			if !ok {
				return nil, errors.New("url_redirect missing url parameter")
			}
			body, err := GetBody(url)
			if err != nil {
				return nil, err
			}
			results[rule.Name] = body
		case "url_redirect":
			url, ok := rule.Parameters["url"]
			if !ok {
				return nil, errors.New("url_redirect missing url parameter")
			}
			result, err := GetRequestURL(url)
			if err != nil {
				return nil, errors.New("url_redirect unable to get request url: " + err.Error())
			}

			results[rule.Name] = result
		case "rule_match":
			r, ok := rule.Parameters["rule"]
			if !ok {
				return nil, fmt.Errorf("rule_match missing rule")
			}
			v, ok := results[r]
			if !ok {
				return nil, fmt.Errorf("rule_match %s result not found. Run the rule to define %s first", r, r)
			}
			m, ok := rule.Parameters["match"]
			if !ok {
				return nil, errors.New("rule_match match parameter not found")
			}
			s := Matches(v, m)
			if s == nil {
				return nil, fmt.Errorf("rule_match didn't match: %s (%s)", v, m)
			}
			results[rule.Name] = s[0]
		case "static_string":
			value, ok := rule.Parameters["value"]
			if !ok {
				return nil, errors.New("static_string value parameter not found")
			}
			results[rule.Name] = value
		case "rule_regexp_replace":
			r, ok := rule.Parameters["rule"]
			if !ok {
				return nil, errors.New("rule_regexp_replace rule parameter not found")
			}
			src, ok := results[r]
			if !ok {
				return nil, errors.New("rule_regexp_replace could not find rule defined in paraemter")
			}
			match, ok := rule.Parameters["match"]
			if !ok {
				return nil, errors.New("rule_regexp_replace match parameter not found")
			}
			repl, ok := rule.Parameters["repl"]
			if !ok {
				return nil, errors.New("rule_regexp_replace repl parameter not found")
			}
			re := regexp.MustCompile(match)
			results[rule.Name] = re.ReplaceAllString(src, repl)
		case "rule_replace_rule":
			r, ok := rule.Parameters["rule"]
			if !ok {
				return nil, errors.New("rule_replace_rule rule parameter not found")
			}
			src, ok := results[r]
			if !ok {
				return nil, errors.New("rule_replace_rule could not find rule defined in paraemter")
			}
			match, ok := rule.Parameters["match"]
			if !ok {
				return nil, errors.New("rule_replace_rule match parameter not found")
			}
			replRule, ok := rule.Parameters["repl_rule"]
			if !ok {
				return nil, errors.New("rule_replace_rule repl_rule parameter not found")
			}
			repl, ok := results[replRule]
			if !ok {
				return nil, errors.New("rule_replace_rule repl_rule not found in cache")
			}
			re := regexp.MustCompile(match)
			results[rule.Name] = re.ReplaceAllString(src, repl)
		}

	}
	for k := range results {
		if strings.HasPrefix(k, "_") {
			delete(results, k)
		}
	}
	return results, nil
}

//GetPlatformResults runs through Platforms, Processes rules,
//and populates a map with results
func GetPlatformResults(platforms []Platform) (map[string]map[string]string, error) {
	results := map[string]map[string]string{}
	for _, platform := range platforms {
		r, err := GetRuleResults(platform.Rules)
		if err != nil {
			return nil, err
		}
		results[platform.Name] = r
	}
	return results, nil
}

//GetSoftwareResults runs through Software, Platforms, Processes rules,
//and populates a map with results
func GetSoftwareResults(software []Software) (map[string]map[string]map[string]string, error) {
	results := map[string]map[string]map[string]string{}
	for _, s := range software {
		r, err := GetPlatformResults(s.Platforms)
		if err != nil {
			return nil, err
		}
		results[s.Name] = r
	}
	return results, nil
}

//GetRequestURL will attempt to GET from url and return the final request url
//This is useful to determine the final url in a series of redirects.
func GetRequestURL(url string) (string, error) {
	log := logxi.New("GetRequestURL")
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Error("retreiving url", "url", url, "err", err)
		return "", err
	}
	if resp.StatusCode != 200 {
		log.Warn("Expecting a 200", "url", url, "StatusCode",
			resp.StatusCode)
		err = errors.New("Expecting a 200")
		return "", err
	}
	rurl := resp.Request.URL.String()

	log.Debug("", "rurl", rurl)
	log.Debug("", "url", url)
	return rurl, nil
}

//GetBody retrieves the response.body of a url
func GetBody(url string) (string, error) {
	log := logxi.New("GetBody")
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Error("retreiving url", "url", url, "err", err)
		return "", err
	}
	if resp.StatusCode != 200 {
		log.Warn("Expecting a 200", "url", url, "StatusCode",
			resp.StatusCode)
		err = errors.New("Expecting a 200")
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn("Reading Body", "url", url, "err", err)
		return "", err
	}
	log.Debug("body contents", "body", string(body))
	return string(body), nil
}

//Matches returns matches of a string
func Matches(s string, matchexp string) []string {
	re := regexp.MustCompile(matchexp)
	return re.FindAllString(s, -1)
}

//GetSoftwareFiles scans a directory dir for yaml files
//If dir is a file it returns only that file.
//nested directories are not traversed.
func GetSoftwareFiles(dir string) ([]string, error) {
	finfo, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if finfo.IsDir() == false {
		return []string{finfo.Name()}, nil
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	result := []string{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if strings.HasSuffix(f.Name(), ".yml") == true {
			result = append(result, f.Name())
		}
	}
	return result, err
}

//GetSoftwareFilesResult parses yaml files in a directory and fills
//a map with results. The key of the map will be the short name of the
//software.
//if dir is a file, it will be processed as a single file instead of a directory.
//nested directories are not traversed.
func GetSoftwareFilesResult(dir string) (map[string]map[string]map[string]string, error) {
	result := map[string]map[string]map[string]string{}
	files, err := GetSoftwareFiles(dir)
	if err != nil {
		return nil, err
	}
	finfo, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	dirIsFile := false
	if finfo.IsDir() == false {
		dirIsFile = true
	}
	for _, f := range files {
		if dirIsFile {
			f = dir
		} else {
			f = dir + string(os.PathSeparator) + f
		}
		s, err := SoftwareFromYAMLFile(f)
		if err != nil {
			return nil, err
		}
		r, err := GetSoftwareResults(s)
		if err != nil {
			return nil, err
		}
		for k, v := range r {
			result[k] = v
		}
	}
	return result, err
}
