package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	a := [3]int{4, 5, 6}

	ref := a[:]

	fmt.Println("Existing ref:\t", ref)

	t := append(s, ref...)
	fmt.Println("t slice:\t", t)

	s = append(s, ref...)

	fmt.Println("Existing s:\t", s)

	s = append(s, s...)
	fmt.Println("s+s\t\t", s)
}
