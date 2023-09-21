package main

import (
	"fmt"
)

var input string = "ababa"

func main() {
	firstCount := ""

	for i := 0; i < len(input); i++ {
		firstCount := input[i]
		// fmt.Printf("%c %c\n", firstCount, input[i])
		totalCount := 0

		for j := 0; j < len(input); j++ {

			if firstCount == input[j] {
				// fmt.Printf("%c %c\n", firstCount, input[j])
				totalCount += 1

			}

		}
		fmt.Printf("%c %v\n", firstCount, totalCount)
	}
	// fmt.Printf("%c", input[i])
	fmt.Println("for loop is finished", firstCount)
}
