# Manejador (Handler)

_Esto es una de las cosas más importantes que debes conocer_

[http.Handler](https://godoc.org/net/http#Handler)

```Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Un controlador (handler) responde a una solicitud HTTP. [Handler.ServeHTTP] debe escribir los encabezados (headers) y los datos de respuesta en ResponseWriter y, a continuación, devolver.

Un controlador recibe una solicitud (request), la procesa y devuelve una respuesta (response)

## Flujo típico de un controlador en una API REST

1. Recibe la solicitud HTTP (GET, POST, PUT, DELETE, etc.).

2. Extrae y valida los datos de la solicitud.

3. Llama a los servicios o lógica de negocio para procesar la solicitud.

4. Devuelve una respuesta al cliente, normalmente en formato JSON.

---

# Servidor (Server)

[http.ListenAndServe](https://godoc.org/net/http#ListenAndServe)

```Go
func ListenAndServe(addr string, handler Handler) error
```

ListenAndServe escucha en la dirección de red TCP addr y luego llama a Serve con un controlador para gestionar las solicitudes en las conexiones entrantes. Las conexiones aceptadas se configuran para habilitar las conexiones activas TCP.

El controlador normalmente es nulo, en cuyo caso se utiliza DefaultServeMux .

ListenAndServe siempre devuelve un error distinto de nulo.

[http.ListenAndServeTLS](https://godoc.org/net/http#ListenAndServeTLS)

```Go
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```

_Observa que ambas funciones anteriores toman un manejador (handler)_

---

# Solicitud (Request)

Consulta [http.Request](https://godoc.org/net/http#Request) en la documentación.

Una solicitud (Request) representa una solicitud HTTP recibida por un servidor o que será enviada por un cliente.

Aquí está con la mayoría de los comentarios y algunos campos eliminados:

```go
type Request struct {
    Method string
    URL *url.URL
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string
    // Este campo solo está disponible después de llamar a ParseForm.
    Form url.Values
    // Este campo solo está disponible después de llamar a ParseForm.
    PostForm url.Values
    MultipartForm *multipart.Form
    // RemoteAddr permite que los servidores HTTP y otro software registren
	// la dirección de red que envió la solicitud, generalmente para
	// propósitos de registro.
    RemoteAddr string
}
```

También consulta el índice que muestra el tipo [Request]() del paquete `http`.

Algunas cosas interesantes que puedes hacer con una solicitud:

## Obtener datos de URL y formularios

`http.Request` es una estructura. Tiene los campos `Form` y `PostForm`. Si leemos la documentación sobre estos, vemos:

```
    // Form contiene los datos del formulario analizados, incluidos tanto
    // los parámetros de consulta del campo URL como los datos del formulario POST o PUT.
    // Este campo solo está disponible después de llamar a **ParseForm**.
    // El cliente HTTP ignora Form y usa Body en su lugar.
    Form url.Values

    // PostForm contiene los datos del formulario analizados desde los parámetros
    // del cuerpo de POST, PATCH o PUT.
    // Este campo solo está disponible después de llamar a **ParseForm**.
    // El cliente HTTP ignora PostForm y usa Body en su lugar.
    PostForm url.Values
```

Si observamos **ParseForm**,

```go
func (r *Request) ParseForm() error
```

vemos que es un método adjunto a `*http.Request`.

---

Si observamos **FormValue**,

```go
func (r *Request) FormValue(key string) string
```

vemos que es un método adjunto a `*http.Request`. `FormValue` devuelve el primer valor del componente nombrado en la consulta. Los parámetros del cuerpo de POST y PUT tienen prioridad sobre los valores de la cadena de consulta de la URL. `FormValue` llama a `ParseMultipartForm` y `ParseForm` si es necesario e ignora cualquier error devuelto por estas funciones. Si la clave no está presente, `FormValue` devuelve una cadena vacía. Para acceder a múltiples valores de la misma clave, llama a `ParseForm` y luego inspecciona directamente `Request.Form`.

---

## Ver el método HTTP

El tipo `http.Request` es una estructura que tiene un campo `Method`.

---

## Ver valores de URL

El tipo `http.Request` es una estructura que tiene un campo `URL`. Observa que el tipo es un `*url.URL`.

Consulta el tipo `url.URL`:

```go
type URL struct {
    Scheme     string
    Opaque     string    // datos opacos codificados
    User       *Userinfo // información de usuario y contraseña
    Host       string    // host o host:puerto
    Path       string
    RawPath    string // sugerencia de ruta codificada (solo en Go 1.5 y posteriores; ver método EscapedPath)
    ForceQuery bool   // agrega un signo de interrogación ('?') incluso si RawQuery está vacío
    RawQuery   string // valores de consulta codificados, sin el signo '?'
    Fragment   string // fragmento para referencias, sin el signo '#'
}
```

---

## Trabajar con los encabezados HTTP

El tipo `http.Request` es una estructura que tiene un campo `Header`.

De la documentación sobre la estructura `http.Header`, vemos que:

```
type Header map[string][]string
```

---

# Respuesta (Response)

[http.ResponseWriter](https://godoc.org/net/http#ResponseWriter)

Un controlador HTTP utiliza una interfaz ResponseWriter para construir una respuesta HTTP.

No se puede utilizar un ResponseWriter después de que [Handler.ServeHTTP] haya regresado.

```Go
type ResponseWriter interface {
    // Header devuelve el mapa de encabezados que se enviará con WriteHeader.
    // Cambiar el encabezado después de llamar a WriteHeader (o Write) no tiene efecto.
    Header() Header

    // Write escribe los datos en la conexión como parte de una respuesta HTTP.
    //
    // Si WriteHeader aún no se ha llamado, Write lo llama
    // con WriteHeader(http.StatusOK) antes de escribir los datos.
    // Si el encabezado no contiene una línea Content-Type, Write agrega un Content-Type
    // establecido en el resultado de pasar los primeros 512 bytes de datos escritos a
    // DetectContentType.
    Write([]byte) (int, error)

    // WriteHeader envía un encabezado de respuesta HTTP con un código de estado.
    // Si WriteHeader no se llama explícitamente, la primera llamada a Write
    // activará implícitamente WriteHeader(http.StatusOK).
    // Por lo tanto, las llamadas explícitas a WriteHeader se utilizan principalmente para
    // enviar códigos de error.
    WriteHeader(int)
}
```

---

## Configurar un encabezado de respuesta

Un `http.ResponseWriter` tiene un método `Header()` que devuelve un `http.Header`.

Consulta la documentación de `http.Header`:

```Go
type Header map[string][]string
```

Consulta los métodos adjuntos al tipo `http.Header`:

```go
type Header
func (h Header) Add(key, value string)
func (h Header) Del(key string)
func (h Header) Get(key string) string
func (h Header) Set(key, value string)
func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```

Podemos configurar encabezados para una respuesta así:

```Go
res.Header().Set("Content-Type", "text/html; charset=utf-8")
```
