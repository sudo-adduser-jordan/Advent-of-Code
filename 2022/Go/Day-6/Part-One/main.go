package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input = "input.txt"

var signal string
var marker string
var counter = 5

func main() {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		signal = scanner.Text()
	}

	marker = signal[0:4]
	signal = signal[4:]

	if ValidateMarker(marker) {
		fmt.Printf("marker: %v \n", marker)
		fmt.Printf("counter: %v \n", counter)
	}

	for _, value := range signal {
		marker = marker[1:] + string(value)

		if ValidateMarker(marker) {
			fmt.Printf("marker: %v \n", marker)
			fmt.Printf("counter: %v \n", counter)
			break
		}
		counter++
	}

}

func ValidateMarker(s string) bool {

	if strings.Contains(string(s[1:4]), string(s[0])) ||
		strings.Contains(string(s[0])+s[2:4], string(s[1])) ||
		strings.Contains(string(s[0:2])+s[3:4], string(s[2])) {
		return false
	}

	return true
}
