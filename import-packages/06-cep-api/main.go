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

func main() {
	http.HandleFunc("/", getCepHandler)
	http.ListenAndServe(":8080", nil)
}

func getCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	address, err := getRemoteCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(address)
}

func getRemoteCep(cep string) (*Address, error) {
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail: %v\n", err)
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail: %v\n", err)
	}

	var c Address

	err = json.Unmarshal(res, &c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail: %v\n", err)
	}

	return &c, nil
}
