package main

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}
func main() {
	planets := []string{
		"M", "V", "E", "M", "J", "S", "U", "N", "P",
	}
	reclassify(&planets)
	fmt.Println(planets)
}
