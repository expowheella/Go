package main

import (
	"fmt"
)

// Let's create two structures for our example
type Person struct {
	name string
}

type Car struct {
	model string
}

// Let's create an interface to access structure attributes via methods
type action interface {
	ChangeName(arg string)
	WhatItDoes()
}

func (p *Person) WhatItDoes() {
	fmt.Printf("method WhatItDoes result: %v\n", "running")
}

func (p *Car) WhatItDoes() {
	fmt.Printf("method WhatItDoes result: %v\n", "running")
}

func (p *Person) ChangeName(arg string) {
	p.name = arg
	fmt.Printf("method ChangeName result: %v\n", p.name)

}

func (c *Car) ChangeName(arg string) {
	c.model = arg
	fmt.Printf("method ChangeName result: %v\n", c.model)
}

func main() {
	// create instances of the structures connected to the interface
	var bu, m action = &Person{}, &Car{}
	bu.ChangeName("BU")
	m.ChangeName("MNW")
	bu.WhatItDoes()
	// fmt.Printf("Print: %v\n", bu.ChangeName("Bill"))
	// fmt.Printf("Print: %v\n", m.ChangeName("BMW"))

}
