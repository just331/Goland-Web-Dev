package main

import (
	"net/http"
	"io"
)

func d(res http.ResponseWriter, req *http.Response){
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Response) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
