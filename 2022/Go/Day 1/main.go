package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// get file
	file, error := os.Open("calories.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	// init scanner
	scanner := bufio.NewScanner(file)

	//init map
	calories := make(map[int]int)

	// loop file by line
	var total int
	counter := 1
	for scanner.Scan() {

		line := scanner.Text()
		number, error := strconv.Atoi(line)
		if error != nil {
			// fmt.Println(error)
		}

		if number == 0 {
			calories[counter] = total
			total = 0
			counter += 1
		}

		if number != 0 {
			total += number
		}

	}
	var k int
	var v int
	// does not check for matching values
	for key, value := range calories {
		if value > v {
			k = key
			v = value
		}
	}
	fmt.Printf("The elf that has the most snacks is elf%v with a total of %v calories\n", k, v)
}
