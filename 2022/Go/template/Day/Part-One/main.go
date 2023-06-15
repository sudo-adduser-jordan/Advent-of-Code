package main

import (
	"bufio"
	"fmt"
	"os"
)

var input = "input.txt"

func main() {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Print(line)
	}
}
