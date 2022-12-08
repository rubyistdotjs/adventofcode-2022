package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(rawData), "\n")
	rows := make([][]uint8, len(data))
	cols := make([][]uint8, len(data[0]))

	for ri, line := range data {
		for ci, char := range line {
			num, err := strconv.ParseUint(string(char), 10, 8)
			if err != nil {
				panic(err)
			}

			tree := uint8(num)
			rows[ri] = append(rows[ri], tree)
			cols[ci] = append(cols[ci], tree)
		}
	}

	var scenicScore uint
	visibleTrees := len(rows[0])*2 + (len(cols[0])-2)*2

	for ri := 1; ri < len(rows)-1; ri++ {
		row := rows[ri]

		for ci := 1; ci < len(row)-1; ci++ {
			tree := row[ci]
			col := cols[ci]
			score := visibility(col[ri+1:], tree) *
				reverseVisibility(row[:ci], tree) *
				reverseVisibility(col[:ri], tree) *
				visibility(row[ci+1:], tree)

			if allLessThan(row[:ci], tree) ||
				allLessThan(row[ci+1:], tree) ||
				allLessThan(col[:ri], tree) ||
				allLessThan(col[ri+1:], tree) {
				visibleTrees += 1
			}

			if score > scenicScore {
				scenicScore = score
			}
		}
	}

	fmt.Println("Part 1: ", visibleTrees)
	fmt.Println("Part 2: ", scenicScore)
}

func allLessThan(elements []uint8, num uint8) bool {
	for _, element := range elements {
		if element >= num {
			return false
		}
	}

	return true
}

func visibility(elements []uint8, num uint8) uint {
	var count uint

	for _, element := range elements {
		count += 1

		if element >= num {
			break
		}
	}

	return count
}

func reverseVisibility(elements []uint8, num uint8) uint {
	var count uint

	for i := len(elements) - 1; i >= 0; i-- {
		count += 1

		if elements[i] >= num {
			break
		}
	}

	return count
}
