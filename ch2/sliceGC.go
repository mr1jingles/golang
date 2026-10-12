package main

import (
	"runtime"
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

	runtime.GC()

	_ = structure[0]
}
