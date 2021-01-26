package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Q: 这里改成 lat, long float64 最后打印出的是 {}
	// A: 大写字母才能被导出
	type location struct {
		Lat, Long float64
	}
	curiosity := location{Lat: -4.234, Long: 173.34}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
