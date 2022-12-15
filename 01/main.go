package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var calories []int
	var currentCalories int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			calories = append(calories, currentCalories)
			currentCalories = 0
		} else {
			calories, _ := strconv.Atoi(line)
			currentCalories += calories
		}
	}

	sort.Slice(calories, func(i int, j int) bool {
		return calories[i] > calories[j]
	})

	var part2 int
	for _, amount := range calories[0:3] {
		part2 += amount
	}

	fmt.Println("Part1:", calories[0])
	fmt.Println("Part2:", part2)
}
