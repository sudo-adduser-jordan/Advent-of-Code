package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	keys := make([]int, 0, len(calories))
	for k := range calories {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return calories[keys[i]] < calories[keys[j]]
	})

	var amount int
	for i := len(keys) - 3; i < len(keys); i++ {
		amount += calories[keys[i]]
	}

	fmt.Printf("The amount of calories that the top three elves are carrying is: %v calories", amount)
}
