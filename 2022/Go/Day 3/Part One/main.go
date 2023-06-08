package main

import (
	"bufio"
	"fmt"
	"os"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var priority []int

func main() {

	file, error := os.Open("input.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		char := matchCharacter(line)
		x := numberValueOf(char)
		priority = append(priority, x)

	}
	sum := sumArray(priority)
	fmt.Printf("The sum of the matching items =  %v\n", sum)
}

func sumArray(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func numberValueOf(c rune) int {
	number := 0
	for index, value := range alphabet {
		if c == value {
			number = index + 1
			break
		}
	}
	return number
}

func matchCharacter(line string) rune {
	var char rune
	midpoint := len(line) / 2

outerloop:
	for i := 0; i < midpoint; i++ {
		charOne := rune(line[i])

		for j := midpoint; j < len(line); j++ {
			charTwo := rune(line[j])

			if charOne == charTwo {
				char = charOne
				break outerloop
			}
		}
	}
	return char
}
