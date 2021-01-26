package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}
	opportunity := location{lat: -1.94234, long: 324.34}
	fmt.Println(opportunity)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)

	spirit := location{-14.234, 175.23}
	fmt.Println(spirit)

	curiosity := location{-4.566, 137.4417}
	fmt.Printf("%v\n", curiosity)
	fmt.Printf("%+v", curiosity)
}
