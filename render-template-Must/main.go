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
	//short template com myst
	render := template.Must(template.New("template.html").ParseFiles("render.email.html"))
	err := render.Execute(os.Stdout, []Cep{myCep})
	if err != nil {
		panic(err)
	}
}
