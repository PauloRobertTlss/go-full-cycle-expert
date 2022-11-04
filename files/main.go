package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	//quando sabemos que será gravao um string
	fileLength, err := file.WriteString("Hello, big")

	if err != nil {
		panic(err)
	}

	fmt.Printf("tamamno: %v  \n", fileLength)
	file.Close()

	fileBytes, _ := os.Create("arquivobytes.txt")
	fileLengthBytes, err := fileBytes.Write([]byte("Escrevendo em bytes"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("tamamno em bytes: %v \n", fileLengthBytes)
	fileBytes.Close()

	// leitura
	arquivo, err := os.ReadFile("arquivobytes.txt") // os.Open() também funcionaria mas é verboso + usar o .close()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo)) // sempre é bytes tem que converter

	// leitura em stream | pense em um arquivo de 10 Gigas bem maior que sua memória disponivel

	readPerBuffetFile, err := os.Open("arquivobytes.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(readPerBuffetFile)
	buffer := make([]byte, 4) // definir o tamanho de leiture []bytes

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break // será quando chegar no fim do arquivo || qualquer erro
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivobytes.txt")

	if err != nil {
		panic(err)
	}

}
