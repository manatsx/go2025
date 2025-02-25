package main

import (
	"fmt"
	"net/http"
)

type product int

func (m product) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
