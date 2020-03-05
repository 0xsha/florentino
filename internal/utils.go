/*
 id: utils.go
 This file is part of Florentino (C) 2020 0xsha
 me[at]0xsha.io
 @0xsha
*/

package internal

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func AnalysisExists(command string, checksum string) {
	if FileExists(command+"_"+checksum+".txt") {
		fmt.Println("Analyze for this file already exists")
		os.Exit(-1)
	}
}

func CreateTextFile(filePath string , dataString string) bool  {

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = file.WriteString(dataString)
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}

	return true
}

// fun function !
func FlorentinoQuotes()   string {

	var quotes [16]string
	quotes[0] = "Fairest ladies my lips are like whatever I'll finish this later ...."
	quotes[1] = "Flowers, women – I desire all that is beautiful."
	quotes[2] = "Ah, Shall I compare thee to a summer's day?"
	quotes[4] = "HaHa, A wonderful day for a duel among gentlemen!"
	quotes[5] = "Hmm the fragrance of romance is truly intoxicating."
	quotes[6] = "Luck be with me!"
	quotes[7] = "I'd wear a fedora but they haven't invented them yet!"
	quotes[8] = "that's good to spend sometimes with lower folks!"
	quotes[9] = "Step first, Thrust later."
	quotes[10] = "Ah... relationships are such a bother!"
	quotes[11] = "EN garde buffoon"
	quotes[12] = "Fairly well!"
	quotes[13] = "You have bad form my friend."
	quotes[14] = "Any ladies here that need saving?"
	quotes[15] = "Here we go !"

	// initializing the seed
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 15
	randomInt := rand.Intn(max - min) + min

	return quotes[randomInt]
}

