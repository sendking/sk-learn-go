package main

import "fmt"

func main() {
	fmt.Printf("%v is a %[1]T\n", "listeral string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)
}
