package main

import (
	"bytes"
	"fmt"
)

func main() {
	total, _ := sum(1, 5, 8, 7, 4, 5, 8, 7, 8, 75, 74, 5, 7, 5, 7, 4, 4, 1, 5)
	PrintSeparator()
	fmt.Printf("x Args Total: %d\n", total)

	PrintSeparator()

}

func PrintSeparator() {
	var buffer bytes.Buffer
	for i := 0; i < 35; i++ {
		buffer.WriteString("---")
	}
	fmt.Println(buffer.String())
}

func sum(args ...int) (int, error) {
	total := 0
	for _, number := range args {
		total += number
	}
	return total, nil
}
