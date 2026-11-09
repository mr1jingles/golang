package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const layout string = "02-Jan-2006"

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s 02-Jan-2006\n", filepath.Base(os.Args[0]))
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
	fmt.Printf("Year %d Month %d Day %d\n", dt.Year(), dt.Month(), dt.Day())
}
