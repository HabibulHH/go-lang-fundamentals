package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()

	panic("panic triggered")
}
