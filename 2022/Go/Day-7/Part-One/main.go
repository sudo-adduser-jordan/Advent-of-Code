package main

import (
	"Go/cast"
	"Go/styles"

	"bufio"
	"fmt"
	"os"
	"strings"
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

	x := cast.ToString(SumDirectories(root))

	fmt.Println()
	fmt.Println(styles.GreenText("The sum of the directories with sizes <100000 is: " + x))
}

func ParseInput(scanner *bufio.Scanner) *Directory {

	root := &Directory{
		Name:  "root",
		Child: map[string]*Directory{},
	}

	iterator := root

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println()
		fmt.Println(styles.PurpleLabel(line))

		array := strings.Split(line, " ")
		switch array[0] {
		// Execute Command
		case "$":
			switch array[1] {
			// Change Directory
			case "cd":
				fmt.Println(styles.RedText(array[2]))
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
				fmt.Println(styles.GreenText("Skip"))
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
		fmt.Println(styles.GreenStruct(*iterator))
		for index, value := range *&iterator.Files {
			s := index + " " + cast.ToString(value)
			fmt.Println(styles.BlueText(s))
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
		size += v
	}

	iterator.Size = size

	fmt.Println()
	fmt.Println(styles.RedLabel(iterator.Name))
	fmt.Println(iterator.Size)

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
