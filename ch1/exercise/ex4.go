package main

import (
	"bytes"
	"log"
	"os"
)

var LOGFILE1 = "/tmp/myLog1"
var LOGFILE2 = "/tmp/myLog2"

func main() {
	f1, err := os.OpenFile(LOGFILE1, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.OpenFile(LOGFILE2, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	var buffer []byte
	buf := bytes.NewBuffer(buffer)

	mylog := log.New(buf, "customLog: ", log.LstdFlags)

	mylog.Println("Hello there!")
	mylog.Println("Another log entry!")

	n, err := f1.Write(buf.Bytes())

	if err != nil {
		log.Println(err)
	} else {
		log.Println("wrote", n, "bytes")
	}

	n, err = f2.Write(buf.Bytes())

	if err != nil {
		log.Println(err)
	} else {
		log.Println("wrote", n, "bytes")
	}
}
