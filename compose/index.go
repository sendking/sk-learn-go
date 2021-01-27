package main

import "fmt"

type report struct {
	sol         int
	temperature temperature
	location    location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (r report) average() celsius {
	return r.temperature.average()
}

func main() {
	bradbury := location{-4.45, 137.234}
	t := temperature{high: -1.0, low: -78.0}
	report1 := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report1)
	fmt.Printf("a balmy %v C\n", report1.temperature.high)

	// average
	fmt.Printf("average %v", t.average())

	// report
	report2 := report{sol: 15, temperature: t}
	fmt.Print(report2.temperature.average())
}
