# Técnicas que aprenderemos

- Concatenar
- Canalización de comandos CLI - salida a un archivo con `>`

# Código que utilizaremos de la biblioteca estándar

## [os.Create](https://godoc.org/os#Create)

Esto nos permite crear un archivo.

```Go
func Create(name string) (*File, error)
```

Create crea o trunca el archivo nombrado. Si el archivo ya existe, se trunca. Si el archivo no existe, se crea con el modo 0o666 (antes de umask). Si la operación es correcta, se pueden utilizar los métodos del archivo devuelto para E/S; el descriptor de archivo asociado tiene el modo O_RDWR. El directorio que contiene el archivo ya debe existir. Si hay un error, será del tipo \*PathError.

**Truncar un archivo** significa reducir su tamaño a cero, es decir, borrar todo su contenido. Cuando un archivo ya existe y se "trunca", todo lo que contenía se elimina, pero el archivo sigue existiendo en el sistema de archivos. Esta acción se realiza sin eliminar el archivo en sí; solo se borra su contenido.

### ¿Qué son esos números `0o666`?

Esos números son un **modo de archivo** que especifican los permisos de acceso del archivo en el sistema operativo. Están representados en **notación octal** (base 8), y cada dígito tiene un significado específico sobre qué acciones se permiten sobre el archivo.

#### Desglosando `0o666`:

- **`0o`**: El prefijo `0o` indica que el número es en **sistema octal** (base 8).
- **`666`**: Es un número en octal que define los **permisos de lectura y escritura** para el propietario del archivo, el grupo y otros usuarios.

##### Desglosando `666`:

- El primer **6** (más a la izquierda) se refiere a los permisos del **propietario** del archivo.
- El segundo **6** se refiere a los permisos del **grupo** al que pertenece el archivo.
- El tercer **6** se refiere a los permisos de **otros usuarios**.

Cada número (0-7) en notación octal es la suma de los permisos que puedes otorgar:

- **4**: Permiso de lectura (r)
- **2**: Permiso de escritura (w)
- **1**: Permiso de ejecución (x)

Por lo tanto, el número `6` se calcula así:

- **4 (lectura) + 2 (escritura) = 6**.

Esto significa que el propietario, el grupo y los demás usuarios tienen **permiso de lectura y escritura** sobre el archivo.

### Resumen:

- **Truncar un archivo** significa borrar su contenido sin eliminar el archivo en sí.
- **`0o666`** son los permisos de acceso al archivo, donde:
  - **6** (propietario): puede leer y escribir.
  - **6** (grupo): puede leer y escribir.
  - **6** (otros): puede leer y escribir.

Por lo tanto, `os.Create` crea o trunca un archivo, y si el archivo ya existe, lo trunca, dejando que el propietario, el grupo y otros usuarios puedan leerlo y escribir en él.

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
