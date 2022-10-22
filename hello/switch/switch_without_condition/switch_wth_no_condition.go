package main

import (
	"fmt"
)

func main() {

	a := 3
	b := 3

	switch {
	case a < b:
		fmt.Println("a < b")
	case a == b:
		fmt.Print("a == b")
	default:
		fmt.Println("a > b")
	}

}
