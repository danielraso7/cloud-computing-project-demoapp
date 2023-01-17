package main

import (
	"fmt"
	"net/http"
)

func HttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", HttpRequest)
	http.ListenAndServe(":8080", nil)
}
