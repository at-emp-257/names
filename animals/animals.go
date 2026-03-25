package animals

import (
	"fmt"
)

var people = []string{
	"anaconda",
	"bear",
	"cat",
}

func PrintAll() {
	result := "Printing animals\n\n"

	for i := 0; i < len(people); i++ {
		line := fmt.Sprintf("%d. %s\n", i + 1, people[i])
		result += line
	}

	fmt.Print(result)
}

