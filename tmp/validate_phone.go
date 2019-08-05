package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const urlEndpoint = "http://apilayer.net/api"

// APIResponse struct with numverify Json structure
type APIResponse struct {
	valid    bool
	number   string
	lineType string
}

func main() {

	// req, err := http.NewRequest("GET", "http://apilayer.net/api/validate", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// q := req.URL.Query()
	// q.Add("access_key", "")
	// q.Add("country_code", "")
	// q.Add("format", "")
	// q.Add("number", "")

	// req.URL.RawQuery = q.Encode()

	res := APIResponse{}
	str := `{"valid":true,"number":"17873626144","line_type":"mobile"}`
	err := json.Unmarshal([]byte(str), &res)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v", res)
}
