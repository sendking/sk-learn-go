package main

import "fmt"

func main() {
	value := "1"
	var boolean bool
	switch value {
	case "true", "yes", "1":
		boolean = true
	case "false", "no", "0":
		boolean = false
	}
	fmt.Println(boolean)
}
