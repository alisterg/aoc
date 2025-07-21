package aoc_2024_10

import (
	"strconv"
	"strings"
)

const (
	TRAILHEAD = 0
	TRAILEND  = 9
)

type coordinate struct {
	x, y, val int
}

// Sum of all scores of all trailheads
// a score is the number of TRAILENDS it can reach
func solve(input string) int {
	grid := parseInput(input)
	trailheads := getTrailheads(grid)

	result := 0
	for _, trailhead := range trailheads {
		result += numTrailsForTrailhead(grid, trailhead)
	}

	return result
}

func getTrailheads(grid [][]int) []coordinate {
	starts := []coordinate{}
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == TRAILHEAD {
				starts = append(starts, coordinate{row, col, TRAILHEAD})
			}
		}
	}

	return starts
}

func numTrailsForTrailhead(grid [][]int, trailhead coordinate) int {
	reached := make(map[coordinate]bool)
	queue := []coordinate{trailhead}

	score := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if reached[curr] {
			continue
		}

		reached[curr] = true

		if curr.val == TRAILEND {
			score += 1
			continue
		}

		directions := [][]int{
			{1, 0},
			{-1, 0},
			{0, -1},
			{0, 1},
		}

		for _, dir := range directions {
			// add all valid destinations to the queue
			if next := checkDirection(grid, curr, dir); next != nil {
				queue = append(queue, *next)
			}
		}
	}

	return score
}

// returns next coordinate, if it's a valid path
func checkDirection(grid [][]int, curr coordinate, direction []int) *coordinate {
	nextX := curr.x + direction[0]
	nextY := curr.y + direction[1]

	if nextX < 0 || nextX > len(grid)-1 {
		return nil
	}
	if nextY < 0 || nextY > len(grid)-1 {
		return nil
	}
	nextVal := grid[nextX][nextY]

	if nextVal == curr.val+1 {
		return &coordinate{nextX, nextY, nextVal}
	}

	return nil
}

func parseInput(input string) [][]int {
	arr := strings.Split(input, "\n")

	var result [][]int
	for _, line := range arr {
		if len(line) == 0 {
			continue
		}

		var current []int

		for _, r := range line {
			height, _ := strconv.Atoi(string(r))
			current = append(current, height)
		}

		result = append(result, current)
	}

	return result
}
