package main

import "fmt"

type celsius float64
type kelvin float64

// 函数
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

// 方法
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (c celsius) celsius() kelvin {
	return kelvin(c + 273.15)
}
func main() {
	// var k kelvin = 294.0
	// c := kelvinToCelsius(k)

	// var c celsius = 127.0
	// k := celsiusToKelvin(c)
	// fmt.Print(k, "K is ", c, "C")

	// 方法
	// var k kelvin = 294.0
	// c := k.celsius()
	// fmt.Print(c)
	var c celsius = 127.0
	k := c.celsius()
	fmt.Print(k)
}
