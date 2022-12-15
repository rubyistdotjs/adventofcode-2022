package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var moves = map[string]string{
	"A": "ROCK",
	"B": "PAPER",
	"C": "SCISSORS",
	"X": "ROCK",
	"Y": "PAPER",
	"Z": "SCISSORS",
}

var outcomes = map[string]string{
	"ROCK-PAPER":     "WIN",
	"ROCK-SCISSORS":  "LOOSE",
	"PAPER-ROCK":     "LOOSE",
	"PAPER-SCISSORS": "WIN",
	"SCISSORS-ROCK":  "WIN",
	"SCISSORS-PAPER": "LOOSE",
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
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var totalScore uint

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rawRound := strings.Split(line, " ")
		opponentMove := moves[rawRound[0]]
		ownMove := moves[rawRound[1]]

		var outcome string
		if opponentMove == ownMove {
			outcome = "DRAW"
		} else {
			outcome = outcomes[opponentMove+"-"+ownMove]
		}

		totalScore += uint(scores[ownMove] + scores[outcome])
	}

	fmt.Print(totalScore)
}
