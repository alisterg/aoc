package aoc_2024_06

import (
	"bufio"
	"fmt"
	"os"
)

const (
	UP    = '^'
	RIGHT = '>'
	DOWN  = 'v'
	LEFT  = '<'
	BLANK = '.'
	WALL  = '#'
)

/*
The player moves forward, and turns 90deg to the right whenever it hits a '#'.
How many distinct positions will the player visit?
- It leaves the matrix after hitting one of the walls.
- It turns into a '^', '<', '>', 'v' depending on orientation.

....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
*/

func solve() int {
	input := loadInput()

	x, y := 0, 0
	dir := UP
	visited := make(map[string]bool)

	// Find the starting position
	for i, row := range input {
		for j, cell := range row {
			if cell == UP {
				x, y = i, j
				input[i][j] = '.'
			}
		}
	}

	visited[fmt.Sprintf("%d,%d", x, y)] = true

	runs := 0
	for {
		// In case there is no exit
		if runs > 100000 {
			fmt.Println("Error: limit reached")
			break
		}

		if isFinalPosition(input, x, y, dir) {
			break
		}

		if canMoveForward(input, x, y, dir) {
			x, y = moveForward(x, y, dir)
		} else {
			dir = turnRight(dir)
			x, y = moveForward(x, y, dir)
		}

		visited[fmt.Sprintf("%d,%d", x, y)] = true
	}

	return len(visited)

}

func isFinalPosition(input [][]rune, x, y int, dir rune) bool {
	switch dir {
	case UP:
		return y == 0
	case DOWN:
		return x == len(input)-1
	case LEFT:
		return x == 0
	case RIGHT:
		return y == len(input)-1
	default:
		return false
	}
}

func canMoveForward(input [][]rune, x, y int, dir rune) bool {
	switch dir {
	case UP:
		return input[x-1][y] != WALL
	case DOWN:
		return input[x+1][y] != WALL
	case LEFT:
		return input[x][y-1] != WALL
	case RIGHT:
		return input[x][y+1] != WALL
	default:
		return false
	}
}

func moveForward(x, y int, dir rune) (int, int) {
	switch dir {
	case UP:
		return x - 1, y
	case DOWN:
		return x + 1, y
	case LEFT:
		return x, y - 1
	case RIGHT:
		return x, y + 1
	default:
		return x, y
	}
}

func turnRight(dir rune) rune {
	switch dir {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	default:
		return UP
	}
}

// loadInput loads the file as a [][]rune
func loadInput() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		var current []rune

		for _, r := range line {
			current = append(current, r)
		}

		result = append(result, current)
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return result
}
