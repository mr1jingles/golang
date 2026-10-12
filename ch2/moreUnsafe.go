package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{0, 1, -2, 3, 4}

	pointer := &array[0]
	fmt.Print(*pointer, " ")
	memoryAddr := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(*pointer)

	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddr))
		fmt.Print(*pointer, " ")
		memoryAddr = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(*pointer)
	}
	fmt.Println()

	pointer = (*int)(unsafe.Pointer(memoryAddr))
	fmt.Print("one more: ", *pointer, " ")
	fmt.Println()
}
