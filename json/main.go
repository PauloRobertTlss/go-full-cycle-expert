package main

import (
	json2 "encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
	Senha  int `json:"-"` //omiti
}

// https://mholt.github.io/json-to-go/

func main() {

	conta := Conta{Numero: 15422, Saldo: 1222}
	//pego o valor
	json, err := json2.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(json))
	// já entrego para um arquivo, ou webserve
	encoder := json2.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	//DECODE recebr um Json e hidrata a conta
	jsonPuro := []byte(`{"numero": 15422,"saldo": 15422}`)
	constXP := Conta{}
	err = json2.Unmarshal(jsonPuro, &constXP) //<-- importante passar o ponteiro
	if err != nil {
		panic(err)
	}
	println(constXP.Saldo)

	//DECODE com tags no struct
	jsonPuroTag := []byte(`{"n": 15422,"s": 15422}`)
	constXPTag := Conta{}
	err = json2.Unmarshal(jsonPuroTag, &constXPTag) //<-- importante passar o ponteiro
	if err != nil {
		panic(err)
	}
	println(constXPTag.Saldo)

}
