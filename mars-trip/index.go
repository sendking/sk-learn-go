package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Printf("%-15v%10v%10v%10v\n", "太空航行公司", "飞行天数", "飞行类型", "价格")
	i := 0
	for i < 10 {
		company := "Space Adventures"
		speed := rand.Intn(30) + 1
		price := rand.Intn(5000) / 100
		tripTime := (62100000 / speed) / (60 * 60 * 24)
		tripType := "单程"
		switch tripTypeIndex := rand.Intn(2); tripTypeIndex {
		case 0:
			tripType = "单程"
		case 1:
			tripType = "往返"
		}
		switch companyIndex := rand.Intn(2); companyIndex {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		default:
			company = "Space Adventures"
		}

		fmt.Printf("%-15v%10v%10v%10v\n", company, tripTime, tripType, price)
		i++
	}
}
