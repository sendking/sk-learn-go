package main

import "fmt"

type location struct {
	lat, long float64
}

// 因为fmt包申明了Stringer
func (l location) String() string {
	return fmt.Sprintf("%v %v", l.lat, l.long)
}

func main() {
	curiosity := location{-4.58, 137.4417}
	fmt.Println(curiosity)
}
