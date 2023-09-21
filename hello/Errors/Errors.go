package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)

		// Exercise: Errors
		fmt.Println(Sqrt(2))
		fmt.Println(Sqrt(-2))
		//
		//
		//
	}
}

// Exercise: Errors
func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		cache := z
		z -= (z*z - x) / (2 * z)

		fmt.Println("start", z, cache)

		if (cache - z) >= 0.001 {
			fmt.Println("if", z, cache)
			return z
		}

		fmt.Println("for", z, cache)

	}

	return z
}

//
//
//
