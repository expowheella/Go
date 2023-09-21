package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	for i := 0; i < 10; i++ {
		cache := z
		z -= (z*z - x) / (2 * z)
		if (cache - z) >= 0.001 {
			return x, ErrNegativeSqrt(x) // ErrNegativeSqrt(x) /* Error() is called here */
		} else {
			return x, nil
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(ErrNegativeSqrt(-10))
}
