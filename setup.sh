#!/bin/zsh

YEAR="${1:-}"
DAY="${2:-}"

if [ -z "$YEAR" ] || [ -z "$DAY" ]; then
	echo "Usage: $0 <YEAR> <DAY>"
	exit 1
fi

# pad day to 2 digits
DAYPADDED=$(printf "%02d" $DAY)
DIR="./$YEAR/$DAYPADDED"
PACKAGE="aoc_${YEAR}_${DAYPADDED}"

function template() {
	cat <<EOF
package $PACKAGE

func solve(input string) int {
    return input
}
EOF
}

function test_template() {
    cat <<EOF
package $PACKAGE

import (
	"testing"
)

const testInput = \`

\`

func TestSolution(t *testing.T) {
	expected := "abc"

	if actual := solve(testInput); actual != expected {
		t.Errorf("===FAIL===: Expected \`%v\`, got \`%v\`", expected, actual)
	}
}

func TestSolutionRealInput(t *testing.T) {
	data, _ := os.ReadFile("input.txt")
	input := string(data)
	expected := -1

	if actual := solve(input); actual != expected {
		t.Errorf("===FAIL===: Expected `%v`, got `%v`", expected, actual)
	}
}
EOF
}

# Create directories if required
if [ ! -d "$DIR" ]; then
	mkdir -p "$DIR"
fi

# Create source files
if [ ! -f "$DIR/solution.go" ]; then
	template > "$DIR/solution.go"
fi
if [ ! -f "$DIR/solution_test.go" ]; then
	test_template > "$DIR/solution_test.go"
fi
