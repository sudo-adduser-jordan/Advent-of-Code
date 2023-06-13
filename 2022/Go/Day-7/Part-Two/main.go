package main

import (
	"Go/cast"
	"Go/mathy"
	"Go/styles"

	"bufio"
	"fmt"
	"math"
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
var folder_name string
var folder_size int

func main() {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	root := ParseInput(scanner)
	total_space := 70000000
	target := 30000000
	folder_to_delete := target - (total_space - root.Size)
	x := FindDirectory(root, folder_to_delete)

	fmt.Println()
	fmt.Print(styles.PurpleLabel("Folder to be deleted for update:"))
	fmt.Println(" " + cast.ToString(x))
}

// Find Directory to delete
func FindDirectory(iterator *Directory, target int) int {
	min_size := math.MaxInt64
	if iterator.Size >= target {
		min_size = mathy.MinInt(min_size, iterator.Size)
	}

	for _, Child := range iterator.Child {
		min_size = mathy.MinInt(min_size, FindDirectory(Child, target))
	}

	return min_size
}

func ParseInput(scanner *bufio.Scanner) *Directory {

	root := &Directory{
		Name:  "root",
		Child: map[string]*Directory{},
	}

	iterator := root

	for scanner.Scan() {

		// Print Current Directory
		fmt.Print(styles.RedLabel("Current directory:"))
		fmt.Print(" ")
		fmt.Println(styles.GreenStruct(*&iterator.Name))

		line := scanner.Text()

		fmt.Println()
		fmt.Print(styles.PurpleLabel("Input:"))
		fmt.Println(" " + line)

		array := strings.Split(line, " ")
		switch array[0] {
		// Execute Command
		case "$":
			switch array[1] {
			// Change Directory
			case "cd":
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
				fmt.Println(styles.GreenText("List"))
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
			fmt.Print(styles.BlueLabel("Directory Added:"))
			fmt.Println(styles.BlueText(" " + array[1]))
		// Add File
		default:
			iterator.Files[array[1]] = cast.ToInt(array[0])
			fmt.Print(styles.BlueLabel("File Added:"))
			fmt.Println(styles.BlueText(" " + array[1]))
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
	fmt.Print(styles.RedLabel(iterator.Name + " size:"))
	fmt.Println(" " + cast.ToString(iterator.Size))

	return size
}
