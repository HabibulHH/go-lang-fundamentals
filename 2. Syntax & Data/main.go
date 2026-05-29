package main

import (
	"fmt"
	"strconv"
)

// | Conversion Type | Method            |
// | --------------- | ----------------- |
// | int → float64   | `float64(i)`      |
// | float64 → int   | `int(f)`          |
// | int → string    | `strconv.Itoa(i)` |
// | string → int    | `strconv.Atoi(s)` |

func main() {
	fmt.Println("Chapter 2: Syntax & Data Containers")

	i1 := 10
	f1 := float64(i1)
	fmt.Println(f1)

	f2 := 11.22
	i2 := int(f2)
	fmt.Println(i2)

	i3 := 65
	s1 := strconv.Itoa(i3)
	fmt.Println(s1)

	s2 := "A"
	i4, _ := strconv.Atoi(s2)
	fmt.Println(i4)
}
