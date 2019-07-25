package main

import (
	"log"
	"net/http"
)

func main() {

	req, err := http.NewRequest("GET", "http://apilayer.net/api/validate", nil)
	if err != nil {
		log.Fatalln(err)
	}

	q := req.URL.Query()
	q.Add("access_key", "")
	q.Add("country_code", "")
	q.Add("format", "")
	q.Add("number", "")

	req.URL.RawQuery = q.Encode()

}
