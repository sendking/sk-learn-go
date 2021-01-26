package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	bradbury := location{-4.56, 138.44}
	curiosity := bradbury
	curiosity.long += 0.0106
	fmt.Println(bradbury, curiosity)
}
