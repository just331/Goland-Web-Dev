package _2_hands_on

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
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":8080", nil)
}
