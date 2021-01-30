package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	count := 0

	for r := 0; r < 100; r++ {
		go func() {
			mutex.Lock()
			count++
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("the count is : ", count)
}
