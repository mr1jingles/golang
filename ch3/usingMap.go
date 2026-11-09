package main

import (
	"fmt"
)

func main() {
	iMap := make(map[string]int)

	iMap["k1"] = 12
	iMap["k2"] = 13

	fmt.Println("k1:", iMap["k1"])
	fmt.Println("notExists:", iMap["notExists"])

	anotherMap := map[string]int{
		"k3": 33,
		"k4": 44,
	}
	fmt.Println(anotherMap)

	delete(anotherMap, "k3")
	delete(anotherMap, "k3")
	delete(anotherMap, "k1")
	fmt.Println(anotherMap)

	_, ok := iMap["notExists"]

	if ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key not exist")
	}

	for k, v := range iMap {
		fmt.Println("key:", k, "value:", v)
	}

	secondMap := map[string]string{
		"a": "11",
		"b": "22",
	}

	fmt.Printf("%T\n", secondMap["asd"])
}
