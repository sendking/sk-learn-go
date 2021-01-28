package main

import "fmt"

func main() {
	var nowhere *int
	if nowhere != nil {
		fmt.Println(*nowhere)
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	panic("I forgot my tower")
}
