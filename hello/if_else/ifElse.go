package main

import "fmt"

func myFitstFunction(x, y int) []int {
	lst := []int{}
	result := x * y
	lst = append(lst, result)
	return lst
}

func main() {
	if x := 1; x > 0 {
		result := myFitstFunction(x, 2)
		fmt.Println(result)
	}
}
