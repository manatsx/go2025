# Técnicas que aprenderemos

- Concatenar
- Canalización de comandos CLI - salida a un archivo con `>`

# Código que utilizaremos de la biblioteca estándar

## [os.Create](https://godoc.org/os#Create)

Esto nos permite crear un archivo.

```Go
func Create(name string) (*File, error)
```

---

## [defer](https://golang.org/ref/spec#Defer_statements)

La palabra clave `defer` nos permite posponer la ejecución de una sentencia hasta que la función en la que se ha colocado el `defer` termine de ejecutarse.

---

## [io.Copy](https://godoc.org/io#Copy)

Esto nos permite copiar de una fuente a un destino.

```Go
func Copy(dst Writer, src Reader) (written int64, err error)
```

## [strings.NewReader](https://godoc.org/strings#NewReader)

`NewReader` devuelve un nuevo lector que lee desde una cadena.

```Go
func NewReader(s string) *Reader
```

## [os.Args](https://godoc.org/os#pkg-variables)

`Args` es una variable del paquete `os`. `Args` contiene los argumentos de la línea de comandos, comenzando con el nombre del programa.

```Go
var Args []string
```

---
