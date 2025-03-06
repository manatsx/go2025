package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", cat)
	http.ListenAndServe(":8080", nil)
}

func cat(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/15/Cat_August_2010-4.jpg/181px-Cat_August_2010-4.jpg">
	`)
}
