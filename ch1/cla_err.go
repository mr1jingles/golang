package main

import (
	"errors"
	"log"
	"os"
	"strconv"
)

func main() {
	var err = errors.New("not a digit")
	args := os.Args
	k := 1

	if len(args) == 1 {
		log.Fatalln("Give me arguments")
		os.Exit(1)
	}

	var min, max float64
	for err != nil {
		if k >= len(args) {
			break
		}

		n, err := strconv.ParseFloat(args[k], 64)

		if err != nil {
			log.Printf("skip %s", args[k])
		} else {
			min, max = n, n
		}

		k++
	}

	for i := 2; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)

		if err == nil {
			if n > max {
				max = n
			}

			if n < min {
				min = n
			}
		}
	}

	log.Println("Max:", max)
	log.Println("Min:", min)
}
