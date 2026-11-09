package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 5)
	s1 = append(s1, -1000)

	fmt.Println(s1)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)

	reSlice[0] = -100
	reSlice[1] = 123456

	fmt.Println(s1)
	fmt.Println(reSlice)

	s1 = nil

	fmt.Println(s1)
	fmt.Println(reSlice)
}
