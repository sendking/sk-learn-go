package main

import "fmt"

func main() {
	temperature := map[string]int{
		"E":    15,
		"M":    -65,
		"Moon": 0,
	}
	temp := temperature["E"]
	fmt.Printf("On %v C.\n", temp)
	temperature["E"] = 16
	temperature["V"] = 464
	fmt.Println(temperature)
	moon := temperature["Moon"]
	fmt.Println(moon)
	if moon, ok := temperature["Moon"]; ok {
		fmt.Print("I'm ", moon)
	} else {
		fmt.Print("Where is the moon?")
	}
}
