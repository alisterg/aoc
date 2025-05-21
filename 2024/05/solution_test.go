package aoc_2024_05

import (
	"testing"
)

func TestSolution(t *testing.T) {
	expected := 6260

	if actual := solve(); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}

func TestGetMiddlePageValue(t *testing.T) {
	input := "75,47,61,53,29"
	expected := 61

	if actual := getMiddlePageValue(input); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}

func TestSolutionPartTwo(t *testing.T) {
	expected := 5346

	if actual := solvePartTwo(); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}
