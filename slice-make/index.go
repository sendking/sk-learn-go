package main

import "fmt"

func main() {
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "C", "P", "H", "M", "E")
	fmt.Println(dwarfs)
}
