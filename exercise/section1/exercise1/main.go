package main

import "fmt"

func main() {

	arr := []int{4, 2, 5, 6, 7}

	// Si ingresamos un vector con los valores [4, 2, 5, 6, 7] se debe imprimir el siguiente mensaje:
	// Los valores del array son:  [24 22 25 26 27]
	// realizar la funcionalidad
	for i := range arr {
		arr[i] += 20
	}

	fmt.Println("Los valores del array son: ", arr)

}
