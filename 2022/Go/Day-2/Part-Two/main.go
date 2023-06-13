package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, error := os.Open("input.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var score int
	for scanner.Scan() {
		line := scanner.Text()

		switch line {
		case "A X":
			score += 3
		case "A Y":
			score += 4
		case "A Z":
			score += 8

		case "B X":
			score += 1
		case "B Y":
			score += 5
		case "B Z":
			score += 9

		case "C X":
			score += 2
		case "C Y":
			score += 6
		case "C Z":
			score += 7
		default:
			fmt.Println("error")
		}

	}

	fmt.Printf("Score: %v", score)
}
