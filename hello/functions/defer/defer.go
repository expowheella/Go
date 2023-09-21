package main

import "fmt"

func main() {
	defer fmt.Println("2. This is defer function's printout")
	fmt.Println("1. This is standart printout")

	// using defer func we can revert the que of a function
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}
