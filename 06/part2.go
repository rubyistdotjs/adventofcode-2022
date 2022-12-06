package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	rawData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	data := string(rawData)

	for i, char := range data {
		chars := string(char)

		for j := 1; j < 14; j++ {
			nextChar := string(data[i+j])

			if strings.Contains(chars, nextChar) {
				break
			} else {
				chars += nextChar
			}
		}

		if len(chars) == 14 {
			fmt.Print(i + 14)
			break
		}
	}
}
