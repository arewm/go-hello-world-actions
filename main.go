package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, out there!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
