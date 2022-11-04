package main

import (
	"io"
	"net/http"
)

func main() {

	request, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	defer request.Body.Close() //so ciclo de execução defer adiciona essa linha 16 por ultimo, servi para algo importante e pode ficar no top.
	readContent, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(readContent))
	request.Body.Close()

}
