package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	reqBody := `{
		"name": "LordGift"
	}`

	req, err := http.NewRequest("POST", "https://httpbin.org/post", strings.NewReader(reqBody))

	//request header
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	var client http.Client
	resp, err := client.Do(req)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//response header
	for key, values := range resp.Header {
		for _, val := range values {
			fmt.Printf("%s:%s\n", key, val)
		}
	}

	fmt.Printf("%s\n", b)

}
