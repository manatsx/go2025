package main

import (
	"fmt"
)

// Función genérica para intercambiar dos valores
func Swap[T any](a, b T) (T, T) {
	return b, a
}

// Restricción para tipos numéricos que soporten la operación de suma
type Summable interface {
	int | int32 | int64 | float64 | float32
}

func Sum[T Summable](value []T) T {
	var total T
	for _, v := range value {
		total += v
	}
	return total
}

func anyType[N any](v1, v2 N) {
	fmt.Println(v1, v2)
}

func anyTypeTwo[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println(v1, v2)
}

func comparableType[N comparable](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 != v2)
}

type CustomSlice[V int | string] []V

type Number interface {
	int | int32 | int64 | float64 | float32
}

func MinNumber[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func MinN1[T N1](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Point int

type N1 interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type (
	MyGenericStruct[D any] struct {
		StrValue string
		Data     D
	}

	MyFirstData struct{}

	MySecondData struct{}
)

func (d MyFirstData) PrintOne() {
	fmt.Println("Print first")
}

func (d MySecondData) PrintTwo() {
	fmt.Println("Print second")
}

func main() {
	// Usando Swap con dos enteros
	x, y := 1, 2
	fmt.Println("Antes de Swap:", x, y)
	x, y = Swap(x, y)
	fmt.Println("Después de Swap:", x, y)

	// Usando Swap con dos cadenas
	str1, str2 := "hello", "world"
	fmt.Println("Antes de Swap:", str1, str2)
	str1, str2 = Swap(str1, str2)
	fmt.Println("Después de Swap:", str1, str2)

	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("Suma de enteros:", Sum(ints))

	// Ejemplo con un slice de flotantes
	floats := []float64{1.5, 2.5, 3.5}
	fmt.Println("Suma de flotantes:", Sum(floats))

	anyType(1, 1)
	anyType("1", "1")
	anyTypeTwo(1, "1")

	comparableType(4, 4)

	csInt := CustomSlice[int]{1, 2, 3, 4}
	fmt.Println(csInt)

	csStg := CustomSlice[string]{"a", "b", "4"}
	fmt.Println(csStg)

	//MinNumber(x,y)
	vv := MinN1(x, y)
	fmt.Println(vv)

	fd := MyGenericStruct[MyFirstData]{StrValue: "Test", Data: MyFirstData{}}
	fd.Data.PrintOne()

	sd := MyGenericStruct[MySecondData]{StrValue: "Test", Data: MySecondData{}}
	//sd.Data.PrintOne()
	sd.Data.PrintTwo()
}
