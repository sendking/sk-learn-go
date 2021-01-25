package main

import "fmt"

func main() {
	planets := []string{"M", "V", "E", "M", "J", "S", "U", "N"}
	// 第三位为容量
	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "C")
	fmt.Println(planets, worlds, terrestrial)
}
