package main

import "fmt"

func main() {
	// var green uint8 = 0
	// fmt.Printf("%08b\n", green)
	// green--
	// fmt.Printf("%08b\n", green)
	// fmt.Println(green)

	var sk uint16 = 65535
	sk++
	fmt.Println(sk)
	fmt.Printf("%016b\n", sk)
}
