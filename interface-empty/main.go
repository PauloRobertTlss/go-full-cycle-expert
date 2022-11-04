package main

import (
	"bytes"
	"fmt"
)

func main() {
	var typeOne interface{} = 10
	var typeTwo interface{} = "Novo ypdd"

	PrintSeparator()
	ShowType(typeOne)
	PrintSeparator()
	ShowType(typeTwo)
	PrintSeparator()
	//type assertion
	var xptoString interface{} = "Big Pool"
	fmt.Println(xptoString.(string))
	PrintSeparator()

	number, ok := xptoString.(int)
	fmt.Printf("%s | Int convert [%v] result - sucess? %v \n", xptoString, number, ok)
	PrintSeparator()

}

func ShowType(v interface{}) {
	fmt.Printf("Type da variable %T\n", v)
}
func PrintSeparator() {
	var buffer bytes.Buffer
	for i := 0; i < 35; i++ {
		buffer.WriteString("---")
	}
	fmt.Println(buffer.String())
}
