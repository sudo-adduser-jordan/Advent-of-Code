package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var file_name = "input.txt"

var Stacks []Stack

type Stack struct {
	crate []string
}

func (s *Stack) Peek() (element string) {
	return s.crate[len(s.crate)-1]
}

func (s *Stack) PushToBottom(element string) {
	s.crate = append([]string{element}, s.crate...)
}

func (s *Stack) PushX(elements []string) {
	s.crate = append(s.crate, elements...)
}

func (s *Stack) PopX(n int) (elements []string) {
	elements = s.crate[len(s.crate)-n : len(s.crate)]
	s.crate = s.crate[:len(s.crate)-n]
	return
}

func main() {

	ParseInput()

	ParseInstructions()

	for index := range Stacks {
		fmt.Print(Stacks[index].Peek())
	}

}

func ParseInstructions() {
	file, _ := os.Open(file_name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "m") {
			line := scanner.Text()
			array := strings.Fields(line)
			amount, _ := strconv.Atoi(array[1])
			source, _ := strconv.Atoi(array[3])
			destination, _ := strconv.Atoi(array[5])

			elements := Stacks[source-1].PopX(amount)
			Stacks[destination-1].PushX(elements)

		}
	}
}

func ParseInput() {
	file, _ := os.Open(file_name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var s string
	for scanner.Scan() {
		line := scanner.Text()
		if ContainsOnlyNumbers(line) {
			break
		}
		s = line
	}
	s = strings.ReplaceAll(s, " ", "")
	slice := strings.Split(s, "")
	total_crates := len(slice)

	Stacks = make([]Stack, total_crates)

	file2, _ := os.Open(file_name)
	defer file.Close()
	scanner2 := bufio.NewScanner(file2)

	for scanner2.Scan() {
		line := scanner2.Text()
		if strings.Contains(line, "[") {
			for index, value := range line {
				if unicode.IsUpper(value) {
					Stacks[index/4].PushToBottom(string(value))
				}
			}
		} else {
			break
		}
	}
}

func ContainsOnlyNumbers(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}
