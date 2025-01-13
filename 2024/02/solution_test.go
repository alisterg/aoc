package aoc_2024_02

import (
	"testing"
)

func TestSolution(t *testing.T) {
	//input := ""
	expected := 246

	if actual := solve(); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}

func TestSolutionPartTwo(t *testing.T) {
	//input := ""
	expected := 318

	if actual := solvePartTwo(); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}
