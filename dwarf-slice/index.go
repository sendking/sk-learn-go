package main

import "fmt"

func main() {
	// 切片
	dwarfs := []string{"C", "P", "H", "M", "E"}
	// 数组
	dwarfsArray := [...]string{"C", "P", "H", "M", "E"}

	fmt.Printf("Slice %T value: %v\n", dwarfs, dwarfs)
	fmt.Printf("Array %T value: %v", dwarfsArray, dwarfsArray)
}
