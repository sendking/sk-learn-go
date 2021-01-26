package main

import "fmt"

func main() {
	type location struct {
		lat  float64
		long float64
	}
	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.234234

	var opportunity location
	opportunity.lat = -1.234
	opportunity.long = 354.43543
	fmt.Println(spirit, opportunity)
}
