package main

import (
	"fmt"
	"net/http"
)

type CustomeHandler string

func (ch CustomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from custom handler")
}

func main() {
	var ch CustomeHandler
	http.ListenAndServe(":8081", ch)
}
