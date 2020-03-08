/*
 id: main.go
 This file is part of Florentino (C) 2020 0xsha
 me[at]0xsha.io
 @0xsha
*/

package main

import (
	"encoding/hex"
	"fmt"
	fl "github.com/0xsha/florentino/internal"
	"github.com/akamensky/argparse"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	// Load and Read .env
	err := godotenv.Load()
	fl.Check(err)
	env , err := godotenv.Read()
	fl.Check(err)


	// parse user arguments
	parser := argparse.NewParser("florentino", "Florentino")
	fileArg := parser.String("f", "file", &argparse.Options{Required: true, Help: "File to start."})
	err = parser.Parse(os.Args)
	fl.Check(err)

    // Calculate file checksums
	md5,sha1,sha256 := fl.FileCheckSums(*fileArg)
	//md5Encoded :=  hex.EncodeToString(md5[:])
	sha1Encoded :=  hex.EncodeToString(sha1[:])
	sha256Encoded :=  hex.EncodeToString(sha256[:])
	fmt.Printf("MD5 : %x\nSHA1 %x\nSHA256 %x\n", md5,sha1,sha256)


	//start detection engine (diec)
	fmt.Println(color.GreenString("[+] File type extraction ..."))
	fmt.Println("Flarentino : " + color.CyanString(fl.FlorentinoQuotes() ))
	fl.AnalysisExists("diec",sha1Encoded)
	diecResult :=   fl.SafeExec(env["DIEC_PATH"], *fileArg, "-fullscan:yes"  ,"-showjson:yes", "-showentropy:yes")
	fl.CreateTextFile("data/dic_"+sha1Encoded+".txt",diecResult)

	// VirusTotal scanning
	isScanned := fl.CheckVirusTotal(sha256Encoded)
	if isScanned == "NeverScanned" {
		fmt.Println("Don't forget to scan it")
	}
	if isScanned == sha256Encoded {
		fmt.Println(color.GreenString("[+] File already scanned !"))
	}

	// file detection and binary parsing
	detectedFile := fl.DetectFile(sha1Encoded)
	fmt.Println(detectedFile.DetectedType)
	fl.FileAnalysis(detectedFile,*fileArg,sha1Encoded)
	fmt.Println("Flarentino : " + color.CyanString(fl.FlorentinoQuotes() ))


	// x86 packer detection and unpacking
	entropy := detectedFile.Entropy
	if (fl.IsPacked(entropy)) && (fl.IsPEX86(*fileArg)){
		fl.UnpackPE(*fileArg)
	}





	// strings part
	fmt.Println(color.GreenString("[+] String extraction ..."))
	fmt.Println("Flarentino : " + color.CyanString(fl.FlorentinoQuotes() ))
	fl.AnalysisExists("strings",sha1Encoded)
	stringsResult :=   fl.SafeExec("strings", *fileArg)
	fl.CreateTextFile("data/strings_"+sha1Encoded+".txt",stringsResult)



	// Floss part
	fmt.Println(color.GreenString("[+] Floss extraction ..."))
	fmt.Println("Flarentino : " + color.CyanString(fl.FlorentinoQuotes() ))
	fl.AnalysisExists("floss",sha1Encoded)
	flossResult :=   fl.SafeExec(env["FLOSS_PATH"], *fileArg, "-q"  ,"-g")
	fl.CreateTextFile("data/floss_"+sha1Encoded+".txt",flossResult)


	// IOCExtract
	fmt.Println(color.GreenString("[+] IOC extraction ..."))
	fmt.Println("Flarentino : " + color.CyanString(fl.FlorentinoQuotes() ))
	fl.SafeExec("iocextract", "--input", "data/strings_"+sha1Encoded+".txt" ,
									"--output" ,"data/ioc_"+sha1Encoded+".txt")
	fl.SafeExec("iocextract", "--input", "data/floss_"+sha1Encoded+".txt" ,
									"--output" ,"data/ioc_floss"+sha1Encoded+".txt")

	fmt.Println("[+] Everything done ! Good Hunting ")

}
