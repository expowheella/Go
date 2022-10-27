package main

import "fmt"

func main() {
	callFunction := myFunc()

	fmt.Print(callFunction("Bulat"))
}

func myFunc() func(s string) string {
	result := func(s string) string {
		str := "Hi, " + s + "!"
		return str
	}
	return result
}
