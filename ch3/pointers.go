package main

import (
	"fmt"
)

func getPointers(i *int) {
	*i = *i * *i
}

func returnPointer(i int) *int {
	v := i * i
	return &v
}

func main() {
	var i, j int = -10, 25
	pI := &i
	pJ := &j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	*pI = 123456
	*pI--

	fmt.Println("i:", i)

	getPointers(pI)

	fmt.Println("i:", i)
	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)
}
