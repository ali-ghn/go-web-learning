package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type CustomeHandler string

func (ch CustomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Method        string
		Submission    url.Values
		Header        http.Header
		URL           *url.URL
		ContentLength int64
	}{
		r.Method,
		r.Form,
		r.Header,
		r.URL,
		r.ContentLength,
	}
	fmt.Printf("%v\n", data.Method)
	fmt.Printf("%v\n", data.Submission)
	fmt.Printf("%v\n", data.Header)
	fmt.Printf("%v\n", data.URL)
	fmt.Printf("%v\n", data.ContentLength)
}

func main() {
	var ch CustomeHandler
	http.ListenAndServe(":8081", ch)
}
