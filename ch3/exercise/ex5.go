package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const layout string = "15-04"

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s 15-04\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
	}

	dt, err := time.Parse(layout, os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		printUsage()
	}

	fmt.Println("Correct date")
	fmt.Println("Full:", dt)
	fmt.Printf("__%d : %d__\n", dt.Hour(), dt.Minute())
}
