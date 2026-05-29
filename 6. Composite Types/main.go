package main

import (
	"fmt"
)

func updateValue(x string) {
	x = "Hira"
}

func updateValueByPointer(x *string) {
	*x = "Hira"
}

func main() {
	fmt.Println("Chapter 6: Composite Types")

	name := "PocketSchool"

	updateValue(name)
	fmt.Println(name)

	updateValueByPointer(&name)
	fmt.Println(name)
}
