package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var N = 40000000
	myMap := make(map[int]int)

	for i := 0; i < N; i++ {
		myMap[i] = i
	}
	start := time.Now()
	runtime.GC()
	fmt.Println("It took:", time.Since(start))
	_ = myMap[0]
}
