package main

import "fmt"

func main() {
	dwarfs := [5]string{"C", "P", "H", "M", "E"}
	for i, dwarf := range dwarfs {
		fmt.Print(i, dwarf)
	}
}
