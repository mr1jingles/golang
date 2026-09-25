package main

import (
	"fmt"
)

func main() {
	a1 := "123"
	a2 := 123
	a3 := "Have a nice day\n"
	a4 := "abc"

	fmt.Print(a1, a2, a3, a4)

	fmt.Println()
	fmt.Println(a1, a2, a3, a4)
	fmt.Print(a1, " ", a2, " ", a3, " ", a4, "\n")
	fmt.Printf("%s%d %s %s\n", a1, a2, a3, a4)
}
