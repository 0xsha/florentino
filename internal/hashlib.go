/*
 id: hashlib.go
 This file is part of Florentino (C) 2020 0xsha
 me[at]0xsha.io
 @0xsha
*/

package internal

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"io/ioutil"

)

func FileCheckSums(filePath string) ( [16]byte , [20]byte, [32]byte){

	fileBytes, _ := ioutil.ReadFile(filePath)

	md5Sum := md5.Sum(fileBytes)
	sha1Sum :=  sha1.Sum(fileBytes)
	sha256Sum := sha256.Sum256(fileBytes)


	return   md5Sum ,sha1Sum, sha256Sum
}