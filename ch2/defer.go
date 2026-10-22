package main

import (
	"fmt"
)

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	fmt.Println("d1:")
	d1()
	fmt.Println()
	fmt.Println("d2:")
	d2()
	fmt.Println()
	fmt.Println("d3:")
	d3()
	fmt.Println()
}
