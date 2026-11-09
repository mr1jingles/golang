package main

import "fmt"

// dump для длины, вместимости и содержимого среза
func dump(label string, slice []string) {
	fmt.Printf("%v: длина %v, вместимость %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	dump("dwarfs", dwarfs)           // Выводит: dwarfs: длина 5, вместимость 5 [Церера Плутон Хаумеа Макемаке Эрида]
	dump("dwarfs[1:2]", dwarfs[4:5]) // Выводит: dwarfs[1:2]: длина 1, вместимость 4 [Плутон]
}
