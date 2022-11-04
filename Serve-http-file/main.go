package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //multplex ..  controle | user o nil não é remendável ListenAndServe(":8081", nil)
	fileServe := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileServe)
	log.Fatal(http.ListenAndServe(":8081", mux))

}
