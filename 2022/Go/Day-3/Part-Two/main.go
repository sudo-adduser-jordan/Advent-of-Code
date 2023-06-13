package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
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

	len, error := LineCounter(file)
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	x := (len + 1) / 3
	for i := 0; i < x; i++ {

		scanner.Scan()
		lineOne := scanner.Text()
		scanner.Scan()
		lineTwo := scanner.Text()
		scanner.Scan()
		lineThree := scanner.Text()

		c := matchCharacter(lineOne, lineTwo, lineThree)

		n := numberValueOf(c)
		priority = append(priority, n)
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
	fmt.Printf("number value: %v\n", number)
	return number
}

func matchCharacter(lineOne string, lineTwo string, lineThree string) rune {
	var c rune
	for i := 0; i < len(lineOne); i++ {
		c = rune(lineOne[i])
		if strings.ContainsRune(lineTwo, c) && strings.ContainsRune(lineThree, c) {
			return c
		}
	}
	return c
}

func LineCounter(r io.Reader) (int, error) {

	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}
