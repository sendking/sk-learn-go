package main

import "fmt"

func main() {
	type location struct {
		name string
		lat  float64
		long float64
	}
	locations := []location{
		{name: "B", lat: -4.5234, long: 137.234},
		{name: "C", lat: 4.54, long: -137.234},
	}
	fmt.Println(locations)
}
