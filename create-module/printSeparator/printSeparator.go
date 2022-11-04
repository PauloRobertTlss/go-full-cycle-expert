package printSeparator

import (
	"bytes"
	"fmt"
)

func PrintSeparator() {
	var buffer bytes.Buffer

	for i := 0; i < 35; i++ {
		buffer.WriteString("---")
	}
	fmt.Println(buffer.String())
}
