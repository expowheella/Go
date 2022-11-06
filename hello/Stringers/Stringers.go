// A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface to print values.
package main

import (
	"fmt"
)

// create a structure Person
type Person struct {
	Name string
	Age  int
}

// create a method String
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	// create instances of Person structure
	b := Person{"Bulat Iakhin", 33}
	a := Person{"Alsu Iakhina", 34}
	fmt.Println(b, a) // here method String() is called (Stringer)
}
