package aoc_2024_04

import (
	"testing"
)

func TestSolution(t *testing.T) {
	expected := 2654

	if actual := solve(); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}

func TestSolutionPartTwo(t *testing.T) {
	expected := 1990

	if actual := solvePartTwo(); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}
