package main

import (
	"fmt"
	"runtime"
	"time"
)

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var structure []data

	structure = make([]data, N)
	for i := 0; i < N; i++ {
		structure[i] = data{i, i}
	}

	start := time.Now()
	runtime.GC()
	fmt.Println("It took:", time.Since(start))

	_ = structure[0]
}
