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
	ter := planets[:4]
	gas := planets[4:6]
	ice := planets[6:]
	all := planets[:]
	fmt.Print(ter, gas, ice, all)
}
