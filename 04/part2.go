package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var overlaps uint16

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sectors := parseSectors(scanner.Text())

		if sectorOverlap(sectors[0], sectors[1]) || sectorOverlap(sectors[1], sectors[0]) {
			overlaps += 1
		}
	}

	fmt.Print(overlaps)
}

func parseSectors(line string) [][]uint8 {
	sectors := make([][]uint8, 0, 2)

	for _, sector := range strings.Split(line, ",") {
		sections := make([]uint8, 0, 2)

		for _, rawSection := range strings.Split(sector, "-") {
			section, err := strconv.ParseUint(rawSection, 10, 8)
			if err != nil {
				panic(err)
			}

			sections = append(sections, uint8(section))
		}

		sectors = append(sectors, sections)
	}

	return sectors
}

func sectorOverlap(firstSector []uint8, secondSector []uint8) bool {
	return (firstSector[0] >= secondSector[0] && firstSector[0] <= secondSector[1]) || (firstSector[1] >= secondSector[0] && firstSector[1] <= secondSector[1])
}
