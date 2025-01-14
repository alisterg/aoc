package aoc_2024_03

import (
	"testing"
)

func TestSolution(t *testing.T) {
	//input := ""
	expected := 161085926

	if actual := solve(); actual != expected {
		t.Errorf("Expected `%v`, got `%v`", expected, actual)
	}
}
