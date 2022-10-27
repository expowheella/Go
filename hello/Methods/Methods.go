package main

import (
	"fmt"
)

type Person struct {
	name, surname string
}

// create a method - fullName() for Person structure which returns string
func (receiver Person) fullName() string {
	return (receiver.name + " " + receiver.surname)
}

func main() {
	// create an instance of structure (class in Python)
	human := Person{"Bulat", "Iakhin"}
	human.pointerMethod()
	// call method fullName() on human
	fmt.Println(human.fullName())

}

// Pointer receiver does not return anything
// It is just influence a structure
func (p *Person) pointerMethod() {
	p.name = "Mr. " + p.name
}
