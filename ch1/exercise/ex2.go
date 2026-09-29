package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	var count int = 0
	var sum float64 = 0
	if len(os.Args) == 1 {
		log.Println("need float arguments")
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.ParseFloat(os.Args[i], 64)

		if err != nil {
			log.Println(err)
			continue
		}

		count++
		sum += n
	}

	log.Println("Result is:", sum/float64(count))
}
