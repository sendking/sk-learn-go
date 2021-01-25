package main

import "fmt"

func main() {
	dwarfs := []string{"C", "P", "H", "M", "E"}
	dwarfs = append(dwarfs, "S", "Q", "S")
	fmt.Println(dwarfs, len(dwarfs))
}
