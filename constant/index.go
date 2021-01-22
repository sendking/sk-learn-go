package main

import "fmt"

func main() {
	// 如果定义了类型那么会有错误提示 const distance uint64 = 2400000000000000000000000
	// 无类型可以直接显示大数
	const distance = 2400000000000000000000000
	fmt.Println(2400000000000000000000000 / 299792 / 86400)

	const lightSpeed = 299792
	const secondsPerDay = 86400

	const days = distance / lightSpeed / secondsPerDay
	fmt.Println(days)
}
