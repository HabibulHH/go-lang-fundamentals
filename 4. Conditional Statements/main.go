package main

import "fmt"

func main() {
	fmt.Println("Chapter 4: Conditional Statements")

	// for {
	// 	fmt.Println("It will run forever")
	// }

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// age := 1
	// for age <= 5 {
	// 	fmt.Println(age)
	// 	age++
	// }

	// for age := 1; age <= 5; age++ {
	// 	if age == 3 {
	// 		break
	// 	}
	// 	fmt.Println(age)
	// }

	for age := 1; age <= 5; age++ {
		if age == 3 {
			continue
		}
		fmt.Println(age)
	}

	fmt.Println("Thank you")
}
