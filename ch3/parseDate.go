package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s '02 January 2006'\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	layout := "02 January-2006"
	date, err := time.Parse(layout, os.Args[1])

	if err != nil {
		fmt.Printf("Wrong date format: %s\n", os.Args[1])
		fmt.Printf("Usage: %s '02 January 2006'\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	} else {
		fmt.Println("Full:", date)
		fmt.Printf("Day: %d, Month: %s, Year: %d", date.Day(), date.Month(), date.Year())
	}
}
