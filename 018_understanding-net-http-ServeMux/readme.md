# Términos sinónimos en Programación Web

- router
- request router (enrutador de solicitudes)
- multiplexer (multiplexor)
- mux
- servemux
- server (servidor)
- http router (enrutador HTTP)
- http request router (enrutador de solicitudes HTTP)
- http multiplexer (multiplexor HTTP)
- http mux
- http servemux
- http server (servidor HTTP)

---

En electrónica, un [multiplexor (o mux)](https://en.wikipedia.org/wiki/Multiplexer) es un dispositivo que selecciona una de varias señales de entrada y reenvía la entrada seleccionada a una única línea de salida.

El término **multiplexor** ha sido adoptado en la programación web para referirse al proceso de enrutamiento de solicitudes.

Un servidor web recibe solicitudes en diferentes rutas y a través de distintos métodos HTTP. Por ejemplo, podríamos tener las siguientes solicitudes:

### **Solicitud #1**

- Ruta: `/cat`
- Método: `GET`

### **Solicitud #2**

- Ruta: `/apply`
- Método: `GET`

### **Solicitud #3**

- Ruta: `/apply`
- Método: `POST`

Con base en las solicitudes entrantes, el servidor debe determinar cómo responder a cada una. Para cada solicitud recibida, se ejecutará un código diferente.

Aunque hemos estado usando la palabra **servidor**, también podríamos haber usado el término **multiplexor** o **mux**. El servidor, multiplexor o mux, determina qué código debe ejecutarse en respuesta a cada solicitud entrante.

---

## **ServeMux en Go**

`ServeMux` es un multiplexor de solicitudes HTTP.

Un `ServeMux` **compara la URL de cada solicitud entrante con una lista de patrones registrados** y llama al controlador (handler) que más se asemeje a la URL.

Los patrones pueden ser rutas fijas o subrutas. Algunos ejemplos:

- `"/favicon.ico"` → Ruta exacta
- `"/images/"` → Subruta (los patrones con `/` al final abarcan todos los elementos dentro de esa subruta)

Si hay rutas similares, **los patrones más largos tienen prioridad**. Por ejemplo:

- `/images/` y `/images/thumbnails/`
- Una solicitud a `/images/thumbnails/photo.jpg` será manejada por `/images/thumbnails/`
- Mientras que `/images/logo.png` será manejada por `/images/`

Además, si un patrón **termina en `/`**, cualquier solicitud a la raíz de la subruta sin la barra final será redirigida automáticamente.

Ejemplo:

- Si registramos `"/images/"`, una solicitud a `"/images"` se redirigirá a `"/images/"` automáticamente.

Pero si registramos **tanto** `"/images"` como `"/images/"`, la redirección no ocurrirá.

`ServeMux` también permite patrones con nombres de host:

- `"codesearch.google.com/"` solo manejará solicitudes de ese dominio.

Además, **`ServeMux` limpia las rutas automáticamente**, eliminando `.` y `..` en la URL o múltiples barras `/`.

---

## **`http.ServeMux` en Go**

📌 **Documentación oficial**: [http.ServeMux](https://godoc.org/net/http#ServeMux)

```go
type ServeMux
	func NewServeMux() *ServeMux
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

Cualquier valor de tipo `*http.ServeMux` **implementa la interfaz `http.Handler`**, lo que significa que podemos pasarlo a `http.ListenAndServe`.

🔹 **Interfaz `http.Handler` en Go**:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Esto significa que podemos pasar `*http.ServeMux` a `http.ListenAndServe` como manejador:

```go
func ListenAndServe(addr string, handler Handler) error
```

---

## **Ejemplo de uso de `http.ServeMux`**

```go
package main

import (
	"fmt"
	"net/http"
)

// Manejador para "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenido a la página principal")
}

// Manejador para "/about"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Página de información - Acerca de nosotros")
}

func main() {
	// Crear un nuevo ServeMux
	mux := http.NewServeMux()

	// Registrar rutas y sus manejadores
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	// Iniciar el servidor
	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
```

---

## **Uso de `mux.Handle` con `http.Handler`**

Si tienes un tipo que implementa `http.Handler`, puedes usar `mux.Handle` en lugar de `mux.HandleFunc`:

```go
package main

import (
	"fmt"
	"net/http"
)

// Tipo que implementa http.Handler
type CustomHandler struct{}

func (h CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Respuesta desde un http.Handler")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/custom", CustomHandler{}) // Asigna el manejador personalizado

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
```

---

## **`DefaultServeMux` en Go**

`http.ListenAndServe` inicia un servidor HTTP con una dirección y un manejador. **Si el manejador es `nil`, se usa `http.DefaultServeMux` automáticamente**.

```go
http.ListenAndServe(":8080", nil)
```

Esto significa que podemos registrar rutas sin necesidad de crear un `ServeMux` explícitamente:

```go
package main

import (
	"fmt"
	"net/http"
)

// Manejadores
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenido a la página principal")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Página de información - Acerca de nosotros")
}

func main() {
	// Registrar rutas en DefaultServeMux
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Iniciar el servidor (usará DefaultServeMux)
	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```

---

## **Resumen**

✅ **`ServeMux` es un multiplexor HTTP que enruta solicitudes basadas en la URL**.  
✅ **Los patrones más largos tienen prioridad sobre los más cortos**.  
✅ **Redirige automáticamente las rutas sin barra `/` al final** (a menos que se registren manualmente).  
✅ **Puede manejar rutas específicas por host**.  
✅ **Sanitiza las rutas eliminando `.` y `..`**.  
✅ **Implementa `http.Handler`, por lo que se puede pasar a `http.ListenAndServe`**.  
✅ **Si `nil` se pasa a `ListenAndServe`, se usa `DefaultServeMux` automáticamente**.

---
