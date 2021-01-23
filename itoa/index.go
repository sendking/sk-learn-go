package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown := 10
	str := "Launch is T minus " + strconv.Itoa(countdown) + " seconds"
	fmt.Println(str)
	str1 := fmt.Sprintf("Lanuch is T minus %v seconds", countdown)
	fmt.Println(str1)

	countdown1, err := strconv.Atoi("10")
	if err != nil {
		// err
	}
	fmt.Println(countdown1)
}
