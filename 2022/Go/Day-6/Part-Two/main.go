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
var counter = 15

func main() {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		signal = scanner.Text()
	}

	marker = signal[0:14]
	signal = signal[14:]

	if ValidateMarker(marker) {
		fmt.Printf("counter: %v \n", counter)
	}

	for _, value := range signal {
		marker = marker[1:] + string(value)

		if ValidateMarker(marker) {
			fmt.Printf("counter: %v \n", counter)
			break
		}

		counter++
	}

}

func ValidateMarker(s string) bool {

	for i := 0; i < len(s); i++ {
		if strings.Contains(
			strings.Replace(s, string(s[i]), "", 1),
			string(s[i])) {
			return false
		}
	}
	return true
}
