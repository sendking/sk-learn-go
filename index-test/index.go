package main

import "fmt"

func main() {
	message := "shalom"
	for i := 0; i < len(message); i++ {
		// fmt.Println(message[i])
		fmt.Printf("%c\n", message[i])
	}
}
