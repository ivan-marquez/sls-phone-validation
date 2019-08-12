package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const urlEndpoint = "http://apilayer.net/api"

// APIResponse struct with numverify Json structure
type APIResponse struct {
	Valid    bool
	Number   string
	LineType string
}

func main() {
	res := APIResponse{
		LineType: "mobile",
		Number:   "17873626144",
		Valid:    true,
	}

	b, err := json.Marshal(res)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(b))
}

func setRequest() (*http.Request, error) {
  req, err := http.NewRequest("GET", "http://apilayer.net/api/validate", nil)
    if err != nil {
		log.Fatalln(err)
		panic(err)
	}

  q := req.URL.Query()
  q.Add("access_key", "")
  q.Add("country_code", "")
  q.Add("format", "1")
  q.Add("number", os.Args[1])

  req.URL.RawQuery = q.Encode() //?
  
  return req, nil	
}

func wip() (*http.Response, error) {
  t := time.Duration(3 * time.Second)
  client := http.Client{
    Timeout: t,
  }
  
  req := setRequest()
  
  res, err := client.Do(req)
    if err != nil {
	  log.Fatalln(err)
	}

  defer res.Body.Close()

  return res, nil
}
