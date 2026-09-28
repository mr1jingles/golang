package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	rsyslog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "LOGFATAL")

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(rsyslog)
	}

	log.Fatal(rsyslog)
	fmt.Println("will you see it?")
}
