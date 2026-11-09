package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	logs := []string{
		"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200] 325504",
		"127.0.0.1 - - [16/Nov/2017:10:16:41 +0200] GET /CVEN HTTP/1.1\" 200 12531 \"-\"\"Mozilla/5.0 AppleWebKit/537.36",
		"127.0.0.1 200 9412 - - [12/Nov/2017:06:26:05 +0200] \"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
		"[12/Nov/2017:16:27:21 +0300]",
		"[12/Nov/2017:20:88:21 +0200]",
		"[12/Nov/2017:20:21 +0200]",
	}

	re := regexp.MustCompile(`\d\d/\w{3}/\d\d\d\d:\d\d:\d\d:\d\d\s+\+\d\d\d\d`)
	layout := "02/Jan/2006:15:04:05 -0700"
	for _, s := range logs {
		if date := re.FindString(s); date != "" {
			dt, err := time.Parse(layout, date)
			if err != nil {
				fmt.Println("Wrong datetime format", err)
				continue
			}

			fmt.Println(dt.Format(time.RFC850))

		} else {
			fmt.Printf("String not match: %s\n", s)
		}
	}
}
