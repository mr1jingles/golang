package main

import "fmt"

func a() {
	fmt.Println("In a() function")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Catch panic!")
		}
	}()

	fmt.Println("Do something")
	b()

	fmt.Println("exiting a()")
}

func b() {
	fmt.Println("In b() function")

	panic("Oh.. shit!")
	fmt.Println("exiting b()")
}

func main() {
	a()

	fmt.Println("exiting main()")
}
