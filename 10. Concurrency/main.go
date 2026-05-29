package main

import (
	"fmt"
	"time"
)

var counter int

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			counter++ // race condition!
		}()
	}

	time.Sleep(time.Second)

	fmt.Println(counter)
}
