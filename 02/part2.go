package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var moves = map[string]string{
	"A-DRAW":  "ROCK",
	"A-WIN":   "PAPER",
	"A-LOOSE": "SCISSORS",
	"B-LOOSE": "ROCK",
	"B-DRAW":  "PAPER",
	"B-WIN":   "SCISSORS",
	"C-WIN":   "ROCK",
	"C-LOOSE": "PAPER",
	"C-DRAW":  "SCISSORS",
}

var outcomes = map[string]string{
	"X": "LOOSE",
	"Y": "DRAW",
	"Z": "WIN",
}

var scores = map[string]uint8{
	"ROCK":     1,
	"PAPER":    2,
	"SCISSORS": 3,
	"LOOSE":    0,
	"DRAW":     3,
	"WIN":      6,
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var totalScore uint
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		outcome := outcomes[round[1]]
		move := moves[round[0]+"-"+outcome]
		totalScore += uint(scores[move] + scores[outcome])
	}

	fmt.Print(totalScore)
}
