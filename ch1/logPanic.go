package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	rsyslog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "LOGPANIC")

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(rsyslog)
	}

	log.Panic(rsyslog)
	fmt.Println("Will you see it?")
}
