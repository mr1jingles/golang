package main

import (
	"fmt"
)

func main() {
	myArray := [5]string{"aa", "bb", "cc", "dd", "ee"}
	myMap := make(map[int]string)

	for idx, value := range myArray {
		myMap[idx] = value
	}

	fmt.Println(myMap)
}
