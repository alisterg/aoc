package aoc_2024_08

import (
	"strings"
)

type coordinate struct {
	x, y int
}

// return the number of antinodes that should be added to the input
// we add an antinode for each pair of nodes that are the same character, and
// the antinode is placed in the opposite direction of the vector between the two nodes;
// the antinode is only added if it is within the bounds of the matrix and if it is
// two times the distance between the nodes
func solve(input string) int {
	matrix := parseInput(input)
	table := make(map[rune][]coordinate)
	antinodeCoords := make(map[coordinate]bool)

	// create a hashmap of each character's coords
	// then for each we can just use math to find inserted nodes
	for y, row := range matrix {
		for x, col := range row {
			if col == '.' {
				continue
			}
			table[col] = append(table[col], coordinate{x, y})
		}
	}

	for _, coords := range table {
		for i := range coords {
			if i >= len(coords) {
				break
			}

			// use math to find where we should put antinodes
			for j := i + 1; j < len(coords); j++ {
				current, next := coords[i], coords[j]
				c1 := coordinate{2*next.x - current.x, 2*next.y - current.y}
				c2 := coordinate{2*current.x - next.x, 2*current.y - next.y}

				if isWithinBounds(c1, matrix) {
					antinodeCoords[c1] = true
				}

				if isWithinBounds(c2, matrix) {
					antinodeCoords[c2] = true
				}
			}
		}
	}

	return len(antinodeCoords)
}

func isWithinBounds(coord coordinate, matrix [][]rune) bool {
	tooLow := coord.x < 0 || coord.y < 0
	tooHigh := coord.x >= len(matrix[0]) || coord.y >= len(matrix)
	return !tooLow && !tooHigh
}

func parseInput(inp string) [][]rune {
	arr := strings.Split(inp, "\n")

	var result [][]rune
	for _, line := range arr {
		if len(line) == 0 {
			continue
		}

		var current []rune

		for _, r := range line {
			current = append(current, r)
		}

		result = append(result, current)
	}

	return result
}
