package aoc_2024_10

import (
	"os"
	"testing"
)

const testInput = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func TestSolution(t *testing.T) {
	expected := 36
	if actual := solve(testInput); actual != expected {
		t.Errorf("===FAIL===: Expected `%v`, got `%v`", expected, actual)
	}
}

func TestSolutionRealInput(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	expected := -1

	if actual := solve(input); actual != expected {
		t.Errorf("===FAIL===: Expected `%v`, got `%v`", expected, actual)
	}
}
