package main

import (
	"runtime"
)

func main() {
	var N = 40000000
	myMap := make(map[int]*int, N)

	for i := 0; i < N; i++ {
		myMap[i] = &i
	}

	runtime.GC()

	_ = myMap[0]
}
