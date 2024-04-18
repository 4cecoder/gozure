// Package main github.com/4cecoder/gozure
package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler serves the HTTP requests.
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		return
	}
}

// main sets up the HTTP server.
func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
		return
	}
}
