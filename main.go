package main

import (
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
	"os"
)

// Read url.txt
func loadDestination() (url string) {
	f, err := os.Open("url.txt")
	if err != nil {
		fmt.Print("Erorr")
	}
	defer f.Close()

	txt, _ := ioutil.ReadAll(f)
	 url = string(txt)
	return url
}

// send http request
func post(userin, url string) error {
	jsonStr := `{"content":"` + userin + `"}`
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func main() {
	var userin string
	var url string

	url = loadDestination()
	fmt.Scan(&userin)
	resp := post(userin, url)
	if resp != nil {
		fmt.Println(resp)
	} else {
		fmt.Print("Send your message!")
	}
}
