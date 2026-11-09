package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second)

	fmt.Println("time.Sleep(1s):", time.Since(start))

	start = time.Now()
	time.Sleep(2 * time.Second)

	fmt.Println("time.Sleep(2s):", time.Since(start))

	start = time.Now()

	for i := 0; i < 200_000_000; i++ {
		_ = i
	}

	fmt.Println("Empty cycle:", time.Since(start))

	sum := 0
	start = time.Now()
	for i := 0; i < 200_000_000; i++ {
		sum += i
	}

	fmt.Println("Sum cycle:", time.Since(start))
}
