package main

import "fmt"

// Implement a function that computes the difference between two lists.
// The function should remove all occurrences of elements from the first list (a) that are present in the second list (b).
// The order of elements in the first list should be preserved in the result.

// Examples
// If a = [1, 2] and b = [1], the result should be [2]
// If a = [1, 2, 2, 2, 3] and b = [2], the result should be [1, 3]
func ArrayDiff(a, b []int) []int {
	return a
}

func main() {
	arr1 := []int{1, 2}
	arr2 := []int{1, 2, 2, 2, 3}
	fmt.Println(ArrayDiff(arr1, arr2))

	arr3 := []int{1, 2}
	arr4 := []int{1}
	fmt.Println(ArrayDiff(arr3, arr4))
}

// Implemente una función que calcule la diferencia entre dos listas.
// La función debe eliminar todas las apariciones de elementos de la primera lista () que estén presentes en la segunda lista ().
// El orden de los elementos de la primera lista debe conservarse en el resultado.ab

// Tengo una función que recibe dos arrays de enteros
// Debo eliminar todos los elementos de la primera lista que se encuentren en la segunda lista
// Debo devolver el resultado de la diferencia de la primera lista
