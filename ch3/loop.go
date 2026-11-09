package main

import (
	"log"
)

func main() {

	var i int = 0
	log.SetFlags(log.LstdFlags)

	for {
		if i > 10 {
			break
		}

		log.Println("i:", i)
		i++
	}
}
