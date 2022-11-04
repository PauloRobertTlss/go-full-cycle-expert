package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8081", nil)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		fmt.Println("Request cancelled. Timeout reached.")
	case <-time.After(5 * time.Second):
		fmt.Println("Request finalizada booked.")
	}
}
