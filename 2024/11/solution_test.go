package aoc_2024_11

import (
	"os"
	"testing"
)

const testInput = `
125 17
`

func TestSolution(t *testing.T) {
	expected := 55312

	if actual := solve(testInput); actual != expected {
		t.Errorf("===FAIL===: Expected `%v`, got `%v`", expected, actual)
	}
}

func TestSolutionRealInput(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	expected := 183248

	if actual := solve(input); actual != expected {
		t.Errorf("===FAIL===: Expected `%v`, got `%v`", expected, actual)
	}
}
