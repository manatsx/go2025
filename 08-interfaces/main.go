package main

import (
	"fmt"

	"github.com/go2025/08-interfaces/vehicles"
)

func main() {
	fmt.Println("VEHICLES")

	carV := vehicles.Car{Time: 160}
	fmt.Println(carV.Distance())

	vArray := []string{"Car", "Motorcycle", "Truck", "Sarasa"}
	var d float64
	for _, v := range vArray {
		fmt.Printf("Vechile: %s\n", v)
		veh, err := vehicles.New(v, 400)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println()
			continue
		}
		distance := veh.Distance()
		fmt.Printf("Distance: %.2f\n", distance)
		fmt.Println()
		d += distance
	}
	fmt.Printf("Total distance: %.2f\n", d)
}
