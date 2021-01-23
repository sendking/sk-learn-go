package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"
	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
	zimu := "¿abcdefghijklmnopqrstuvwxyz"
	// 字节
	fmt.Printf("%v\n", len(zimu))
	// 符文
	fmt.Printf("%v\n", utf8.RuneCountInString(zimu))
}
