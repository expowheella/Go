package main

import "fmt"

func main() {
	// initialization, condition, iteration
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)

		for j := 0; j < 5; {
			j++
			fmt.Println("j=", j)
		}
	}

}
