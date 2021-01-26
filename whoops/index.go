package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earch": "Sector zz9",
		"Mars":  "Sector zz9",
	}
	planetsMarkII := planets
	planets["Earch"] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	delete(planets, "Earch")
	fmt.Println(planetsMarkII)
	//
	temperature := make(map[float64]int, 8)
	fmt.Print(temperature)
}
