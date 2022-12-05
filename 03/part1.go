package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var priorities int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bag := scanner.Text()
		halfSize := len(bag) / 2
		firstPouch := bag[0:halfSize]
		secondPouch := bag[halfSize:]

		for _, char := range firstPouch {
			if strings.ContainsRune(secondPouch, char) {
				priorities += charPriority(char)
				break
			}
		}
	}

	fmt.Print(priorities)
}

func charPriority(char rune) int {
	priority := int(char)

	if priority > 96 {
		priority -= 96 // ASCII a to z starts at 97
	} else {
		priority -= 38 // ASCII A to Z starts at 65 but we want to start at 27
	}

	return priority
}
