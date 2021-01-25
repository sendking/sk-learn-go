package main

import "fmt"

func main() {
	planets := [...]string{
		"M",
		"V",
		"E",
		"M",
		"J",
		"S",
		"U",
		"N",
	}
	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]
	fmt.Println(terrestrial, gasGiants, iceGiants)
	fmt.Println(gasGiants[0])
	// 更改了切片，原数组会发生改变
	iceGiantsMarkII := iceGiants
	iceGiants[1] = "P"
	fmt.Println(planets)
	fmt.Println(iceGiants, iceGiantsMarkII)
	// 复制原数组不会发生改变
	iceGiantsMarkIII := planets
	planets[0] = "S"
	fmt.Println(iceGiantsMarkIII)
}
