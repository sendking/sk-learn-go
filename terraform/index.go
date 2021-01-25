package main

import "fmt"

func terraform(planets [8]string) (result [8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
	return planets
}
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
	result := terraform(planets)
	fmt.Print(planets, result)
}
