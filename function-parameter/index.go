package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v K\n", k)
	}
}
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func main() {
	measureTemperature(3, fakeSensor)
}
