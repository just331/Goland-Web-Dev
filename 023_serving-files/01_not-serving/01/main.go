package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--not serving from our servers-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}
