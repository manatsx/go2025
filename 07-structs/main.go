package main

import (
	"encoding/json"
	"fmt"

	"github.com/go2025/07-structs/structs"
)

func main() {
	var p1 structs.Product
	p1.ID = 1
	p1.Name = "Product 1"
	fmt.Println(p1)

	p2 := structs.Product{
		ID:   2,
		Name: "Product 2",
		TypeProduct: structs.Type{
			ID:          1,
			Code:        "TP1",
			Description: "Type Product 1",
		},
		Count: 10,
		Price: 100.50,
		Tags:  []string{"tag1", "tag2"},
	}
	fmt.Println(p2)

	p3 := structs.Product{
		ID:   3,
		Name: "Heladera marca sarasa",
		TypeProduct: structs.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Price: 100.50,
		Count: 10,
		Tags:  []string{"heladera", "sarasa", "freezer", "refrigerador"},
	}
	v, err := json.Marshal(p3)
	fmt.Println(err)
	fmt.Println(string(v))

	fmt.Println(p3.GetPrice())

	p3.SetName("Heladera marca sarasa 2")
	fmt.Println(p3.Name)

	p3.AddTags("nuevo", "usado")
	fmt.Println(p3.Tags)
}
