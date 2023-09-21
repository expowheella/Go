package main

import "fmt"

type Coordinate struct {
	X, Y float64
}

// we can initialise a variable with name m and type map of structure with string fields
// var m map[string]Coordinate
// but it does not needed because further we create a map of structure

// create a nil map (empty map)
var m = make(map[string]Coordinate)

func main() {
	fmt.Println(m)

	m["Bell Labs"] = Coordinate{40.68433, -74.39967}
	fmt.Printf("Coordinates of Bell Labs: %v \n", m["Bell Labs"])

	testMap := make(map[string]int)
	fmt.Printf("Print testMap %v", testMap)
}
