package main

import "fmt"

func main() {
	fmt.Println("Chapter 5: Functions")

	func() {
		fmt.Println("This is a basic anonymous function")
	}()

	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(5, 3))
}
