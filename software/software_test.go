package software

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCoreFunctions(t *testing.T) {
	Convey("GetBody can fetch response body of google.com", t, func() {
		body, err := GetBody("https://google.com")
		So(err, ShouldBeNil)
		So(body, ShouldNotBeEmpty)
	})
	Convey("GetRequestURL of https://www.google.com works on non-redirects", t, func() {
		result, err := GetRequestURL("https://www.google.com")
		So(err, ShouldBeNil)
		So(result, ShouldNotBeEmpty)
		So(result, ShouldEqual, "https://www.google.com")
	})
	Convey("GetRequestURL of https://google.com works on redirects", t, func() {
		result, err := GetRequestURL("https://google.com")
		So(err, ShouldBeNil)
		So(result, ShouldNotBeEmpty)
		So(result, ShouldEqual, "https://www.google.com/")
	})
	Convey("Matches works with regular expressions", t, func() {
		values := Matches("test", "test")
		So(values, ShouldNotBeNil)
		So(len(values), ShouldEqual, 1)

		values = Matches("skjfksj>>http://happy.dmg", "http://.{1,}\\.dmg")
		So(values, ShouldNotBeNil)
		So(len(values), ShouldEqual, 1)
		So(values[0], ShouldEqual, "http://happy.dmg")
	})
}

func TestYAML(t *testing.T) {
	Convey("YAML can be marshaled from a file", t, func() {
		software, err := SoftwareFromYAMLFile("../definitions/mozilla.yml")
		So(err, ShouldBeNil)
		So(software, ShouldNotBeEmpty)
	})
}

func TestGetRuleResults(t *testing.T) {
	Convey("Rule Results are working", t, func() {
		software, err := SoftwareFromYAMLFile("../definitions/mozilla.yml")
		So(err, ShouldBeNil)
		So(software, ShouldNotBeEmpty)
		s := software[0]
		So(s.Name, ShouldNotBeEmpty)
		So(s.DisplayName, ShouldNotBeEmpty)
		So(s.Description, ShouldNotBeEmpty)
		So(s.Platforms, ShouldNotBeEmpty)
		p := s.Platforms[0]
		So(p.Rules, ShouldNotBeEmpty)
		results, err := GetRuleResults(p.Rules)
		So(err, ShouldBeNil)
		So(results, ShouldNotBeEmpty)
	})
}

func TestGetPlatformResults(t *testing.T) {
	Convey("Rule Results are working", t, func() {
		software, err := SoftwareFromYAMLFile("../definitions/mozilla.yml")
		So(err, ShouldBeNil)
		So(software, ShouldNotBeEmpty)
		s := software[0]
		So(s.Name, ShouldNotBeEmpty)
		So(s.DisplayName, ShouldNotBeEmpty)
		So(s.Description, ShouldNotBeEmpty)
		So(s.Platforms, ShouldNotBeEmpty)
		So(len(s.Platforms), ShouldBeGreaterThanOrEqualTo, 2)
		So(s.Platforms[0], ShouldNotBeEmpty)
		So(s.Platforms[1], ShouldNotBeEmpty)
		So(s.Platforms[0].Name, ShouldEqual, "mac")
		So(s.Platforms[1].Name, ShouldEqual, "windows")
		results, err := GetPlatformResults(s.Platforms)
		So(err, ShouldBeNil)
		So(results, ShouldNotBeEmpty)
	})
}

func TestGetSoftwareResults(t *testing.T) {
	Convey("Rule Results are working", t, func() {
		software, err := SoftwareFromYAMLFile("../definitions/mozilla.yml")
		So(err, ShouldBeNil)
		So(software, ShouldNotBeEmpty)
		results, err := GetSoftwareResults(software)
		So(err, ShouldBeNil)
		So(results, ShouldNotBeEmpty)
		So(results["firefoxesr"], ShouldNotBeEmpty)
		So(results["firefoxesr"]["mac"], ShouldNotBeEmpty)
		So(results["firefoxesr"]["windows"], ShouldNotBeEmpty)
	})
}

func TestGetSoftwareFiles(t *testing.T) {
	Convey("GetSoftwareFiles can get a list of files from a directory", t, func() {
		files, err := GetSoftwareFiles("../definitions")
		So(err, ShouldBeNil)
		So(files, ShouldNotBeEmpty)
	})

	Convey("GetSoftwareFiles can get a single file", t, func() {
		files, err := GetSoftwareFiles("../definitions/mozilla.yml")
		So(err, ShouldBeNil)
		So(files, ShouldNotBeEmpty)
		So(files[0], ShouldEqual, "mozilla.yml")
	})
}

func TestGetSoftwareFilesResult(t *testing.T) {
	Convey("GetSoftwareFilesResult can get results from a directory of files", t, func() {
		results, err := GetSoftwareFilesResult("../definitions")
		So(err, ShouldBeNil)
		So(results, ShouldNotBeEmpty)
		So(results, ShouldContainKey, "firefoxesr")
	})
	Convey("GetSoftwareFilesResult can get results from a single file", t, func() {
		results, err := GetSoftwareFilesResult("../definitions/mozilla.yml")
		So(err, ShouldBeNil)
		So(results, ShouldNotBeEmpty)
		So(results, ShouldContainKey, "firefoxesr")
	})
}

func TestStaticStringAndRegExpReplace(t *testing.T) {
	Convey("Set Static String and Do a RegExpReplace", t, func() {
		results, err := GetSoftwareFilesResult("../definitions/adobe.yml")
		So(err, ShouldBeNil)
		So(results, ShouldNotBeEmpty)
		So(results, ShouldContainKey, "adobeflashplayer")
		flash := results["adobeflashplayer"]
		So(flash, ShouldContainKey, "mac")
		mac := flash["mac"]
		So(mac, ShouldContainKey, "current_version")
	})
}
