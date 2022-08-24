package main

import (
	"fmt"
	"net/http"
)

type CustomeHandler01 string

type CustomeHandler02 string

func (ch CustomeHandler01) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello there 01")
}

func (ch CustomeHandler02) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello there 02")
}

func main() {
	var ch1 CustomeHandler01
	var ch2 CustomeHandler02

	mux := http.NewServeMux()
	// will handle /01/* route
	mux.Handle("/01/", ch1)
	// will only handle /02 route
	mux.Handle("/02", ch2)
	http.ListenAndServe(":8081", mux)

}
