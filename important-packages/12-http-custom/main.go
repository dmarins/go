package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
