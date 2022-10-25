package main

import (
	"fmt"
)

func main() {

	// Create a slice with make
	slice := make([]int, 0)
	fmt.Printf("This is a slice: %v \n", slice)

	// Create a slice without make() function
	slice1 := []int{}
	fmt.Printf("This is a slice: %v \n", slice1)

	// There are also ways to create a slice containig information
	slice2 := []int{0, 0, 0, 0, 0}
	slice3 := make([]int, 5)
	fmt.Printf("There are a slice2 and a slice3:\n %v\n %v \n", slice2, slice3)

	// Create leveled arrays/slices
	doubleSlice := [][]int{{}}
	tripleSlice := [][][]string{{{}}}
	fmt.Printf("This is a double slice --> %v\nand triple slice --> %v \n", doubleSlice, tripleSlice)

	// Create a map. Map is similar to dict in Python.
	map0 := make(map[string]int)
	fmt.Printf("This is an empty map: %v \n", map0)

	// Insert a value
	map0["first"] = 1
	map0["secondary"] = 2
	fmt.Printf("This is a map: %v \n", map0)

	// Show a value by key
	value0 := map0["secondary"]
	fmt.Printf("This is an element: %v \n", value0)

	// Delete key value
	delete(map0, "first")
	fmt.Printf("Delete element with key 'first' from the map0\n")
	fmt.Printf("Show the map %v\n", map0)

	// Check if key-value in the map0
	element, yes := map0["secondary"]
	element1, yes1 := map0["third"]
	fmt.Printf("The element %v in map0: %v \n", element, yes)
	fmt.Printf("The element %v in map0: %v ", element1, yes1)

}
