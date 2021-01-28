package main

import "fmt"

func main() {
	var admin *string
	scolese := "Chris"
	admin = &scolese
	fmt.Println(*admin)

	bolden := "Charles"
	admin = &bolden
	fmt.Println(*admin)

	bolden = "Frank"
	fmt.Println(*admin)

	*admin = "Maj"
	fmt.Println(bolden)

	major := admin
	*major = "Maj SK"
	fmt.Println(bolden)
	// 赋值指向了相同的内存地址
	fmt.Println(admin == major)

	lightfoot := "Robert"
	admin = &lightfoot
	fmt.Println(admin == major)

	//
	charles := *major
	*major = "Charles Major"
	fmt.Println(charles)
	fmt.Println(bolden)

}
