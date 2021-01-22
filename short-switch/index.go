package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space A")
	case 1:
		fmt.Println("Space B")
	case 2:
		fmt.Println("Space C")
	default:
		fmt.Println("Space X")
	}
}
