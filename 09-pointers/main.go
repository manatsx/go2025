package main

import "fmt"

func main() {
	var i int = 23
	fmt.Println(&i) // vemos la direccion de memoria de la variable utilizando &
	fmt.Println(i)

	var iP *int = &i
	fmt.Println(i)
	fmt.Println(iP)

	*iP = 5

	fmt.Printf("val i: %d, val pointer i: %d\n", i, *iP)
	fmt.Println()

	myVar := 30
	fmt.Printf("addrs: %p, values: %v\n", &myVar, myVar)
	increment(myVar)
	fmt.Println(myVar)
	incrementP(&myVar)
	fmt.Println(myVar)
	fmt.Println()

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("addrs: %p, values: %v\n", mySlice, mySlice)
	fmt.Printf("addrs 1: %p, addrs 2: %p, addrs 3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Println(mySlice)
	fmt.Println()
}

func increment(val int) {
	fmt.Println(&val)
	val++
}

func incrementP(val *int) {
	fmt.Println(val)
	*val++
}

func updateSlice(v []int) {
	fmt.Printf("addrs: %p\n", v)
	fmt.Printf("addrs 1: %p, addrs 2: %p, addrs 3: %p\n", &v[0], &v[1], &v[2])
	v[0] = 12
}

// func updateStruct(v *MyStruct) {
// 	fmt.Printf("addrs in function: %p\n", v)
// 	v.ID = 999
// 	v.Name = "Update Struct"
// }
