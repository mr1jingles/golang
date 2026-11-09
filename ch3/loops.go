package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}

	fmt.Println()

	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}

	fmt.Println()

	i = 0

	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Printf("expression: %d\n", i)
		i++
	}

	myArr := [5]int{1, -1, 0, 2, 3}
	for i, value := range myArr {
		fmt.Println("index:", i, "value:", value)
	}
}
