package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var priorities1 int
	var priorities2 int
	groupBags := make([]string, 0, 3)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bag := scanner.Text()
		groupBags = append(groupBags, bag)

		priorities1 += part1BagPriority(bag)

		if len(groupBags) == 3 {
			priorities2 += part2BagsPriority(groupBags)
			groupBags = groupBags[:0]
		}
	}

	fmt.Println("Part 1:", priorities1)
	fmt.Println("Part 2:", priorities2)
}

func part1BagPriority(bag string) int {
	halfSize := len(bag) / 2
	firstPouch := bag[0:halfSize]
	secondPouch := bag[halfSize:]

	for _, char := range firstPouch {
		if strings.ContainsRune(secondPouch, char) {
			return charPriority(char)
		}
	}

	return 0
}

func part2BagsPriority(bags []string) int {
	var commonChars string

	for _, char := range bags[0] {
		if strings.ContainsRune(bags[1], char) {
			commonChars += string(char)
		}
	}

	for _, char := range commonChars {
		if strings.ContainsRune(bags[2], char) {
			return charPriority(char)
		}
	}

	return 0
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
