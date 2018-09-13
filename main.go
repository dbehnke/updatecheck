package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/dbehnke74/updatecheck/software"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(`%s checks for software updates

results		  
  %s results [dir or single file (default: ./definitions)] [output file (default: results.json)]
  Scans all the software in yaml files and outputs to a results json file

query
  %s query [file.json] string1 <string2> <string3>
  i.e. %s query results.json thunderbird mac download_url

`, os.Args[0], os.Args[0], os.Args[0], os.Args[0])
	}
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "results":
			dir := "./definitions"
			if len(os.Args) >= 3 {
				dir = os.Args[2]
			}
			outputfile := "results.json"
			if len(os.Args) >= 4 {
				outputfile = os.Args[3]
			}
			results, err := software.GetSoftwareFilesResult(dir)
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

		case "query":
			resultfile := ""

			if len(os.Args) >= 3 {
				resultfile = os.Args[2]
			} else {
				fmt.Println("check parameters")
				os.Exit(1)
			}
			b, err := ioutil.ReadFile(resultfile)
			if err != nil {
				fmt.Println("error reading file", err)
			}
			var m map[string]interface{}
			err = json.Unmarshal(b, &m)
			if err != nil {
				fmt.Println("error couldn't unmarshal: ", err)
			}
			//fmt.Println(m)
			if len(os.Args) == 4 {
				k := os.Args[3]
				v, ok := m[k]
				if !ok {
					fmt.Println("error key not found:", k)
					os.Exit(1)
				}
				fmt.Println(v)
			}
			if len(os.Args) == 5 {
				k := os.Args[3]
				kk := os.Args[4]
				v, ok := m[k].(map[string]interface{})[kk]
				if !ok {
					fmt.Println("error key not found:", k, kk)
					os.Exit(1)
				}
				fmt.Println(v)
			}
			if len(os.Args) == 6 {
				k := os.Args[3]
				kk := os.Args[4]
				kkk := os.Args[5]
				v, ok := m[k].(map[string]interface{})[kk].(map[string]interface{})[kkk]
				if !ok {
					fmt.Println("error key not found:", k, kk, kkk)
					os.Exit(1)
				}
				fmt.Println(v)
			}
		}
	}
}
