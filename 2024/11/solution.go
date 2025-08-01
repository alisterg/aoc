package aoc_2024_11

import (
	"strconv"
	"strings"
)

const (
	NUM_BLINKS       = 25
	STONE_ZERO       = "0"
	STONE_ONE        = "1"
	STONE_MULTIPLIER = 2024
)

// `blink` 25 times and return the number of stones at 25 blinks
func solve(input string) int {

	result := strings.Fields(input)

	for range NUM_BLINKS {
		result = blink(result)
	}

	return len(result)
}

/*
Every time you blink, the stones each simultaneously change according to the first applicable rule in this list:

1. If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.

2. If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)

3. If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.

No matter how the stones change, their order is preserved
*/
func blink(stones []string) []string {

	var result []string

	for _, v := range stones {
		if v == STONE_ZERO {
			result = append(result, STONE_ONE)
		} else if isEven(v) {
			first, second := splitStone(v)
			// first half of v
			result = append(result, first)
			// second half of v
			result = append(result, second)
		} else {
			// v * 2024
			intV, _ := strconv.Atoi(v)
			resV := strconv.Itoa(intV * STONE_MULTIPLIER)
			result = append(result, resV)
		}
	}

	return result
}

// remove leading zeroes
func normalise(inp string) string {
	intV, _ := strconv.Atoi(inp)
	return strconv.Itoa(intV)
}

// split a string in half
func splitStone(inp string) (string, string) {
	half := len(inp) / 2
	return normalise(inp[:half]), normalise(inp[half:])
}

func isEven(inp string) bool {
	return len(inp)%2 == 0
}
