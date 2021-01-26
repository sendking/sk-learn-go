package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}
func main() {
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)

	planets := []string{"V", "M", "J"}
	newPlanets := terraform("N", planets...)
	fmt.Println(newPlanets)
}
