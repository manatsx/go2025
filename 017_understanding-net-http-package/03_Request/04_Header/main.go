package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type product int

// Este método se ejecuta cuando el servidor recibe una solicitud
func (m product) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//  analiza los datos del formulario enviado en una solicitud HTTP
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string      // Método de solicitud HTTP (GET, POST, PUT, DELETE, etc.)
		URL         *url.URL    // URL de la solicitud
		Submissions url.Values  // Datos del formulario enviado en una solicitud HTTP
		Header      http.Header // Encabezados de la solicitud
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
	}

	// renderiza la plantilla index.gohtml y pasa los datos del formulario (req.Form) a la plantilla
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var p product
	http.ListenAndServe(":3000", p)

}
