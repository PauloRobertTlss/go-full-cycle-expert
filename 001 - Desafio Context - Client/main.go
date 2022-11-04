package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Cep struct {
	Numero string `json:"cep"`
	Url    string `json:"-"`
}

func main() {

	cep := Cep{Numero: "06280090", Url: "https://viacep.com.br/ws/01001000/json/"}

	for _, input := range os.Args[1:] {
		cep.Numero = input
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8088/", nil)
		if err != nil {
			panic(err)
		}

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		//select {
		//case <-ctx.Done():
		//
		//	break
		//
		//}

		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		content, err := io.ReadAll(response.Body)
		var data ViaCEP
		err = json.Unmarshal(content, &data)
		if err != nil {
			panic(err)
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
		io.Copy(os.Stdout, response.Body)
		//reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		//text, _ := reader.ReadString('\n')
		//fmt.Fprint(os.Stdout)

	}

}
