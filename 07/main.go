package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	folders := make(map[string]int)

	var pwd []string
	var totalSize int

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "$") {
			command := strings.Split(line, " ")

			if command[1] != "cd" {
				continue
			} else if command[2] == ".." {
				pwd = pwd[0 : len(pwd)-1]
			} else {
				pwd = append(pwd, command[2])
			}
		} else if strings.HasPrefix(line, "dir") {
			continue
		} else {
			stats := strings.Split(line, " ")
			size, _ := strconv.Atoi(stats[0])

			for i := range pwd {
				folders[strings.Join(pwd[:i+1], "_")] += size
			}

			totalSize += size
		}
	}

	var part1 int
	var part2 []string

	spaceToRecover := 30000000 - (70000000 - totalSize)

	for path, size := range folders {
		if size <= 100000 {
			part1 += size
		}

		if size >= spaceToRecover {
			part2 = append(part2, path)
		}
	}

	sort.Slice(part2, func(i, j int) bool {
		return folders[part2[i]] < folders[part2[j]]
	})

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", folders[part2[0]])
}
