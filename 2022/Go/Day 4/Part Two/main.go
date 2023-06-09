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

		rangeOneStart, error := strconv.Atoi(arrayTwo[0])
		if error != nil {
			panic(error)
		}
		rangeOneEnd, error := strconv.Atoi(arrayTwo[1])
		if error != nil {
			panic(error)
		}
		rangeTwoStart, error := strconv.Atoi(arrayOne[0])
		if error != nil {
			panic(error)
		}
		rangeTwoEnd, error := strconv.Atoi(arrayOne[1])
		if error != nil {
			panic(error)
		}

		if rangeOneStart <= rangeTwoEnd && rangeOneEnd >= rangeTwoStart || rangeTwoStart <= rangeOneEnd && rangeTwoEnd >= rangeOneStart {
			total++
		}
	}

	fmt.Printf("Amount of assignment pairs that the ranges overlap: %v \n", total)
}
