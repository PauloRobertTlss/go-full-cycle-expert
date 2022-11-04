package main

//"text/template" <-- não é recomentado | o Html blida de alguma injesões..
import (
	"html/template"
	"os"
)

type Cep struct {
	Numero string `json:"cep"`
	Url    string `json:"-"`
}

type Collection []Cep

func main() {

	myCep := Cep{Numero: "06280090", Url: "http://"}
	myCepTwo := Cep{Numero: "06280091", Url: "http://"}
	//short template com myst
	identifyRender := "render.html"
	render := template.Must(template.New(identifyRender).ParseFiles(identifyRender))
	err := render.Execute(os.Stdout, Collection{myCep, myCepTwo})
	if err != nil {
		panic(err)
	}
}
