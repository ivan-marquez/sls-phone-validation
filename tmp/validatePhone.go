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

func main() {
	//res := APIResponse{
	//	LineType: "mobile",
	//	Number:   "17873626144",
	//	Valid:    true,
	//}

	//b, err := json.Marshal(res)

	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	//	os.Exit(1)
	//}

	res, err := getPhoneValidation(os.Args[1])
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	defer res.Body.Close()
	data := ioutil.ReadAll(res.Body)

	fmt.Println(data)
}

func createRequest(phoneNumber string) (*http.Request, error) {
	const urlEndpoint = "http://apilayer.net/api/validate"
	req, err := http.NewRequest("GET", urlEnpoint, nil)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	q := req.URL.Query()
	q.Add("access_key", "")
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

	req := createRequest(phoneNumber)

	res, err := client.Do(req)
	if err != nil {
		// TODO: how to handle this error?
		log.Fatalln(err)
	}

	defer res.Body.Close()
	return res, nil
}
