package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Crear una nueva fuente de aleatoriedad
	source := rand.NewSource(time.Now().Unix())
	// Crear un nuevo generador de números aleatorios
	r := rand.New(source)

	// Obtener un número aleatorio en el rango de 0 a 2
	x := r.Intn(5)
	fmt.Println(x)
}
