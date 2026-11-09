package main

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func checkVersion() {
	version := runtime.Version()
	version = version[2:]

	major, minor := strings.Split(version, ".")[0], strings.Split(version, ".")[1]

	majorInt, err := strconv.Atoi(major)

	if err != nil {
		panic(err)
	}

	minorInt, err := strconv.Atoi(minor)

	if err != nil {
		panic(err)
	}

	if majorInt < 1 || minorInt < 8 {
		panic("need go version > 1.8")
	}

	fmt.Println("version is valid")
}

type aPerson struct {
	person string
	height int
	weight int
}

func main() {
	checkVersion()

	aSlice := make([]aPerson, 0)

	aSlice = append(aSlice, aPerson{"Mihalis", 180, 90})
	aSlice = append(aSlice, aPerson{"Bill", 134, 45})
	aSlice = append(aSlice, aPerson{"Marietta", 155, 45})
	aSlice = append(aSlice, aPerson{"Epifanios", 144, 50})
	aSlice = append(aSlice, aPerson{"Athina", 134, 40})

	fmt.Println("0:", aSlice)

	sort.Slice(aSlice, func(x, y int) bool {
		return (aSlice[x]).height < (aSlice[y]).height
	})

	fmt.Println("<:", aSlice)

	sort.Slice(aSlice, func(x, y int) bool {
		return (aSlice[x]).height > (aSlice[y]).height
	})

	fmt.Println(">:", aSlice)
}
