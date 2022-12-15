package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var commonDenominator int = 1

type monkey struct {
	items       []int
	operation   string
	operand     int
	divisibleBy int
	trueMonkey  int
	falseMonkey int
	inspections int
}

type monkeys []monkey

func main() {
	rawData, _ := os.ReadFile("./input.txt")
	data := strings.Split(string(rawData), "\n\n")

	var part1Monkeys monkeys
	for _, monkeyData := range data {
		lines := strings.Split(monkeyData, "\n")

		m := monkey{
			items:       parseItems(lines[1]),
			operation:   parseOperation(lines[2]),
			operand:     parseEndNumber(lines[2]),
			divisibleBy: parseEndNumber(lines[3]),
			trueMonkey:  parseEndNumber(lines[4]),
			falseMonkey: parseEndNumber(lines[5]),
		}

		commonDenominator *= m.divisibleBy
		part1Monkeys = append(part1Monkeys, m)
	}

	part2Monkeys := make(monkeys, len(part1Monkeys))
	copy(part2Monkeys, part1Monkeys)

	part1Monkeys.playRounds(20, true)
	part2Monkeys.playRounds(10000, false)

	fmt.Println(part1Monkeys.result())
	fmt.Println(part2Monkeys.result())
}

func parseItems(line string) []int {
	var items []int

	for _, item := range strings.Split(strings.Split(line, ":")[1], ",") {
		number, _ := strconv.Atoi(strings.Trim(item, " "))
		items = append(items, number)
	}

	return items
}

func parseOperation(line string) string {
	if strings.Contains(line, "+") {
		return "+"
	} else {
		return "*"
	}
}

func parseEndNumber(line string) int {
	number, _ := strconv.Atoi(strings.Trim(line[len(line)-2:], " "))
	return number
}

func (ms *monkeys) playRounds(rounds int, relieved bool) {
	for rounds > 0 {
		for mi, m := range *ms {
			for _, item := range m.items {
				if m.operation == "+" {
					if m.operand == 0 {
						item += item
					} else {
						item += m.operand
					}
				} else {
					if m.operand == 0 {
						item *= item
					} else {
						item *= m.operand
					}
				}

				if relieved {
					item = int(math.Floor(float64(item) / 3))
				} else {
					item %= commonDenominator
				}

				if item%m.divisibleBy == 0 {
					(*ms)[m.trueMonkey].items = append((*ms)[m.trueMonkey].items, item)
				} else {
					(*ms)[m.falseMonkey].items = append((*ms)[m.falseMonkey].items, item)
				}
			}

			(*ms)[mi].inspections += len(m.items)
			(*ms)[mi].items = nil
		}

		rounds--
	}
}

func (ms monkeys) result() int {
	inspections := make([]int, len(ms))

	for i, m := range ms {
		inspections[i] = m.inspections
	}

	sort.Slice(inspections, func(i int, j int) bool {
		return inspections[i] > inspections[j]
	})

	return inspections[0] * inspections[1]
}
