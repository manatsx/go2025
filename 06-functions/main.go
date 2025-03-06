package main

import (
	"fmt"

	"github.com/go2025/06-functions/function"
)

func main() {
	res := function.Add(1, 2)
	fmt.Println(res)

	function.RepeatString(5, "as")

	fmt.Println()
	v, err := function.Calc(function.SUM, 3, 6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value: ", v)
	}

	v, err = function.Calc(function.DIV, 3, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value: ", v)
	}
}
