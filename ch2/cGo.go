package main

// #include <stdio.h>
// void callC() {
//    printf("Calling C code!\n");
// }
import "C"

import "fmt"

func main() {
	fmt.Println("Hello from go")
	C.callC()
	fmt.Println("Another go program")
}
