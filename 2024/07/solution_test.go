package aoc_2024_07

import (
	"os"
	"testing"
)

const testInput = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestSolution(t *testing.T) {
	expected := 3749

	if actual := solve(testInput); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}

func TestSolutionRealInput(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	expected := 1260333054159

	if actual := solve(input); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}

func TestSolutionPartTwo(t *testing.T) {
	expected := 11387

	if actual := solvePartTwo(testInput); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}

func TestSolutionPartTwoRealInput(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	expected := 162042343638683

	if actual := solvePartTwo(input); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}
