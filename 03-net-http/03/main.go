package main

import (
	"fmt"
	"net/http"
)

type CustomeHandler string

func (ch CustomeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<h1>Hello there</h1>")
}

func main() {
	var ch CustomeHandler
	http.ListenAndServe(":8081", ch)
}
