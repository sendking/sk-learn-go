package main

import "fmt"

func main() {
	dwarfs := [5]string{"C", "P", "H", "M", "E"}
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}
}
