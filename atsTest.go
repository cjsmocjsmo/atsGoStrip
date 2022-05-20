package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpRequestHandler)
	http.ListenAndServeTLS(":8081", "go-server.crt", "go-server.key", nil)
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
