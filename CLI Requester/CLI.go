package main

import (
	"CLI/Configuration"
	"bytes"
	"fmt"
	"github.com/borodun/SberLab/core"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"net/http"
)

var offset string
var limit string
var configPath string
var config Configuration.Config

func makeRequest(reqUrl string) {
	signer := core.Signer{
		Key:    config.AccessKey,
		Secret: config.SecretKey,
	}

	req, err := http.NewRequest("GET", reqUrl, ioutil.NopCloser(bytes.NewBuffer([]byte(""))))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("content-type", "application/json")
	err = signer.Sign(req)
	if err != nil {
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func showECS(cmd *cobra.Command, args []string) {
	var reqUrl = fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers/detail?offset=%s&limit=%s", config.ProjectID, offset, limit)
	makeRequest(reqUrl)
}

func main() {
	var cmdECS = &cobra.Command{
		Use:   "ecs",
		Short: "Operations with ecs",
		Long:  `Operations with elastic cloud servers on SberCloud`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			config, _ = Configuration.LoadConfig(configPath)
		},
	}

	var cmdShowECS = &cobra.Command{
		Use:   "show",
		Short: "Show list of ecs",
		Long:  `Prints list of elastic cloud servers in project`,
		Run:   showECS,
	}
	cmdECS.AddCommand(cmdShowECS)
	cmdShowECS.PersistentFlags().StringVarP(&offset, "offset", "o", "1", "Number of first entity")
	cmdShowECS.PersistentFlags().StringVarP(&limit, "limit", "l", "1", "Number of entities")

	var rootCmd = &cobra.Command{
		Use:   "sber",
		Short: "CLI until for SberCloud",
		Long:  "CLI util for SberCloud with api reqests",
	}
	rootCmd.AddCommand(cmdECS)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", ".", "Path to the folder with config")

	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
