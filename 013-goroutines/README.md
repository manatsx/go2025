### **Goroutines y Channels en Go**

Go es un lenguaje diseñado para la concurrencia, y su característica principal para lograr esto son las **goroutines** y los **channels**.

---

## **1. Goroutines: Ejecución Concurrente en Go**

Una **goroutine** es una función que se ejecuta de manera concurrente con otras goroutines dentro del mismo programa. Es similar a un hilo (thread), pero más liviana y manejada por el runtime de Go en lugar del sistema operativo.

### **Ejemplo Básico de Goroutine**

```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("¡Hola desde una goroutine!")
}

func main() {
	go sayHello() // Ejecutamos la función en una goroutine

	fmt.Println("¡Hola desde main!")

	time.Sleep(1 * time.Second) // Espera para que la goroutine termine antes de que el programa finalice
}
```

### **Explicación**

1. **`go sayHello()`**: Lanza la función `sayHello` en una goroutine.
2. **No hay garantía de orden**: El `main` puede terminar antes de que la goroutine se ejecute, por lo que usamos `time.Sleep` para darle tiempo.

---

## **2. Channels: Comunicación entre Goroutines**

Las goroutines no comparten memoria de forma automática, por lo que Go proporciona los **channels**, que permiten la comunicación segura entre goroutines.

### **Ejemplo Básico de Channel**

```go
package main

import "fmt"

func main() {
	messages := make(chan string) // Creamos un channel de tipo string

	// Goroutine que envía un mensaje
	go func() {
		messages <- "Hola desde la goroutine"
	}()

	// Recibimos el mensaje en el main
	msg := <-messages
	fmt.Println(msg)
}
```

### **Explicación**

1. **`make(chan string)`**: Creamos un channel de tipo `string`.
2. **`messages <- "Hola"`**: La goroutine envía un mensaje al channel.
3. **`<-messages`**: Bloquea la ejecución hasta que reciba un valor del channel.

---

## **3. Channels Bufferizados vs. No Bufferizados**

- **Channel NO bufferizado** (bloquea hasta que haya un receptor):
  ```go
  ch := make(chan int) // No tiene espacio de almacenamiento
  ```
- **Channel bufferizado** (permite almacenar valores sin necesidad de un receptor inmediato):
  ```go
  ch := make(chan int, 3) // Puede almacenar hasta 3 valores antes de bloquearse
  ```

### **Ejemplo de Channel Bufferizado**

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Channel bufferizado con capacidad de 2

	ch <- 1
	ch <- 2

	fmt.Println(<-ch) // Imprime 1
	fmt.Println(<-ch) // Imprime 2
}
```

---

## **4. Cierre de Channels y Range en Channels**

Cuando un channel se **cierra**, ya no puede enviar más valores, pero aún se pueden recibir los valores restantes.

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // Cerramos el canal

	// Leer valores con un range
	for val := range ch {
		fmt.Println(val)
	}
}
```

---

## **5. Select para Multiples Channels**

El `select` en Go permite escuchar múltiples canales y ejecuta el caso que esté listo primero.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Mensaje de ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Mensaje de ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Recibido:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Recibido:", msg2)
	}
}
```

### **Explicación**

- Se ejecuta el **primer canal que esté listo**.
- Si ambos canales están bloqueados, `select` esperará hasta que uno de ellos tenga datos.

---

## **Resumen Final**

✅ **Goroutines** → Son funciones concurrentes, se crean con `go function()`.  
✅ **Channels** → Permiten comunicación segura entre goroutines (`make(chan tipo)`).  
✅ **Bufferizados vs. No bufferizados** → Los no bufferizados bloquean hasta que se lean los datos.  
✅ **Cierre de channels** → Se usa `close(chan)` para indicar que no se enviarán más datos.  
✅ **Select** → Escucha múltiples canales y ejecuta el primero que esté listo.
