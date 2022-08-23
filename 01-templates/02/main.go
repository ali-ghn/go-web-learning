package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	names := []string{"James", "John"}

	err := tpl.Execute(os.Stdout, names)
	if err != nil {
		log.Fatalln(err)
	}
}
