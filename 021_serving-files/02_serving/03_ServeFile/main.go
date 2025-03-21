package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", cat)
	http.HandleFunc("/cat.jpg", catPic)
	http.ListenAndServe(":3030", nil)
}

func cat(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="cat.jpg">	`)
}

func catPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "cat.jpg")
}
