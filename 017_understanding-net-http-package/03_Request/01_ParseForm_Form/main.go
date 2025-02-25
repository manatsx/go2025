package main

import (
	"html/template"
	"log"
	"net/http"
)

type product int

// Este método se ejecuta cuando el servidor recibe una solicitud
func (m product) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//  analiza los datos del formulario enviado en una solicitud HTTP
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// renderiza la plantilla index.gohtml y pasa los datos del formulario (req.Form) a la plantilla
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var p product
	http.ListenAndServe(":3000", p)

}
