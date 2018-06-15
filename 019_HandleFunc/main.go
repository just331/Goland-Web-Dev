package main

import (
	"net/http"
	"io"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res,"dog dog dog")
}

func c(res http.ResponseWriter, req *http.ResponseWriter) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)

}