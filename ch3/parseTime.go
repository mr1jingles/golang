package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func printUsage() {
	fmt.Printf("Usage: %s HH:MM\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
	}

	layout := "15:04"

	if dt, err := time.Parse(layout, os.Args[1]); err == nil {
		fmt.Println("Full:", dt)
		fmt.Printf("Time: hour %d, minute %d\n", dt.Hour(), dt.Minute())
	} else {
		fmt.Println(err)
		printUsage()
	}
}
