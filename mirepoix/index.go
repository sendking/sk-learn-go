package main

import "fmt"

func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "carrot")
}
func main() {
	soup := mirepoix(nil)
	fmt.Println(soup)
}
