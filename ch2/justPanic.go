package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("Give me argument(s)")
	}

	fmt.Println("Thanks for the argiment(s)")
}
