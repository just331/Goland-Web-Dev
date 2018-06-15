package _3

import (
	"net/http"
	"io"
	"html/template"
	"log"
)

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo go")
}

func bar(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w,"bar go")
}

func myName(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil{
		log.Fatalln("error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "template.gohtml", "Justin")
	if err != nil{
		log.Fatalln("error executing template", err)
	}
}


func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.HandleFunc("/dog/", http.HandlerFunc(bar))
	http.HandleFunc("/me/", http.HandlerFunc(myName))
	http.ListenAndServe(":8080", nil)
}
