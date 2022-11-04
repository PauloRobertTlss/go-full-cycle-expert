package main

import (
	"bytes"
	"fmt"
)

func Soma[T int | float64](m map[string]T) T {
	var total T
	for _, number := range m {
		total += number
	}
	return total

}

func main() {
	m := map[string]int{"big": 1542, "lara": 154}

	fmt.Printf("Sum | Sun [%v] \n", Soma(m))
	PrintSeparator()

}

func PrintSeparator() {
	var buffer bytes.Buffer
	for i := 0; i < 35; i++ {
		buffer.WriteString("---")
	}
	fmt.Println(buffer.String())
}
