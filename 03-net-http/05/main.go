package main

import (
	"fmt"
	"net/http"
)

func Handler01(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello there 01")
}

func Handler02(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello there 02")
}

func main() {

	http.HandleFunc("/01/", Handler01)

	http.HandleFunc("/02", Handler02)

	http.ListenAndServe(":8081", nil)

}
