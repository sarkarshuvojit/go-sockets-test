package main

import (
	"fmt"
	"net/http"
)

const (
	Host     = "localhost"
	Port     = "9988"
	Protocol = "tcp"
)

func main() {
	fmt.Println("Http server starting...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Http works!")
	})
	http.ListenAndServe(":8080", nil)
}
