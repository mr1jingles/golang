package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"
)

var dateNotValid error = errors.New("Date is not valid")

func parseText(line string) (time.Time, error) {
	layout1 := "02/Jan/2006:15:04:05 -0700"
	layout2 := "02-Jan-2006:15:04:05 -0700"

	dt, err := time.Parse(layout1, line)

	if err == nil {
		return dt, nil
	}

	dt, err = time.Parse(layout2, line)

	if err != nil {
		return dt, fmt.Errorf("%w: %w", dateNotValid, err)
	}

	return dt, nil
}

func main() {
	logs := []string{
		"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200] 325504",
		"127.0.0.1 - - [16/Nov/2017:10:16:41 +0200] GET /CVEN HTTP/1.1\" 200 12531 \"-\"\"Mozilla/5.0 AppleWebKit/537.36",
		"127.0.0.1 200 9412 - - [12/Nov/2017:06:26:05 +0200] \"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
		"[12/Nov/2017:16:27:21 +0300]",
		"[12/Nov/2017:20:88:21 +0200]",
		"[12/Nov/2017:20:21 +0200]",
		"127.0.0.1 - - [16-Nov-2017:10:49:46 +0200] 325504",
		"127.0.0.1 - - [16-Nov-2017:10:16:41 +0200] GET /CVEN HTTP/1.1\" 200 12531 \"-\"\"Mozilla/5.0 AppleWebKit/537.3",
		"127.0.0.1 200 9412 - - [12-Nov-2017:06:26:05 +0200] \"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
	}

	re := regexp.MustCompile(`\[(.*)\]`)

	for _, s := range logs {
		if date := re.FindStringSubmatch(s); date != nil {
			dt, err := parseText(date[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			fmt.Println(dt.Format(time.RFC850))

		} else {
			fmt.Printf("String not match a regular expression: %s\n", s)
		}
	}
}
