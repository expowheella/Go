package main

import (
	"fmt"
	"math"
)

func getArea(s shape) float64 {
	return s.area()
}

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	round := circle{5}
	quadro := rect{3, 4}
	shapes := []shape{&round, &quadro}

	for _, s := range shapes {
		fmt.Printf("getArea(s): %v\n", getArea(s))
		fmt.Printf("s.area(): %v\n", s.area())
	}

	fmt.Printf("round : %v\n", round.area())
	fmt.Printf("quadro : %v\n", quadro.area())

}
