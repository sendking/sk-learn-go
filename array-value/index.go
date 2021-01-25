package main

import "fmt"

func main() {
	planets := [...]string{"C", "P", "H", "M", "E"}
	planetsMarkII := planets
	planets[2] = "whoops"
	fmt.Print(planets)
	fmt.Print(planetsMarkII)
}
