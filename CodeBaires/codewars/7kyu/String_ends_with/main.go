package main

import (
	"fmt"
	"strings"
)

// Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).
// Examples:
// solution("abc", "bc") // returns true
// solution("abc", "d") // returns false
// Link: https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d/train/go

func main() {
	// se le pasa a la function solution los argumentos
	fmt.Println(solution("abc", "bc"))
}

// Function con parametros de tipo string
func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)

}

// Debo crear una funcion que reciba dos parametros que son string
// Debo verificar si el string del primer argumento termina con un string en el segundo argumento
// Si el primer argumento termina devolver true, caso contrario devolver false
