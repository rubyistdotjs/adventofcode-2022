package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var calories []int32
	var currentCalories int32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			calories = append(calories, currentCalories)
			currentCalories = 0
		} else {
			calories, err := strconv.ParseInt(line, 10, 32)
			if err != nil {
				panic(err)
			}

			currentCalories += int32(calories)
		}
	}

	sort.SliceStable(calories, func(i int, j int) bool {
		return calories[i] > calories[j]
	})

	var topThree int32
	for _, amount := range calories[0:3] {
		topThree += int32(amount)
	}

	fmt.Println(topThree)
}
