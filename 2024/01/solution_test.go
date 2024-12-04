package aoc_2024_01

import (
	"testing"
)

func TestSolution(t *testing.T) {
	//input := ""
	expected := 11

	if actual := solve(); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}

func TestSolutionPartTwo(t *testing.T) {
	//input := ""
	expected := 11

	if actual := solveSimilarityScore(); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}
