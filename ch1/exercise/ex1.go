package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	var sum int = 0

	if len(os.Args) == 1 {
		log.Println("give me arguments")
		return
	}

	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.Atoi(os.Args[i])

		if err != nil {
			log.Println("not a digit!", os.Args[i])
			continue
		}

		sum += n
	}

	log.Println("Sum is:", sum)
}
