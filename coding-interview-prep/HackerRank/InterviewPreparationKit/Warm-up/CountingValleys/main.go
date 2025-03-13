package main

import "fmt"

func main() {
	fmt.Println(countingValleys2(12, "DDUUDDUDUUUD"))
}

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	level := 0  // nivel del mar
	valley := 0 // valle

	// Recorro los pasos que dio el excursionista
	for _, step := range path {
		// Verifico si dio pasos hacia arriba (U) o hacia abajo (D)
		if step == 'U' {
			// si dio pasos hacia arriba significa que esta por encima del nivel del mar
			level++
		} else if step == 'D' {
			// si dio pasos hacia abajo significa que esta por debajo del nivel del mar
			level--
		}

		// si esta en el nivel 0 y dio pasos hacia arriba entonces está en un valle
		if level == 0 && step == 'U' {
			valley++
		}
	}

	return int32(valley)
}

func countingValleys2(steps int32, path string) int32 {
	level := 0
	valley := 0

	for _, step := range path {
		if step == 'D' {
			if level == 0 {
				valley++

			}
			level -= 1

		} else {
			level += 1
		}
	}
	return int32(valley)
}

// Un excursionista realiza una caminata en las montañas y va anotando todos los pasos que da
// Si los pasos que da es hacia arriba escribe U y si sus pasos son para abajo, escribe D
// Lo que tengo que hacer es encontrar el total de Valles que atravezó
// Cada paso inicia sobre el nivel del mar

// Una montaña es una secuencia de pasos consecutivos por encima del nivel del mar, que comienza con un paso hacia arriba (U) desde el nivel del mar y termina con un paso hacia abajo (D) hasta el nivel del mar
// Un valle es una secuencia de pasos consecutivos por debajo del nivel del mar, que comienza con un paso hacia abajo (D) desde el nivel del mar y termina con un paso hacia arriba (U) hasta el nivel del mar.

// El valle empieza con pasos hacia abajo (D) desde el nivel del mar y termina con pasos arriba (U) hasta el nivel del mar
// Para entrar en un valle tiene que bajar una vez desde el nivel del mar y subir una vez hasta el nivel del mar

// Si el excursionista anota los pasos que da, significa que está llevando adelante un conteo (voy a necesitar un contador)
// Si debo encontrar el total de valles que atravezo entonces también necesito hacer un conteo del valle que atravezó (voy a necesitar otro contador)
// Cada paso inicia desde el nivel del mar (0) (voy a necesitar un contador para saber cuando entro o salio del nivel del mar)

// Paso 1. Debo inicializar dos variables (nivel del mar (contador), valle (contador) )
// Paso 3. Debo recorrer cada paso
// Paso 4: Debo verificar si dio un paso hacia arriba U e indicar dio pasos por encima del nivel del mar
// Paso 5: Debo verificar si dio un paso hacia abajo D e indicar que dio pasos por debajo del nivel del mar
// Paso 6: Debo verificar si su paso esta en el nivel del mar y luego dio pasos hacia arriba, de ser asi encontro un valle
// Paso 7: Retornar la cantidad de valles recorridos
