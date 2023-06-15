package utils

import (
	"fmt"
)

// PrintMatrixGrid will print given 2D array/slice as a grid
// Supported types are:
//   - [][]int
func PrintMatrixGrid(matrix [][]int) {

	rows := len(matrix)
	for i := 0; i < rows; i++ {

		columns := len(matrix[i])
		for k := 0; k < columns; k++ {
			fmt.Print(matrix[i][k], " ")
		}

		fmt.Print("\n")
	}
}
