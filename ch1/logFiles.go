package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	basename := filepath.Base(os.Args[0])

	rsyslog, err := syslog.New(syslog.LOG_DEBUG|syslog.LOG_LOCAL0, basename)

	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(rsyslog)
	log.Println("Hello syslog")
	rsyslog, err = syslog.New(syslog.LOG_MAIL, "Some program!")

	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(rsyslog)

	log.Println("LOG_MAIL: my message")
	fmt.Println("Can you see it?")
}
