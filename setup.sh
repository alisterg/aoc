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

func solve(input string) string {
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

func TestSolution(t *testing.T) {
	input := ""
	expected := "abc"

	if actual := solve(input); actual != expected {
		t.Errorf("Expected \`%v\`, got \`%v\`", expected, actual)
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

# Create the doc files. TODO: automatically download them
if [ ! -f "$DIR/problem.md" ]; then
    echo "# $YEAR day $DAY" > "$DIR/problem.md"
	# curl https://adventofcode.com/$YEAR/day/$DAY > "$DIR/problem.html"
fi
if [ ! -f "$DIR/input.txt" ]; then
    touch "$DIR/input.txt"
	# curl https://adventofcode.com/$YEAR/day/$DAY/input > "$DIR/input.txt"
fi
