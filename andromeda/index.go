package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	secondsPerDay1 := new(big.Int)
	secondsPerDay1.SetString("86400", 10)
	distance := new(big.Int)
	distance.SetString("2400000000000000000", 10)
	fmt.Println(distance)
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println((days))
}
