package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	rawJson := bytes.NewBuffer([]byte(`{"name": "wesley"}`))

	c := http.Client{Timeout: time.Second}
	req, err := c.Post("https://www.google.com", "application/json", rawJson)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	io.CopyBuffer(os.Stdout, req.Body, nil)
}
