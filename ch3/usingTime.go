package main

import (
	"fmt"
	"time"
)

func main() {
	utime := time.Now()

	fmt.Println("Unix time:", utime.Unix())

	fmt.Println(utime, utime.Format(time.RFC3339))

	fmt.Println(utime.Weekday(), utime.Day(), utime.Month(), utime.Year())
	time.Sleep(time.Second)
	t1 := time.Now()

	fmt.Println("Elapsed:", t1.Sub(utime))
	fmt.Println("Formatted", t1.Format("02 January 2006"))

	loc, _ := time.LoadLocation("Europe/Moscow")
	londonTime := t1.In(loc)
	fmt.Println("London:", londonTime)
}
