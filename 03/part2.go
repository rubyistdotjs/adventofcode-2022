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
	lines := make([]string, 0, 3)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bag := scanner.Text()
		lines = append(lines, bag)

		if len(lines) == 3 {
			for _, char := range commonChars(lines[0], lines[1]) {
				if strings.ContainsRune(lines[2], char) {
					priorities += charPriority(char)
					break
				}
			}

			lines = lines[:0]
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

func commonChars(firstStr string, secondStr string) string {
	var commonChars string

	for _, char := range firstStr {
		if strings.ContainsRune(secondStr, char) {
			commonChars += string(char)
		}
	}

	return commonChars
}
