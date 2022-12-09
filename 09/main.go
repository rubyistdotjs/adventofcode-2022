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
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	hCoord := coordinates{x: 0, y: 0}
	tCoord := coordinates{x: 0, y: 0}
	positions := make(map[coordinates]bool)
	positions[tCoord] = true

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		count, err := strconv.ParseInt(line[1], 10, 8)
		if err != nil {
			panic(err)
		}

		for i := 0; i < int(count); i++ {
			var move int

			switch line[0] {
			case "U", "R":
				move = 1
			case "D", "L":
				move = -1
			}

			switch line[0] {
			case "U", "D":
				hCoord.y += move

				if !tCoord.around(hCoord) {
					tCoord.y += move

					if tCoord.x != hCoord.x {
						tCoord.x = hCoord.x
					}
				}
			case "R", "L":
				hCoord.x += move

				if !tCoord.around(hCoord) {
					tCoord.x += move

					if tCoord.y != hCoord.y {
						tCoord.y = hCoord.y
					}
				}
			}

			positions[tCoord] = true
		}
	}

	fmt.Println(len(positions))
}

func (t coordinates) around(h coordinates) bool {
	box := []int{-1, 0, 1}

	for _, xEdge := range box {
		for _, yEdge := range box {
			if t.x+xEdge == h.x && t.y+yEdge == h.y {
				return true
			}
		}
	}

	return false
}
