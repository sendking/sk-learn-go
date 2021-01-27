package main

import "fmt"

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type report struct {
	sol int
	temperature
	location
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	report := report{
		sol:         15,
		location:    location{-4.5, 137.4},
		temperature: temperature{high: -1.0, low: -78.0},
	}
	fmt.Println(report.average())
	fmt.Printf("%v C\n", report.high)
	report.high = 32
	fmt.Printf("%v C\n", report.temperature.high)
}
