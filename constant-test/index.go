package main

import "fmt"

func main() {
	const distance = 236e15
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const days = distance / lightSpeed / secondsPerDay
	fmt.Println(days)
}
