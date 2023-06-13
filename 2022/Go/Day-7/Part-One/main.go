package main

import (
	"Go/cast"

	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/ttacon/chalk"
)

type Directory struct {
	Name   string
	Parent *Directory
	Child  map[string]*Directory
	Files  map[string]int
	Size   int
}

var input = "input.txt"

func main() {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	root := ParseInput(scanner)

	x := SumDirectories(root)

	fmt.Println(x)
}

func ParseInput(scanner *bufio.Scanner) *Directory {

	root := &Directory{
		Name:  "root",
		Child: map[string]*Directory{},
	}

	iterator := root

	for scanner.Scan() {
		line := scanner.Text()
		array := strings.Split(line, " ")
		fmt.Println()
		fmt.Println(chalk.Magenta, array, chalk.Reset)

		switch array[0] {
		// Execute Command
		case "$":
			switch array[1] {
			// Change Directory
			case "cd":
				fmt.Println(chalk.Red, array[2], chalk.Reset)
				switch array[2] {
				case "..":
					iterator = iterator.Parent
				default:
					if _, ok := iterator.Child[array[2]]; !ok {
						iterator.Child[array[2]] = &Directory{
							Name:   array[2],
							Parent: iterator,
							Child:  map[string]*Directory{},
							Files:  map[string]int{}}
					}
					iterator = iterator.Child[array[2]]
				}
			// Skip
			case "ls":
				fmt.Println(chalk.Green, "Skip", chalk.Reset)
				continue
			}
		// Add Directory
		case "dir":
			if _, ok := iterator.Child[array[1]]; !ok {
				iterator.Child[array[1]] = &Directory{
					Name:   array[1],
					Parent: iterator,
					Child:  map[string]*Directory{},
					Files:  map[string]int{}}
			}
		// Add File
		default:
			iterator.Files[array[1]] = cast.ToInt(array[0])
		}

		// Print Current Directory
		fmt.Println(ToStringDirectory(*iterator))
		for index, value := range *&iterator.Files {
			fmt.Println(ToStringFiles(index, value))
		}
	}

	FillSize(root)
	return root

}

// Fill Directory Sizes
func FillSize(iterator *Directory) int {
	size := 0

	for _, v := range iterator.Child {
		size += FillSize(v)
	}
	for _, v := range iterator.Files {
		// fmt.Println(v)
		size += v
	}

	iterator.Size = size
	fmt.Println(chalk.Red, iterator.Name)
	fmt.Println(iterator.Size, chalk.Reset)

	return size
}

// Sum Directories < X
func SumDirectories(iterator *Directory) int {
	limit := 100000
	sum := 0

	if iterator.Size <= limit {
		sum += iterator.Size
	}
	for _, v := range iterator.Child {
		sum += SumDirectories(v)
	}
	return sum
}

// Day 7
func ToStringDirectory(t Directory) lipgloss.Style {
	s := fmt.Sprintf("%+v", t)

	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Background(lipgloss.Color("0")).
		Foreground(lipgloss.Color("2"))

	return style
}

// Day 7
func ToStringFiles(s string, i int) lipgloss.Style {
	s = fmt.Sprintf(s, i)

	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Background(lipgloss.Color("0")).
		Foreground(lipgloss.Color("6"))

	return style
}
