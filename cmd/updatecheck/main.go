package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	logxi "github.com/mgutz/logxi/v1"
	"gitlab.com/dbehnke74/updatecheck"
	yaml "gopkg.in/yaml.v2"
)

//2017.07.27
const versionjre = "1.8.0.144"

//2017.07.11
const versionfirefox = "52.3.0"

//2017.08.08
const versionflash = "26.0.0.151"

//2017.08.16
const versionoffice2016 = "15.37.0 [170815] 64-bit"

func olddetect() {
	log := logxi.New("olddetect")
	var check string
	var exitcode int
	if len(os.Args) > 1 {
		check = strings.ToLower(os.Args[1])
		log.Debug("check", "check", check)
	}
	checkall := strings.HasPrefix(check, "all")
	checkjre := strings.Contains(check, "jre") || checkall
	if checkjre {
		u, err := updatecheck.OracleJRE()
		if err != nil {
			log.Error("Error Checking OracleJRE", "err", err)
			exitcode = 1
		} else {
			fmt.Println(u.Product, u.CurrentVersion, versionjre)
			if u.CurrentVersion != versionjre {
				log.Warn("Update Detected", "current", u.CurrentVersion,
					"check", versionjre)
				exitcode = 2
			}
		}
	}

	checkflash := strings.Contains(check, "flash") || checkall
	if checkflash {
		u, err := updatecheck.AdobeFlashPlayer()
		if err != nil {
			log.Error("Error Checking AdobeFlashPlayer", "err", err)
			exitcode = 1
		} else {
			fmt.Println(u.Product, u.CurrentVersion, versionflash)
			if u.CurrentVersion != versionflash {
				log.Warn("Update Detected", "current", u.CurrentVersion,
					"check", versionflash)
				exitcode = 2
			}
		}
	}

	checkfirefox := strings.Contains(check, "firefox") || checkall
	if checkfirefox {
		s := updatecheck.NewFirefoxESR("mac")
		if !s.Refresh() {
			log.Error("Error Checking FirefoxESR")
			exitcode = 1
		} else {
			fmt.Println(s.Name, s.CurrentVersion(), versionfirefox)
			if s.CurrentVersion() != versionfirefox {
				log.Warn("Update Detected", "current", s.CurrentVersion(),
					"check", versionfirefox)
				exitcode = 2
			}
		}
	}

	checkoffice2016 := strings.Contains(check, "office2016") || checkall
	if checkoffice2016 {
		u, err := updatecheck.Office2016()
		if err != nil {
			log.Error("Error Checking Office2016", "err", err)
			exitcode = 1
		} else {
			fmt.Println(u.Product, u.CurrentVersion, versionoffice2016)
			if u.CurrentVersion != versionoffice2016 {
				log.Warn("Update Detected", "current", u.CurrentVersion,
					"check", versionoffice2016)
				exitcode = 2
			}
		}
	}
	os.Exit(exitcode)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(`%s checks for software updates

results		  
  %s results [dir or single file (default: ./software)] [output file (default: lastresults.yml)]
  Scans all the software in yaml files and outputs to a results yaml file

`, os.Args[0], os.Args[0])
		os.Exit(0)
	}
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "results":
			dir := "./software"
			if len(os.Args) >= 3 {
				dir = os.Args[2]
			}
			outputfile := "lastresults.yml"
			if len(os.Args) >= 4 {
				outputfile = os.Args[3]
			}
			results, err := updatecheck.GetSoftwareFilesResult(dir)
			if err != nil {
				fmt.Println("GetSoftwareFilesResult error: ", err)
				os.Exit(1)
			}
			yamlbytes, err := yaml.Marshal(results)
			if err != nil {
				fmt.Println("YAML marshall error: ", err)
				os.Exit(1)
			}
			err = ioutil.WriteFile(outputfile, yamlbytes, 0644)
			if err != nil {
				fmt.Printf("Writing %s: %s\n", outputfile, err)
			}
			fmt.Println("Done.")
			os.Exit(0)
		}
	}
}
