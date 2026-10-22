package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	myversion := runtime.Version()

	major := strings.Split(myversion, ".")[0][2]
	minor := strings.Split(myversion, ".")[1]

	majorInt, err := strconv.Atoi(string(major))

	if err != nil {
		fmt.Println(err)
		return
	}

	minorInt, err := strconv.Atoi(minor)

	if err != nil {
		fmt.Println(err)
		return
	}

	if majorInt < 1 && minorInt < 8 {
		fmt.Println("You're using wrong version")
		return
	}

	fmt.Println("You are using version > 1.8")
}
