package main

import (
	"bytes"
	"fmt"
)

//struct não é uma classe
// GO não tem extends Herança (tudo é composição)

type Client struct {
	name    string
	age     int
	active  bool
	Address //composite
}

type Address struct {
	street string
	number int
	code   string
}

type PersonInterface interface {
	//apenas methods uma interface
	Disable()
}

//comportamentos (Methods)
func (c *Client) Disable() {
	c.active = false
}

func ProvideDisable(person PersonInterface) {
	person.Disable()
}

func main() {
	PrintSeparator()
	bigPool := Client{name: "Bill Gates", age: 89, active: true}
	bigPool.code = "0628090"
	bigPool.number = 1542
	bigPool.street = "Rua tote carvalho"
	fmt.Println(bigPool)
	PrintSeparator()
	bigPool.Disable()
	fmt.Println(bigPool)
	ProvideDisable(&bigPool)
	PrintSeparator()
	fmt.Println(bigPool)

}

func PrintSeparator() {
	var buffer bytes.Buffer
	for i := 0; i < 35; i++ {
		buffer.WriteString("---")
	}
	fmt.Println(buffer.String())
}
