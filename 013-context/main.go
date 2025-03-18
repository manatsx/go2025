package main

import (
	"context"
	"fmt"
	"time"
)

// El paquete context en Go se usa para manejar cancelaciones, deadlines y valores compartidos entre goroutines de manera controlada.
// Es útil para controlar la vida útil de procesos concurrentes, como solicitudes HTTP, consultas a bases de datos o tareas en segundo plano.

// Definimos un tipo personalizado para las claves de contexto.
// Esto evita colisiones de claves cuando se usan paquetes externos.
type contextKey string

func main() {
	// Creamos un contexto base vacío.
	ctx := context.Background()

	// Agregamos un valor al contexto con la clave "my-key".
	// Se usa `contextKey("my-key")` en lugar de un string literal por seguridad.
	ctx = context.WithValue(ctx, contextKey("my-key"), "sarasa")

	// Llamamos a `viewContext` para ver el valor almacenado en el contexto.
	viewContext(ctx)

	// Creamos un nuevo contexto con un timeout de 5 segundos.
	// Esto significa que el contexto se cancelará automáticamente después de 5 segundos.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // `defer` asegura que se libere el contexto cuando `main` termine.

	// Iniciamos una goroutine que ejecuta `myProcess` con el contexto.
	go myProcess(ctx)

	// Bloqueamos la ejecución hasta que el contexto expire o se cancele.
	<-ctx.Done()

	// Cuando el contexto expira, se ejecuta este código.
	fmt.Println("we've exceeded the deadline") // Mensaje indicando que se superó el tiempo límite.
	fmt.Println(ctx.Err())                     // Imprime el error asociado (`context.DeadlineExceeded`).

	// Intentamos obtener el valor almacenado en el contexto original.
	// Sin embargo, este `ctx` es el de `context.WithTimeout`, por lo que el valor anterior se pierde.
	fmt.Printf("my value is %s\n", ctx.Value("my-key")) // Imprime `<nil>` porque `my-key` no existe en este `ctx`.
}

// Función que recupera e imprime el valor de "my-key" desde el contexto.
func viewContext(ctx context.Context) {
	fmt.Printf("my value is %s\n", ctx.Value(contextKey("my-key")))
}

// Función que simula un proceso en segundo plano.
// Se ejecuta en una goroutine y se detiene cuando `ctx.Done()` es activado.
func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Si el contexto se cancela o expira, la goroutine se detiene.
			fmt.Println("ohh, time out")
			return
		default:
			fmt.Println("doing some process") // Mientras el contexto esté activo, seguimos trabajando.
		}
		time.Sleep(500 * time.Millisecond) // Pausa de 500ms entre iteraciones.
	}
}
