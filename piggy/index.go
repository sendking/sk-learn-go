package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var total float64

	for total <= 20.00 {
		rand.Seed(time.Now().Unix())
		money := 0.05
		switch i := rand.Intn(2); i {
		case 0:
			money = 0.05
		case 1:
			money = 0.10
		case 2:
			money = 0.25
		}
		total += money
	}
	fmt.Printf("total:%.2f", total)
}
