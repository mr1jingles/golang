package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)

	fmt.Println("Ready to read input")
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
