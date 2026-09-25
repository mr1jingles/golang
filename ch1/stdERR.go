package main

import (
	"io"
	"os"
)

func main() {
	myString := ""

	if len(os.Args) == 1 {
		myString = "Give me a string!"
	} else {
		myString = os.Args[1]
	}

	io.WriteString(os.Stdout, "This is stdout\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
