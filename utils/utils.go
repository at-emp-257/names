package utils

import "fmt"

func PrintAll(header string, list []string) {
	fmt.Println("----------------------------------")
	result := header + "\n\n"

	for i := 0; i < len(list); i++ {
		line := fmt.Sprintf("%d. %s\n", i + 1, list[i])
		result += line
	}

	fmt.Println(result)
	fmt.Println("----------------------------------")
}