package main

import "fmt"

// El objetivo de este ejercicio es determinar cuál es la combinación más cara de un teclado y una memoria USB que una persona puede comprar con un presupuesto dado.
// Se te proporcionan dos listas de precios, una para teclados y otra para memorias USB.
// La tarea es encontrar la combinación de un teclado y una memoria USB cuyo precio total sea lo más alto posible sin exceder el presupuesto.
// Si no es posible comprar ambos productos dentro del presupuesto, se debe devolver un valor de `-1`.

// https://www.hackerrank.com/challenges/electronics-shop/problem
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	// Debo sumar cada array y el resultado que me de menor o igual al presupuesto lo guardo en un map
	// Debo comparar las sumas del map para obtener el mayor de todos
	// Debo comparar si la suma más alta es menor o igual al presupuesto
	// Si la suma es menor o igual retorno la suma si es menor retorno -1

	return -1

}

func main() {
	fmt.Println(getMoneySpent([]int32{3, 2, 1}, []int32{5, 2, 8}, 10))
}

// Tengo una función con dos arrays como parametros que representan cada uno a precios
// Debo encontrar el precio más alto posible entre la combinación de precios del teclado y usb sin exceder el presupuesto
// Si no puedo comprar ambos productos retorno -1
