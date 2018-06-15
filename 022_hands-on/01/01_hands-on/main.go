package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "the beginning")
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog dog")
}

func myName(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Justin")
}
