package aoc_2023_01

import (
	"testing"
)

func TestSolution(t *testing.T) {
	input := ""
	expected := ""

	if actual := solve(input); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}
