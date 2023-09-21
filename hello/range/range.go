// Exercise Slices
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// create a 2D emty array with length == dy
	matrix := make([][]uint8, dy)

	for i := range matrix { // where i is index
		matrix[i] = make([]uint8, dx) // create an array named matrix[i] = [   ]
		// matrix[0] = [   ], matrix[1] = [   ], matrix[2] = [   ] ... , matrix[n] = [   ]
		for j := range matrix[i] { // where j is index of matrix[i]
			// for example, matrix[0] = [j0, j1, j2, j3, ... jn], where j is index
			matrix[i][j] = uint8(i ^ j) // here we have a template for 2D array which was instantiated in the beginnig as matrix
			// e.x.: matrix[0][1] = [[0, 1, 2, ... n], [0, 1, 2, ... n], [0, 1, 2, ... n], ...]
		}
	}
	return matrix
}

func main() {
	pic.Show(Pic)
}
