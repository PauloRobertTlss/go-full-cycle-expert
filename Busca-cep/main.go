package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	cep := Cep{Numero: "06280090", Url: "https://viacep.com.br/ws/01001000/json/"}

	for _, input := range os.Args[1:] {
		cep.Numero = input
		req, err := http.Get("https://viacep.com.br/ws/" + input + "/json/")
		if err != nil {
			println(err)
		}
		defer req.Body.Close()
		content, err := io.ReadAll(req.Body)
		var data ViaCEP
		err = json.Unmarshal(content, &data)
		if err != nil {
			println(err)
		}

		fileCep, err := os.Create(cep.Numero + ".json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ciar arquivo %v\n", err)
		}
		defer fileCep.Close()
		//_, err = fileCep.WriteString(fmt.Sprintf("CEP: %s | Localidade: %s", data.Cep, data.Localidade))
		// já entrego para um arquivo, ou webserve
		encoder := json.NewEncoder(fileCep)
		encoder.Encode(data)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprint(os.Stdout)

	}

}
