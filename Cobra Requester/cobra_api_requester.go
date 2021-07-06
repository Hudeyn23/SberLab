package main

import (
	"../core"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
)

var offset string
var limit string

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

	//fmt.Printf("accesKey: %s secretKey: %s\n", accesKey, secretKey)

	signer := core.Signer{
		Key:    accesKey,
		Secret: secretKey,
	}

	return signer
}

func makeRequest(reqUrl string) {
	var signer = createSigner()

	req, _ := http.NewRequest("GET", reqUrl, ioutil.NopCloser(bytes.NewBuffer([]byte(""))))

	req.Header.Add("content-type", "application/json")
	signer.Sign(req)

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

func showEcs(cmd *cobra.Command, args []string) {
	var reqUrl = fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/0b5a73ddd98027372f2ec00668b88856/cloudservers/detail?offset=%s&limit=%s", offset, limit)
	makeRequest(reqUrl)
}

func main() {
	var cmdPrintEcs = &cobra.Command{
		Use:   "ecs --offset --limit",
		Short: "Show ecs",
		Long:  `Show full list of elastic cloud servers with offset and limit`,
		Run:   showEcs,
	}

	var rootCmd = &cobra.Command{Use: "sber"}
	rootCmd.AddCommand(cmdPrintEcs)
	cmdPrintEcs.PersistentFlags().StringVarP(&offset, "offset", "o", "1", "Number of first entity")
	cmdPrintEcs.PersistentFlags().StringVarP(&limit, "limit", "l", "1", "Number of entities")

	rootCmd.Execute()
}
