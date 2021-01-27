package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}
type location struct {
	lat, long float64
}

func (c coordinate) decemal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s + 3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decemal(), long.decemal()}
}
func main() {
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.decemal(), long.decemal())
	curiosity := newLocation(lat, long)
	fmt.Println(curiosity)
}
