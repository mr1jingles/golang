package main

import (
	"fmt"
)

func main() {

	aSlice := make([]int, 20, 100)
	fmt.Printf("Len: %d, Cap: %d\n", len(aSlice), cap(aSlice))

	fmt.Printf("Type of nil slice %T\n", aSlice)

	newSlice := aSlice[:5]

	newSlice[1] = 100

	fmt.Println(aSlice)
}
