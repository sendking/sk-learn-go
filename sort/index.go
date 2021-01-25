package main

import (
	"fmt"
	"sort"
)

func main() {
	dwarfs := []string{"C", "P", "H", "M", "E"}
	sort.StringSlice(dwarfs).Sort()
	fmt.Println(dwarfs)
	planets := []string{"C", "P", "H", "M", "E", "S", "K"}
	sort.Strings(planets)
	fmt.Println(planets)
}
