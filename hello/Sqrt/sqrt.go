package main

import (
	"fmt"
)

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

func main() {
	fmt.Println(Sqrt(2))
}
