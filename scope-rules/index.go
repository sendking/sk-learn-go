package main

import (
	"fmt"
	"math/rand"
	"time"
)

const era = "AD"

func main() {
	rand.Seed(time.Now().Unix())
	year := rand.Intn(2018) + 1
	month := rand.Intn(12) + 1
	dayInMonth := 31
	switch month {
	case 2:
		if year%4 == 0 {
			dayInMonth = 29
		} else {
			dayInMonth = 28
		}
	case 4, 6, 9, 11:
		dayInMonth = 30
	default:
		dayInMonth = 31
	}
	for i := 0; i < 10; i++ {
		day := rand.Intn(dayInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}
