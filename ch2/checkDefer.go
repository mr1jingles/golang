package main

import (
	"log"
)

func main() {
	x := 10

	defer log.Println(x)

	x = 20

	defer log.Println(x)

	for i := 0; i < 3; i++ {
		defer log.Println(x)
	}

	log.Println("closures")
	for i := 0; i < 3; i++ {
		defer func() {
			log.Println(i)
		}()
	}
	x = 100
}
