package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance uint64 = 41.3e12
	// 	uint64类型溢出 var distance uint64 = 24e18
	fmt.Println("Alpha is", distance, "km away")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")
}
