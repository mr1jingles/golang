package main

import "fmt"

const (
	pow4_0 int = 1 << (2 * iota)
	_
	pow4_2
	pow4_3
	pow4_4
)

func main() {
	fmt.Println("pow4_0:", pow4_0)
	fmt.Println("pow4_2:", pow4_2)
	fmt.Println("pow4_3:", pow4_3)
	fmt.Println("pow4_4:", pow4_4)
}
