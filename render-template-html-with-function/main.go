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

	myCep := Cep{Numero: "06280090", Url: "xptop"}
	myCepTwo := Cep{Numero: "06280091", Url: "bitpaool"}
	//short template com myst
	identifyRender := "render.html"
	render := template.Must(template.New(identifyRender).Funcs(template.FuncMap{"ToUpper": ToUpper}).ParseFiles(identifyRender))
	err := render.Execute(os.Stdout, Collection{myCep, myCepTwo})
	if err != nil {
		panic(err)
	}

	// or
	//render = template.New("CepRender")
	//render.Funcs(template.FuncMap{"ToUpper": ToUpper}) <-- required before Parse or ParseFiles
	//template, _ := render.Parse("Seu cep: {{.Numero}} na url: {{.Url}}")
}

func ToUpper(s string) string {
	return `$contact ` + s
}
