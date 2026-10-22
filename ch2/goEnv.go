package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You re using ", runtime.Compiler, " ")
	fmt.Println("on", runtime.GOARCH, "machine")
	fmt.Println("Using version", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}
