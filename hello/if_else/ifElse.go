package main

import "fmt"

func myFitstFunction(x, y int) []int {
	lst := []int{}
	result := x * y
	lst = append(lst, result)
	return lst
}

func main() {

	for x := 0; x < 10; x++ {
		result := myFitstFunction(x, 2)
	}
	fmt.Println(result)
}
