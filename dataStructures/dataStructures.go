package main

import (
	"fmt"
)

func main() {

	// slice
	// динамический непрерывный массив
	slice := []int{}

	slice2 := make([]int, 0)

	fmt.Println(slice)
	fmt.Println(slice2)

	//
	//стек (LIFO)
	stack := make([]int, 0)

	// push a value
	slice = []int{1, 2, 3}
	value := 100
	stack = append(stack, value)
	fmt.Println(stack)

	// push an array
	stack = appendArray(slice, stack)
	fmt.Println(stack)

	// pop a value
	stack, value = stack[:len(stack)-1], stack[len(stack)-1]
	fmt.Printf("pop %v from stack %v", value, stack)

	//
	//
	// Queue (Очередь)FIFO

}

func appendArray(array, s []int) (stack []int) {
	for i := 0; i < len(array); i++ {
		s = append(s, array[i])
	}
	return s
}
