package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Cep struct {
	Numero string `json:"cep"`
	Url    string `json:"-"`
}

// https://mholt.github.io/json-to-go/
// https://viacep.com.br/
type ViaCEP struct {
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
	mux := http.NewServeMux() //multplex ..  controle | user o nil não é remendável ListenAndServe(":8081", nil)
	mux.HandleFunc("/", BuscarCEPHandle)
	http.ListenAndServe(":8081", mux)

}

func BuscarCEPHandle(w http.ResponseWriter, r *http.Request) {

	cep := r.URL.Query().Get("cep")

	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonVia, err := fireApiViaCEP(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Hello, cep: " + cep))
	json.NewEncoder(w).Encode(jsonVia)
}

func fireApiViaCEP(cep string) (*ViaCEP, error) {

	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	content, err := io.ReadAll(req.Body)
	var data ViaCEP
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
