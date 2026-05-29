package main

import (
	"fmt"
)

// func printInts(numbers []int) {
// 	for _, v := range numbers {
// 		fmt.Println(v)
// 	}
// }

// func printStrings(texts []string) {
// 	for _, v := range texts {
// 		fmt.Println(v)
// 	}
// }

// func printSlices(items []any) {
// 	for _, v := range items {
// 		fmt.Println(v)
// 	}
// }

func printSlices[T any](items []T) {
	for _, v := range items {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("Chapter 7: Behavior")

	numbers := []int{1, 2, 3}
	texts := []string{"a", "b", "c"}

	fmt.Println(numbers)
	fmt.Println(texts)
}
