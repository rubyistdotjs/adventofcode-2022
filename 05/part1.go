package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	stacks := make([][]string, 9)
	r, err := regexp.Compile("(\\d+).*(\\d+).*(\\d+)")
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		} else if isCrateLine(line) {
			for i, crate := range parseCrates(scanner.Text()) {
				if crate != "" {
					stacks[i] = append(stacks[i], crate)
				}
			}
		} else if isInstructionLine(line) {
			matches := r.FindStringSubmatch(line)
			from := toInt(matches[2]) - 1
			to := toInt(matches[3]) - 1

			for i := 0; i < toInt(matches[1]); i++ {
				stacks[to] = append([]string{stacks[from][0]}, stacks[to]...)
				stacks[from] = stacks[from][1:]
			}
		}
	}

	var result string
	for _, stack := range stacks {
		result += stack[0]
	}

	fmt.Print(result)
}

func toInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(i)
}

func isInstructionLine(line string) bool {
	return strings.HasPrefix(line, "move")
}

func isCrateLine(line string) bool {
	firstChar := string(line[0])
	return (firstChar == " " || firstChar == "[") && !strings.Contains(line, "2")
}

func parseCrates(row string) []string {
	rows := int(float64(len(row)) / 3.88)
	crates := make([]string, 9)

	for i := 0; i < rows; i++ {
		crate := string(row[i*4+1])

		if crate != " " {
			crates[i] = crate
		}
	}

	return crates
}
