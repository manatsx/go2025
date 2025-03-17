package main

import "fmt"

func plusMinus(arr []int32) {
	totalNums := float32(len(arr))

	pos, neg, n := 0, 0, 0

	for _, v := range arr {
		if v < 0 {
			neg++
		} else if v > 0 {
			pos++
		} else {
			n++
		}
	}

	fmt.Printf("%.6f\n", float32(pos)/totalNums)
	fmt.Printf("%.6f\n", float32(neg)/totalNums)
	fmt.Printf("%.6f\n", float32(n)/totalNums)
}

func main() {
	arr := []int32{1, 1, 0, -1, -1, 0}
	plusMinus(arr)
}
