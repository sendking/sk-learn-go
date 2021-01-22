package main

import "fmt"

func main() {
	piggyBank := 0.1
	for i := 0; i < 10; i++ {
		piggyBank += 0.1
	}
	fmt.Println(piggyBank)
}
