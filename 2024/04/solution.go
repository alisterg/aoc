package aoc_2024_04

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

// solve takes an input 2d matrix and
// performs a wordsearch for 'XMAS'.
// A valid word position could be horizontal,
// vertical, diagonal, backwards, or overlapping
// other words.
// It returns the number of XMAS found.
func solve() int {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	grid       := loadInputGrid()
	searchTerm := "XMAS"
	numRows    := len(grid)
	numCols    := len(grid[0])
	total      := 0

	validPosition := func(row, col int) bool {
		return row >= 0 && row < numRows && col >= 0 && col < numCols
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			for _, c := range directions {
				dx, dy := c[0], c[1]
				x, y   := row, col
				valid  := true

				for k := 0; k < len(searchTerm); k++ {
					if !validPosition(x, y) || grid[x][y] != searchTerm[k] {
						valid = false
						break
					}
					x += dx
					y += dy
				}

				if valid {
					total++
				}
			}
		}
	}

    return total
}

// solvePartTwo takes an input 2d matrix and
// searches for two instances of 'MAS' together
// in an X shape.
// The 'MAS' can be forward or backwards
// It returns the number of matches found.
func solvePartTwo() int {
	grid       := loadInputGrid()
	numRows    := len(grid)
	numCols    := len(grid[0])
	total      := 0

	validPosition := func(row, col int) bool {
		return row > 0 && row < numRows - 1 && col > 0 && col < numCols - 1
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'A' && validPosition(row, col) {
				// Construct diagonal words.
				// Diagonal 1: top-left, center, bottom-right.
				diag1 := string([]byte{
					grid[row-1][col-1],
					grid[row][col],
					grid[row+1][col+1],
				})
				// Diagonal 2: top-right, center, bottom-left.
				diag2 := string([]byte{
					grid[row-1][col+1],
					grid[row][col],
					grid[row+1][col-1],
				})

				// Check if each diagonal forms either "MAS" or "SAM".
				validDiag := func(s string) bool {
					return s == "MAS" || s == "SAM"
				}

				if validDiag(diag1) && validDiag(diag2) {
					total++
				}
			}
		}
	}

    return total
}

// loadInputGrid gets the data in a usable format
// from the input file
func loadInputGrid() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return result
}

func loadTestInputGrid() []string {
	testInp := `
.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

	testInp = strings.TrimSpace(testInp)
	return strings.Split(testInp, "\n")
}
