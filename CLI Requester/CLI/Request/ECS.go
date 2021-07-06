package Request

import (
	"bytes"
	"fmt"
	"github.com/borodun/SberLab/core"
	"io"
	"io/ioutil"
	"net/http"
)

func MakeRequest(reqUrl string, accessKey string, secretKey string) {
	signer := core.Signer{
		Key:    accessKey,
		Secret: secretKey,
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
