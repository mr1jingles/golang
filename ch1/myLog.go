package main

import (
	"log"
	_ "os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Test message")
}
