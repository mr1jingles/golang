package main

import (
	"fmt"
)

const (
	A = 10
	B = 20
)

func main() {
	var x float64 = A / B

	fmt.Println(x)
	var y float64 = A / 20.0
	fmt.Println(y)

	var z float64 = float64(A) / B

	fmt.Println(z)

	const s1 = 123

	var v1 float32 = s1 * 0.5
	var v2 = s1 * 0.5

	fmt.Printf("v1: %v, type: %T\n", v1, v1)
	fmt.Printf("v2: %v, type: %T\n", v2, v2)
	//var v3 int = s1 * 0.5
	//fmt.Printf("v3: %v, type: %T\n", v3, v3)
}
