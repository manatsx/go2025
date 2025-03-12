package main

import "fmt"

// Given an array of integers your solution should find the smallest integer.

// For example:

// Given [34, 15, 88, 2] your solution will return 2
// Given [34, -345, -1, 100] your solution will return -345
// You can assume, for the purpose of this kata, that the supplied array will not be empty.
// Link: https://www.codewars.com/kata/55a2d7ebe362935a210000b2/train/go

func main() {
	fmt.Println(SmallestIntegerFinder([]int{34, 15, 88, 2}))
	fmt.Println(SmallestIntegerFinder([]int{34, -345, -1, 100}))
}

func SmallestIntegerFinder(numbers []int) int {
	smallest := numbers[0]
	for _, val := range numbers {
		if val < smallest {
			smallest = val
		}
	}
	return smallest
}

// Recibo un array de numeros positivos y/o negativos
// y debo encontrar el numero minimo de los elementos
// si lo encuentro lo retorno
