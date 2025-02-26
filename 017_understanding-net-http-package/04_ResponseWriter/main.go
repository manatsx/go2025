package main

import (
	"fmt"
	"net/http"
)

type product int

func (m product) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var d product
	http.ListenAndServe(":8080", d)
}
