package main

import "fmt"

func main() {
	yesNo := "no"
	launch := (yesNo == "yes")
	fmt.Println("Ready for launch:", launch)

	boolean := true
	var booleanInt int
	if boolean == true {
		booleanInt = 1
	} else {
		booleanInt = 0
	}
	fmt.Println(booleanInt)
}
