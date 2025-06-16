package aoc_2024_07

import (
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	input := `
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
	expected := 3749

	if actual := solve(input); actual != expected {
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
