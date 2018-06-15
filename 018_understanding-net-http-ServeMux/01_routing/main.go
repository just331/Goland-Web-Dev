package main

import (
	"net/http"
	"io"
)

type hotdog int

func (m hotdog) ServeHTTO(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")

	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}