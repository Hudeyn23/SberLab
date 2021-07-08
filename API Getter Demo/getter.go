package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/borodun/SberLab/core"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	makeRequest()
}

func parseCredentials(filename string) (string, string) {
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	accesKey := csvLines[1][1]
	secretKey := csvLines[1][2]
	return accesKey, secretKey
}

func createSigner() core.Signer {
	var filename string = "/home/artem/Documents/credentials.csv"
	var accesKey, secretKey = parseCredentials(filename)

	fmt.Printf("accesKey: %s secretKey: %s\n", accesKey, secretKey)

	s := core.Signer{
		Key:    accesKey,
		Secret: secretKey,
	}

	return s
}

func makeRequest() {
	var s = createSigner()

	req, _ := http.NewRequest("GET", "https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/0b5a73ddd98027372f2ec00668b88856/todos/detail?offset=1&limit=10", ioutil.NopCloser(bytes.NewBuffer([]byte(""))))

	req.Header.Add("content-type", "application/json")
	s.Sign(req)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
