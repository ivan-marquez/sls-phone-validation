package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// APIResponse struct with numverify Json structure
type APIResponse struct {
	Valid    bool
	Number   string
	LineType string
}

func validatePhone(phoneNumber string) APIResponse {
	b := make([]byte, 350)
	var response APIResponse

	res, err := getPhoneValidation(phoneNumber)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	defer res.Body.Close()

	err = json.Unmarshal(b, &response)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	return response
}

func createRequest(phoneNumber string) (*http.Request, error) {
	const urlEndpoint = "http://apilayer.net/api/validate"
	req, err := http.NewRequest("GET", urlEndpoint, nil)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	q := req.URL.Query()
	q.Add("access_key", "e0ccae9a9264df43761e785a5434363b")
	q.Add("country_code", "")
	q.Add("format", "1")
	q.Add("number", phoneNumber)

	req.URL.RawQuery = q.Encode() //?

	return req, nil
}

func getPhoneValidation(phoneNumber string) (*http.Response, error) {
	const t = time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: t,
	}

	req, err := createRequest(phoneNumber)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return res, nil
}
