/*
 id: detectengine.go
 This file is part of Florentino (C) 2020 0xsha
 me[at]0xsha.io
 @0xsha
*/
package internal

import (
	"encoding/json"
	"fmt"
	"github.com/VirusTotal/vt-go"
	"github.com/joho/godotenv"
	"io/ioutil"
	"strings"
)

type DetectionEngine struct {
	FileType string
	FileName string
	DetectedType string
	Entropy string
}

func FileAnalysis(detectedFile DetectionEngine, file string, sha1sum string)  {

	if strings.HasPrefix(detectedFile.FileType,"PE") {
		fmt.Println("PE File detected")
		fileReport := ParsePE(file,sha1sum)
		CreateTextFile("data/file_"+sha1sum+".txt",fileReport)

		fmt.Println("[+] FileName:"+ detectedFile.FileName)
		fmt.Println("[+] DetectedType:" +detectedFile.DetectedType)
	}
	if strings.HasPrefix(detectedFile.FileType,"MACH") {
		fmt.Println("MAC-O File detected")
		ParseMaco(file,sha1sum)

		fileReport := ParseMaco(file,sha1sum)
		CreateTextFile("data/file_"+sha1sum+".txt",fileReport)

		fmt.Println("[+] FileName:"+ detectedFile.FileName)
		fmt.Println("[+] DetectedType:" +detectedFile.DetectedType)
	}
	if strings.HasPrefix(detectedFile.FileType,"ELF") {
		fmt.Println("ELF File detected")
		ParseELF(file,sha1sum)

		fileReport := ParseELF(file,sha1sum)
		CreateTextFile("data/file_"+sha1sum+".txt",fileReport)

		fmt.Println("[+] FileName:"+ detectedFile.FileName)
		fmt.Println("[+] DetectedType:" +detectedFile.DetectedType)
	}


}
func DetectFile(fileHash string) DetectionEngine {

	diecJson, err := ioutil.ReadFile("data/dic_"+fileHash+".txt")


	// Json a little different in go
	type DecodedResults struct {
		Detects  []map[string]string
		Entropy string
		Filename string
	}

	var decoded DecodedResults
	err = json.Unmarshal(diecJson, &decoded)
	Check(err)

	// Binary detection and parsing
	detectedResult := DetectionEngine{decoded.Detects[0]["filetype"], decoded.Detects[0]["name"],
		decoded.Detects[0]["type"] , decoded.Entropy }

	return detectedResult
}

func CheckVirusTotal(sha256 string)  string {


	err := godotenv.Load()
	Check(err)
	env , err := godotenv.Read()
	Check(err)

	client := vt.NewClient(env["VIRUSTOTAL_API"])
	file, err := client.GetObject(vt.URL("files/%s", sha256))
	Check(err)
	ls, err := file.GetTime("last_submission_date")
	if err != nil {
		return "NeverScanned"
	}
	fmt.Printf("File %s was submitted for the last time on %v\n", file.ID(), ls)

	return file.ID()

}