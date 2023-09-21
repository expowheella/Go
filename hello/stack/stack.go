package main

import (
	"fmt"
)

// Create an emty stack
var stack []int

func main() {
	// this statement will be printed in the end
	defer fmt.Println("The program is finished!")

	// Print out an empty stack
	fmt.Printf("show empty stack: %v\n", stack)

	// append element into stack
	fmt.Println(appendElement(stack, 2))

	// append array into stack
	array := []int{1, 2, 3, 4}
	fmt.Printf("length of the array %v\n", len(array))
	fmt.Println(appendArray(array))

	// pop item from stack
	fmt.Println(popFromArray(array))
}

func appendElement(array []int, value int) []int {
	array = append(array, value)
	return array
}

// create a function to append array (stack)
func appendArray(someArray []int) []int {
	for i := 0; i < len(someArray); i++ {
		stack = append(stack, someArray[i])
	}
	return stack
}

// pop value from stack
func popFromArray(array []int) (int, []int) {
	value := array[len(array)-1]
	fmt.Printf("Print popping value %v\n", value)
	newArray := array[:len(array)-1]
	return value, newArray
	// fmt.Printf("value %v is popped from the array \n", value)
}
