package main

import (
	"fmt"
)

func main() {
	a4 := []int{-1, -2, -3, -5}
	a6 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("a4:", a4)
	fmt.Println("a6:", a6)

	copy(a6, a4)

	fmt.Println("a4:", a4)
	fmt.Println("a6:", a6)

	a4[0] = 1000
	fmt.Println("a4:", a4)
	fmt.Println("a6:", a6)

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}

	fmt.Println("b4:", b4)
	fmt.Println("b6:", b6)

	copy(b4, b6)

	fmt.Println("b4:", b4)
	fmt.Println("b6:", b6)

	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	copy(s6, array4[:])
	fmt.Println("array4:", array4[:])
	fmt.Println("s6:", s6)

	s6 = []int{1, 1, -1, -1, 5, -5}
	fmt.Printf("array4 type: %T\n", array4)
	fmt.Printf("s6 type: %T\n", s6)

	copy(array4[:], s6)

	fmt.Println("array4:", array4[:])
	fmt.Println("s6:", s6)
	fmt.Printf("array4 type: %T\n", array4)
	fmt.Printf("s6 type: %T\n", s6)
}
