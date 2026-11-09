package main

import (
	"fmt"
)

func main() {
	var amap map[string]int
	fmt.Println(amap)

	for _, _ = range amap {
		fmt.Println("ok")
	}
	amap["test"] = 111
	fmt.Println(amap)
}
