package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	massaAddress string = "A12VViTPc3ZQu3PGp53oYXDPA6iydaeT69abcFTM2QXo6BFKMsv7"
	nodeAddress  string = "http://localhost:33035/"
)

func main() {
	responce, err := getMassaWebsite(massaAddress, nodeAddress)
	if err != nil {
		panic(err)
	}
	err = responceToZipFile(responce)
	if err != nil {
		panic(err)
	}
	unZipFile("site1.zip")
}

func getMassaWebsite(massaAddress string, nodeAddress string) (*http.Response, error) {
	// we are looking for the key "massa_web" in decimal
	body := []byte(`{
		"jsonrpc":"2.0",
		"method":"get_datastore_entries",
		"params":[[{
			"address":"` + massaAddress + `",
			"key":[109,97,115,115,97,95,119,101,98]
		}]],
		"id":1}`)

	resp, err := http.Post(nodeAddress, "application/json", bytes.NewBuffer(body))

	return resp, err
}

type responceBodyGetDatastoreEntries struct {
	Jsonrpc 	float64 		`json:"jsonrpc"`
	Result		[]interface{}	`json:"result"`
	Id			int				`json:"id"`
} 

func responceToZipFile(responce *http.Response) error {
	// TODO

	// var body responceBodyGetDatastoreEntries
	// json.NewDecoder(responce.Body).Decode(body)
	// fmt.Printf("%v\n", body.Result...)

	return nil
}

func unZipFile(fileName string) error {
	// https://golang.cafe/blog/golang-unzip-file-example.html
	dirNameUnZip := "output"
    archive, err := zip.OpenReader(fileName)
    if err != nil {
        panic(err)
    }
    defer archive.Close()

    for _, f := range archive.File {
        filePath := filepath.Join(dirNameUnZip, f.Name)

        if !strings.HasPrefix(filePath, filepath.Clean(dirNameUnZip)+string(os.PathSeparator)) {
            return fmt.Errorf("invalid file path")
        }
        if f.FileInfo().IsDir() {
            os.MkdirAll(filePath, os.ModePerm)
            continue
        }

        if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
            panic(err)
        }

        dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil {
            panic(err)
        }

        fileInArchive, err := f.Open()
        if err != nil {
            panic(err)
        }

        if _, err := io.Copy(dstFile, fileInArchive); err != nil {
            panic(err)
        }

        dstFile.Close()
        fileInArchive.Close()
    }
	return nil
}