package aoc_2024_08

import (
	"os"
	"testing"
)

const testInput = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

func TestSolution(t *testing.T) {
	expected := 14

	if actual := solve(testInput); actual != expected {
		t.Errorf("===FAIL===: Expected `%v`, got `%v`", expected, actual)
	}
}

func TestSolutionRealInput(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	expected := 259

	if actual := solve(input); actual != expected {
		t.Errorf("===FAIL===: Expected `%v`, got `%v`", expected, actual)
	}
}
