package main

import "fmt"

func main() {
	d1 := []string{"C", "P", "H", "M", "E"}
	d2 := append(d1, "O")
	d3 := append(d2, "S", "Q", "S")
	fmt.Print(d1, d2, d3)
}
