# Pasar valores

## Envío de formularios

Podemos pasar valores desde el cliente al servidor a través de la URL o mediante el cuerpo de la solicitud.

Cuando envías un formulario, puedes usar el método "POST" o "GET". El método "POST" envía los valores del formulario a través del cuerpo de la solicitud. El método "GET" para enviar un formulario pasa los valores del formulario a través de la URL.

Lo recuerdo así:

```
post
body (cuerpo)

get
url
```

**Post** tiene cuatro letras y también **form**.

**Get** tiene tres letras y también **url**.

---

## Valores en la URL

Siempre puedes agregar valores a una URL.

Todo lo que va después del `?` es la cadena de consulta (query string), el área donde se almacenan los valores.

![anatomía de una URL](URL.png)

Los valores se almacenan en el formato `identificador=valor`.

Puedes tener varios `identificador=valor` separándolos con el `&` ampersand.

---

## Recuperar valores

Aunque hay múltiples maneras de recuperar valores, nos enfocaremos en:

[func (\*Request) FormValue](https://godoc.org/net/http#Request.FormValue)

```Go
func (r *Request) FormValue(key string) string
```

FormValue devuelve el primer valor para el componente nombrado de la consulta. Los parámetros del cuerpo de las solicitudes POST y PUT tienen prioridad sobre los valores de la cadena de consulta de la URL. FormValue llama a ParseMultipartForm y ParseForm si es necesario e ignora cualquier error devuelto por estas funciones. Si el valor de la clave no está presente, FormValue devuelve una cadena vacía. Para acceder a múltiples valores de la misma clave, llama a ParseForm y luego inspecciona directamente Request.Form.
