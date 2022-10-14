package main

import (
	"fmt"
)

const (
	Big = 1 << 100
)

// declare var a in a package scope
var a int = 1

func main() {

	fmt.Println(a)
	// print value type value type

	// var (
	// 	b int     = 2
	// 	c float64 = 3.45
	// )
	// fmt.Println(b, c)
	// fmt.Printf("%v, %v, %T, %T\n", a, b, a, b)

	// daclare var a in a func scope
	a := 2
	fmt.Println("print var a here:", a)

	b := float64(42)
	fmt.Println(b)

	// var d string = strconv.Itoa(b)
	// fmt.Printf("%v, %T\n", d, d)

	// f, err := strconv.Atoi(d)
	// fmt.Printf("%v, %T, %v\n", f, f, err)
}
