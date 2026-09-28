package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/mGO"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	iLog := log.New(f, "customLogNumber: ", log.LstdFlags)

	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
