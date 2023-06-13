package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var total int

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		array := strings.Split(line, ",")

		stringOne := array[0]
		stringTwo := array[1]

		arrayOne := strings.Split(stringOne, "-")
		arrayTwo := strings.Split(stringTwo, "-")

		rangeOneStart, _ := strconv.Atoi(arrayTwo[0])
		rangeOneEnd, _ := strconv.Atoi(arrayTwo[1])
		rangeTwoStart, _ := strconv.Atoi(arrayOne[0])
		rangeTwoEnd, _ := strconv.Atoi(arrayOne[1])

		if rangeTwoStart >= rangeOneStart && rangeTwoEnd <= rangeOneEnd || rangeOneStart >= rangeTwoStart && rangeOneEnd <= rangeTwoEnd {
			total++
		}

	}

	fmt.Printf("Amount of assignment pairs that one range fully contains the other: %v \n", total)
}
