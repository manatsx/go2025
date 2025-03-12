package main

import (
	"fmt"
	"strings"
)

/*
 * Escribe una función que reciba dos palabras (String) y retorne
 * verdadero o falso (Bool) según sean o no anagramas.
 * - Un Anagrama consiste en formar una palabra reordenando TODAS
 *   las letras de otra palabra inicial.
 * - NO hace falta comprobar que ambas palabras existan.
 * - Dos palabras exactamente iguales no son anagrama.
 */

func main() {
	fmt.Println(isItAnAnagram("roma", "amor"))        // true
	fmt.Println(isItAnAnagram("listen", "silent"))    // true
	fmt.Println(isItAnAnagram("hello", "world"))      // false
	fmt.Println(isItAnAnagram("roma", "roma"))        // false
	fmt.Println(isItAnAnagram("casa", "saca"))        // true
	fmt.Println(isItAnAnagram("anagrama", "nagaram")) // false
}

func isItAnAnagram(word1 string, word2 string) bool {
	a := strings.ToLower(word1)
	b := strings.ToLower(word2)

	if a == b {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	count := make(map[rune]int)

	for _, val := range a {
		count[val]++
	}

	for _, val := range b {
		count[val]--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}

	return true

}
