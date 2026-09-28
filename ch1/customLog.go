package main

import (
	"log"
	"os"
)

var LOGFILE = "/tmp/myLog"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	mylog := log.New(f, "customLog: ", log.LstdFlags)

	mylog.Println("Hello there!")
	mylog.Println("Another log entry!")

}
