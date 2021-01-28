package main

import "fmt"

func main() {
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil)
	fmt.Printf("%#v\n", v)
	var p *int
	v = p
	// 类型和值都为nil时才等于nil
	fmt.Printf("%T %v %v\n", v, v, v == nil)
	fmt.Printf("%#v\n", v)
}
