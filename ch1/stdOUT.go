package main

import (
	"io"
	"os"
)

func main() {
	mystring := ""

	if len(os.Args) == 1 {
		mystring = "Give me a string"
	} else {
		mystring = os.Args[1]
	}

	io.WriteString(os.Stdout, mystring)
	io.WriteString(os.Stdout, "\n")
}
