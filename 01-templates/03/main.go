package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	person := Person{
		Name: "Alighn",
		Age:  19,
	}

	err := tpl.Execute(os.Stdout, person)
	if err != nil {
		log.Fatalln(err)
	}
}
