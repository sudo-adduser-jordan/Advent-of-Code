package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, error := os.Open("input.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calories := make(map[int]int)

	var total int
	counter := 1
	for scanner.Scan() {

		line := scanner.Text()
		number, error := strconv.Atoi(line)
		if error != nil {
			calories[counter] = total
			total = 0
			counter += 1
		} else {
			total += number
		}
	}

	var k int
	var v int
	for key, value := range calories {
		if value > v {
			k = key
			v = value
		}
	}
	fmt.Printf("The elf that has the most snacks is elf%v with a total of %v calories\n", k, v)
}
