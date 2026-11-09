package main

import (
	"fmt"
)

type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12

	fmt.Println(v1)
	fmt.Println(PI)

	const (
		zero Digit = iota
		one
		two
		three
	)

	fmt.Println(zero, one, two, three)

	const (
		p2_0 Power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)

	fmt.Println(p2_0, p2_2, p2_4, p2_6)

	type ServerStatus int

	const (
		Active ServerStatus = iota
		Running
		Failed
		Starting
	)

	fmt.Println(Starting)
}
