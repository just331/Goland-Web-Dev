package main

import (
	"text/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil{
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil{
		log.Fatalln("error creating file", err)
	}

	err = tpl.Execute(nf, nil)
	if err != nil{
		log.Fatalln(err)
	}
}
