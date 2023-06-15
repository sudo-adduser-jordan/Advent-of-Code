package main

import (
	"Go/cast"
	"Go/mathy"
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
	score := cast.ToString(FindPoints(matrix))

	fmt.Print(styles.PurpleLabel("the highest scenic score possible:"))
	fmt.Println(styles.GreenText(" " + score))
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

// Find Point Score
func FindPoints(matrix [][]int) int {
	var score_slice []int
	var score int

	rows := len(matrix)
	for current_row := 1; current_row < rows-1; current_row++ {

		columns := len(matrix[current_row])
		for current_column := 1; current_column < columns-1; current_column++ {

			point := matrix[current_row][current_column]

			top_score := CheckTop(matrix, current_row, current_column, point)
			bottom_score := CheckBottom(matrix, current_row, current_column, point)
			left_score := CheckLeft(matrix, current_row, current_column, point)
			right_score := CheckRight(matrix, current_row, current_column, point)

			score = top_score * bottom_score * left_score * right_score
			score_slice = append(score_slice, score)
		}
	}
	return mathy.MaxIntSlice(score_slice)
}

// Check distance
func CheckTop(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0
	for row := current_row - 1; row >= 0; row-- {
		count++
		if matrix[row][current_column] >= point {
			break
		}
	}
	return count
}

func CheckBottom(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0
	for row := current_row + 1; row <= len(matrix)-1; row++ {
		count++
		if matrix[row][current_column] >= point {
			break
		}
	}
	return count
}

func CheckLeft(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0
	for column := current_column - 1; column >= 0; column-- {
		count++
		if matrix[current_row][column] >= point {
			break
		}
	}
	return count
}

func CheckRight(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0

	for column := current_column + 1; column <= len(matrix)-1; column++ {
		count++
		if matrix[current_row][column] >= point {
			break
		}
	}
	return count
}
