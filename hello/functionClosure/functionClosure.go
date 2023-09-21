package main

import "fmt"

func main() {
	callFunction := myFunc()

	fmt.Printf("%v\n", callFunction("Bulat"))

	a := 0
	b := 1
	c := 1

	f := fibonacci()
	for i := 0; i < 10; i++ {

		fmt.Printf("%v %v %v ", f(a), f(b), f(c))
		a = b + c
		b = a + c
		c = a + b
	}

}

func fibonacci() func(x int) int {
	result := func(x int) int {
		return x
	}
	return result
}

func myFunc() func(s string) string {
	result := func(s string) string {
		str := "Hi, " + s + "!"
		return str
	}
	return result
}
