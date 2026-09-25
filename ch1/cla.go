package main

import (
	"fmt"
	"os"
	"strconv"
)

func convert(s string) (float64, bool) {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("init: skip value", s)
		return 0, false
	}

	return n, true
}

func main() {
	args := os.Args

	init := false
	var min_value float64
	var max_value float64
	for i := 1; i < len(args); i++ {
		n, valid := convert(args[i])

		if valid && !init {
			min_value = n
			max_value = n
			init = true
			continue
		}

		if !valid {
			continue
		}

		if n < min_value {
			min_value = n
		}

		if n > max_value {
			max_value = n
		}
	}

	if !init {
		fmt.Println("No valid arguments")
		os.Exit(1)
	}

	fmt.Println("Max:", max_value)
	fmt.Println("Min:", min_value)
}
