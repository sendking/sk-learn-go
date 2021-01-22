package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	money := 5.00
	total := 0.00
	i := 0
	for total < 2000 {
		rand.Seed(time.Now().Unix())
		switch i := rand.Intn(2); i {
		case 0:
			money = 5.00
		case 1:
			money = 10.00
		case 2:
			money = 25.00
		}
		total += money
		i++
		fmt.Printf("total:$%.2f,count:%v\n", total/100, i)
	}
}
