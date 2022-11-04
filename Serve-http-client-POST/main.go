package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	client := http.Client{Timeout: time.Second} // <-- context deadline exceeded
	/**
	 O Objeto Request é isolado o Client e attachamos c.Do()
	 */
	request, err := http.NewRequest("POST", "https://google.com.br", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Accept", "application/json")

	jsonVar := bytes.NewBuffer([]byte(`{"name": "BigPolll"}`)) //<-- required buffer
	resp, err := client.Post("https://google.com.br", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	io.CopyBuffer(os.Stdout, resp.Body, nil)

}
