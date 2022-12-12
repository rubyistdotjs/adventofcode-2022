package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinates struct {
	x, y int
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	head := coordinates{x: 0, y: 0}
	tails := make([]coordinates, 9)

	part1 := make(map[coordinates]bool)
	part1[tails[0]] = true

	part2 := make(map[coordinates]bool)
	part2[tails[8]] = true

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		direction := line[0]
		steps, _ := strconv.Atoi(line[1])

		for steps > 0 {
			head.move(direction)

			for ti := range tails {
				if ti == 0 {
					tails[ti].follow(head)
				} else {
					tails[ti].follow(tails[ti-1])
				}
			}

			part1[tails[0]] = true
			part2[tails[8]] = true
			steps--
		}
	}

	fmt.Println("Part 1: ", len(part1))
	fmt.Println("Part 2: ", len(part2))
}

func (head *coordinates) move(direction string) {
	var moveBy int

	switch direction {
	case "U", "R":
		moveBy = 1
	case "D", "L":
		moveBy = -1
	}

	switch direction {
	case "U", "D":
		head.y += moveBy
	case "R", "L":
		head.x += moveBy
	}
}

func (tail *coordinates) follow(lead coordinates) {
	if tail.around(lead) {
		return
	}

	if tail.x-lead.x == 2 {
		tail.x--
		tail.ajustY(lead)
	} else if tail.x-lead.x == -2 {
		tail.x++
		tail.ajustY(lead)
	} else if tail.y-lead.y == 2 {
		tail.y--
		tail.ajustX(lead)
	} else if tail.y-lead.y == -2 {
		tail.y++
		tail.ajustX(lead)
	}
}

func (tail *coordinates) around(lead coordinates) bool {
	edges := []int{-1, 0, 1}

	for _, xEdge := range edges {
		for _, yEdge := range edges {
			if tail.x+xEdge == lead.x && tail.y+yEdge == lead.y {
				return true
			}
		}
	}

	return false
}

func (tail *coordinates) ajustX(lead coordinates) {
	if lead.x > tail.x {
		tail.x++
	} else if lead.x < tail.x {
		tail.x--
	}
}

func (tail *coordinates) ajustY(lead coordinates) {
	if lead.y > tail.y {
		tail.y++
	} else if lead.y < tail.y {
		tail.y--
	}
}
