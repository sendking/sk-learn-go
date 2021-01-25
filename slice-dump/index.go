package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v%v\n", label, len(slice), cap(slice), slice)
}
func main() {
	dwarfs := []string{"C", "P", "H", "M", "E"}
	dump("dwarfs", dwarfs)
	// ? 这里为什么是 cap是4
	dump("dwarfs[1:2]", dwarfs[1:2])
}
