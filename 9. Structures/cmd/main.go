package main

import (
	"fmt"
	"structures/utils"

	"github.com/google/uuid"
)

// | Feature           | `main` Package             | Other Packages              |
// | ----------------- | -------------------------- | --------------------------- |
// | Purpose           | Entry point of program     | Reusable code/library       |
// | Usage             | Starts the program         | Used/imported by others     |
// | Execution         | Can run directly           | Cannot run directly         |
// | Required function | Must have `func main()`    | Nothing special             |

func main() {
	fmt.Println("Chapter 9: Structures")

	utils.Add(2, 4)

	id := uuid.New()
	fmt.Println(id)
}
