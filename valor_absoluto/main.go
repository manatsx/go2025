package main

import (
	"fmt"
	"math"
)

func main() {
	// Usando math.Abs para obtener el valor absoluto
	num := -5.3
	absValue := math.Abs(num)

	fmt.Println("El valor absoluto de", num, "es", absValue)
}
