package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

func GetUrlResponse(url, token string) ([]byte, error) {
	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("User-Agent", "curl/7.54.1")

	fmt.Println("requestedURL: ", url, req)
	resp, err := client.Do(req)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("response error: %s, resp: %d", err.Error(), int(resp.StatusCode)))
		return nil, err
	}

	defer resp.Body.Close()
	return jsonData, nil

}
