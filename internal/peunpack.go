package internal

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UnpackPE(filePath string)  {

	respCode , respData := uploadFile(filePath)
	if respCode == "200"{
		fmt.Println("Ok")
		fmt.Println(respData)
	} else {
		fmt.Println("Something went wrong on unpacking ...")
	}

}

func uploadFile(filePath string) (string,[]byte) {

	    const endPoint string = "https://api.unpac.me/api/v1/private/upload"

	    err := godotenv.Load()
		Check(err)
		env , err := godotenv.Read()
		Check(err)

		apiKey := env["UNPACME_API"]
		client := &http.Client{}

		file, err := os.Open(filePath)
		Check(err)
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
		Check(err)
		io.Copy(part, file)
		writer.Close()


		req, err := http.NewRequest("POST", endPoint,body)
		Check(err)

		// Authentication
		req.Header.Add("Authorization", apiKey)
		resp, err := client.Do(req)
		Check(err)
		content, err := ioutil.ReadAll(resp.Body)
		Check(err)

		return resp.Status, content
}

