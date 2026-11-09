package main

import "fmt"

func terraform(planets *[6]string) {
	for idx, p := range *planets {
		planets[idx] = "New" + p
	}
}

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
	}

	fmt.Println(planets)

	terraform(&planets)

	fmt.Println(planets)
}
