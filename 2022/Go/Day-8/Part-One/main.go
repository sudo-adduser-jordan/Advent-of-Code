package main

import (
	"Go/cast"
	"Go/styles"
	"bufio"
	"fmt"
	"os"
)

var input = "input.txt"

func main() {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	matrix := ParseGrid(scanner)
	points := cast.ToString(FindPoints(matrix))

	fmt.Print(styles.PurpleLabel("The amount of trees visiable:"))
	fmt.Println(styles.GreenText(" " + points))

}

// Parse Input
func ParseGrid(scanner *bufio.Scanner) [][]int {
	var matrix [][]int
	for scanner.Scan() {
		row := cast.SliceToInt(scanner.Bytes())
		matrix = append(matrix, row)
	}
	return matrix
}

// Find Points the are not visible from outside grid
func FindPoints(matrix [][]int) int {
	count := (len(matrix) * 4) - 4

	rows := len(matrix)
	for current_row := 1; current_row < rows-1; current_row++ {

		columns := len(matrix[current_row])
		for current_column := 1; current_column < columns-1; current_column++ {

			point := matrix[current_row][current_column]

			// Check if visiable
			if CheckTop(matrix, current_row, current_column, point) ||
				CheckBottom(matrix, current_row, current_column, point) ||
				CheckRight(matrix, current_row, current_column, point) ||
				CheckLeft(matrix, current_row, current_column, point) {
				count++
			}
		}
	}
	return count
}

// Check if Visiable
func CheckTop(matrix [][]int, current_row int, current_column int, point int) bool {

	var max int
	for row := 0; row < current_row; row++ {
		if max < matrix[row][current_column] {
			max = matrix[row][current_column]
		}
	}

	if point > max {
		return true
	}

	return false
}

func CheckBottom(matrix [][]int, current_row int, current_column int, point int) bool {

	var max int
	for row := len(matrix) - 1; row > current_row; row-- {
		if max < matrix[row][current_column] {
			max = matrix[row][current_column]
		}
	}

	if point > max {
		return true
	}

	return false
}

func CheckLeft(matrix [][]int, current_row int, current_column int, point int) bool {

	var max int
	for column := 0; column < current_column; column++ {
		if max < matrix[current_row][column] {
			max = matrix[current_row][column]
		}
	}

	if point > max {
		return true
	}

	return false
}

func CheckRight(matrix [][]int, current_row int, current_column int, point int) bool {

	var max int
	for column := len(matrix[current_row]) - 1; column > current_column; column-- {
		if max < matrix[current_row][column] {
			max = matrix[current_row][column]
		}
	}

	if point > max {
		return true

	}
	return false
}
