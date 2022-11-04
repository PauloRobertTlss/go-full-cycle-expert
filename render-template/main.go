package main

import (
	"os"
	"text/template"
)

type Cep struct {
	Numero string `json:"cep"`
	Url    string `json:"-"`
}

func main() {

	myCep := Cep{Numero: "06280090", Url: "http://"}
	render := template.New("CepRender")
	template, _ := render.Parse("Seu cep: {{.Numero}} na url: {{.Url}}")
	err := template.Execute(os.Stdout, myCep)
	if err != nil {
		panic(err)
	}
}
