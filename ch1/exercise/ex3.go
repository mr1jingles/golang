package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Ready")
	for scanner.Scan() {
		text := scanner.Text()

		if text == "END" {
			log.Println("Bye!")
			break
		}

		n, err := strconv.Atoi(text)

		if err != nil {
			log.Printf("not a digit %s, skip\n", text)
			continue
		}

		log.Printf("valid %d", n)
	}
}
