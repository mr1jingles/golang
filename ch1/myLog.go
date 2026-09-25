package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "my-program:", log.Lshortfile|log.LstdFlags)

	logger.Println("Test message")
}
