package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", cat)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./resources"))))
	http.ListenAndServe(":8080", nil)
}

func cat(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/assets/cat.jpg">`)
}

/*

./assets/cat.jpg

*/
