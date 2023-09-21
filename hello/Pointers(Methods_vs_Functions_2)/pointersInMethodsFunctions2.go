package main

import (
	"fmt"
)

type Person struct {
	name string
}

// -------------------------------------------------------------------------
// USING POINTERS TO THE STRUCTURES WITH METHODS AND FUNCTIONS BACKWARDS WAY
// -------------------------------------------------------------------------

func (p Person) ChangeName() string {
	p.name = "Billy"
	fmt.Printf("it is a method works: %v\n", p.name)
	return p.name
}

func ChangeToNick(p Person) string {
	p.name = "Expowheella"
	fmt.Printf("Change person name to nickname function: %v\n", p.name)
	return p.name
}

func main() {
	instance := Person{"Bulat"}

	fmt.Printf("instance name: %v\n", instance.name)
	fmt.Printf("change name using method: %v\n", instance.ChangeName())
	fmt.Printf("instance name: %v\n", instance.name)
	fmt.Printf("change name using function: %v\n", ChangeToNick(instance))
	fmt.Printf("instance name: %v\n", instance.name)

	fmt.Printf("--------------------------------------------------------------\n")

	// CREATE A POINTER TO THE STRUCTURE
	instance2 := &Person{"Harris"}

	fmt.Printf("instance name: %v\n", instance2.name)
	fmt.Printf("change name using method and pointer: %v\n", instance2.ChangeName())
	fmt.Printf("instance name: %v\n", instance2.name)
	// IN ORDER TO ACCESS TO THE STRUCTURE ATTRIBURE, WE USE ASTERIX(*) POINTER WITH POINTER(&)
	// TO THE STRUCTURE WHEN WE INSTANTIATE AN INSTANCE
	fmt.Printf("change name using function: %v\n", ChangeToNick(*instance2))
	fmt.Printf("instance name: %v\n", instance2.name)
}
