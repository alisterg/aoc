package aoc_2024_06

import (
	"testing"
)

func TestSolution(t *testing.T) {
	expected := 5305

	if actual := solve(); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}
