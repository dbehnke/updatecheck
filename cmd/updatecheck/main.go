package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/dbehnke74/updatecheck"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(`%s checks for software updates

results		  
  %s results [dir or single file (default: ./software)] [output file (default: results.json)]
  Scans all the software in yaml files and outputs to a results json file

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
			outputfile := "results.json"
			if len(os.Args) >= 4 {
				outputfile = os.Args[3]
			}
			results, err := updatecheck.GetSoftwareFilesResult(dir)
			if err != nil {
				fmt.Println("GetSoftwareFilesResult error: ", err)
				os.Exit(1)
			}
			jsonbytes, err := json.MarshalIndent(results, "", "  ")
			if err != nil {
				fmt.Println("JSON marshall error: ", err)
				os.Exit(1)
			}
			err = ioutil.WriteFile(outputfile, jsonbytes, 0644)
			if err != nil {
				fmt.Printf("Writing %s: %s\n", outputfile, err)
			}
			fmt.Println("Done.")
			os.Exit(0)
		}
	}
}
