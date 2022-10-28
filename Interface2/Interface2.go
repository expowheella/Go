package main

import (
	"fmt"
)

type action interface {
	ChangeName(arg string)
	WhatItDoes()
}

type Person struct {
	name string
}

type Car struct {
	model string
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
	var bu, m action = &Person{}, &Car{}
	bu.ChangeName("BU")
	m.ChangeName("MNW")
	bu.WhatItDoes()
	// fmt.Printf("Print: %v\n", bu.ChangeName("Bill"))
	// fmt.Printf("Print: %v\n", m.ChangeName("BMW"))

}
