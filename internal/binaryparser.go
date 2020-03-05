/*
 id: binaryparser.go
 This file is part of Florentino (C) 2020 0xsha
 me[at]0xsha.io
 @0xsha
 */
package internal

import (
	"debug/elf"
	"debug/macho"
	"debug/pe"
	"encoding/json"
	"fmt"
)


func ParsePE(filePath string, sha1sum string) string {


	peFile, err := pe.Open(filePath)
	Check(err)

	peLibs, _ := peFile.ImportedSymbols()
	peSymbols, _ := peFile.ImportedSymbols()


	// its possible to go way deeper, you know !?
	peHeader :=  peFile.FileHeader
	headers := make([]uint32, 3)
	headers[0] = uint32(peHeader.Machine)
	headers[1] = uint32(peHeader.Characteristics)
	headers[2] = peHeader.TimeDateStamp
	fmt.Printf("%x",headers)


	//fmt.Println(peLibs)
	//fmt.Println(peSymbols)

	report := report(peSymbols,peLibs,sha1sum)
	return report

}


func ParseELF(filePath string, sha1sum string)  string  {

	elfFile, err := elf.Open(filePath)
	Check(err)

	elfLibs, _ := elfFile.ImportedLibraries()
	elfSymbols, _ := elfFile.ImportedSymbols()
	elfHeader :=  elfFile.FileHeader


	fmt.Println(elfHeader)
	fmt.Println(elfLibs)
	fmt.Println(elfSymbols)

	var symbols  []string
	for _,symbol:= range elfSymbols{

		symbols = append(symbols,symbol.Name+" "+symbol.Library+" "+symbol.Version)
	}

	report := report(symbols,elfLibs,sha1sum)
	return report

}


func ParseMaco(filePath string, sha1sum string) string  {


	macoFile, err := macho.Open(filePath)
	Check(err)

	macoLibs,_ := macoFile.ImportedLibraries()
	macoSymbols,_ := macoFile.ImportedSymbols()

	fmt.Println(macoSymbols)
	fmt.Println(macoLibs)

	report := report(macoSymbols,macoLibs,sha1sum)
	return report

}


func report(symbols []string, libs []string ,sha1sum string) string {

	type report struct {

		 Symbols  []string `json:"symbols"`
		 Libraries  []string `json:"libraries"`
		 Sha1sum string   `json:"sha1sum"`
	}

	rep := report{Symbols: symbols , Libraries:libs , Sha1sum:sha1sum}

	var jsonData []byte
	jsonData, err := json.Marshal(rep)
	Check(err)
	//fmt.Println(string(jsonData))

	return string(jsonData)
}