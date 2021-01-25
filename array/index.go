package main

import "fmt"

func main() {
	var planets [8]string

	planets[0] = "SK"
	planets[1] = "LB"
	planets[2] = "YS"
	earch := planets[2]
	fmt.Println(earch, " ", len(planets), planets[3] == "")

	i := 8
	planets[i] = "P"
	pluto := planets[i]
	fmt.Print(pluto)

	var intArray [8]int
	fmt.Print(intArray[0] == 0)
}
