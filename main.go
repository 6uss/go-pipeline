package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello worlds t1<h1>")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}