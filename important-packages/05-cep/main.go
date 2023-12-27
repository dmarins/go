package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// go run import-packages/05-cep/main.go 20020080
// cat import-packages/05-cep/address.txt

// go build -o import-packages/05-cep/cep import-packages/05-cep/main.go
// ./import-packages/05-cep/cep 20020080
func main() {
	for _, cep := range os.Args[1:] {

		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail: %v\n", err)
		}

		var data Address

		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail: %v\n", err)
		}

		file, err := os.Create("./import-packages/05-cep/address.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail: %v\n", err)
		}

		defer file.Close()

		file.WriteString(fmt.Sprintf("CEP: %s, Logradouro: %s, UF: %s", data.Cep, data.Logradouro, data.Uf))

	}
}
