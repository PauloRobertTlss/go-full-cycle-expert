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
	O Objeto Request é isolado do Client e attachamos c.Do()
	*/
	jsonVar := bytes.NewBuffer([]byte(`{"name": "BigPolll"}`)) //<-- required buffer
	request, err := http.NewRequest("POST", "https://google.com.br", jsonVar)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Accept", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	io.CopyBuffer(os.Stdout, resp.Body, nil)

}
