package main

import (
	"fmt"
)

type Weekday int

const (
	Monday Weekday = iota
	Thuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println(Monday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Sunday:", Sunday)
}
