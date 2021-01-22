package main

import "fmt"

func main() {
	// var red, green, blue uint8 = 0, 141, 213
	var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("%02x%02x%02x", red, green, blue)
}
