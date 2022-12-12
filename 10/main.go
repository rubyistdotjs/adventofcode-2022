package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type program struct {
	ticker, x, strength int
	screen              [][]string
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	p := program{ticker: 1, x: 1, screen: make([][]string, 6)}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]

		if command == "addx" {
			value, _ := strconv.Atoi(line[1])
			p.tick(2)
			p.x += value
		} else {
			p.tick(1)
		}
	}

	fmt.Println("Part 1: ", p.strength)
	fmt.Println("Part 2:")
	for _, row := range p.screen {
		fmt.Println(row)
	}
}

func (p *program) tick(times int) {
	for times > 0 {
		p.addStrength()
		p.draw()
		p.ticker++
		times--
	}
}

func (p *program) addStrength() {
	if (p.ticker+20)%40 == 0 {
		p.strength += p.x * p.ticker
	}
}

func (p *program) draw() {
	crtPosition := p.ticker - 1
	row := crtPosition / 40
	spritePosition := (crtPosition - (40 * row)) - p.x

	if spritePosition >= -1 && spritePosition <= 1 {
		p.screen[row] = append(p.screen[row], "#")
	} else {
		p.screen[row] = append(p.screen[row], ".")
	}
}
