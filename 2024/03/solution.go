package aoc_2024_03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"io/ioutil"
)

const (
	inputFile      = "input.txt"
	mainRegexMatch = "mul\\(\\d{1,3},\\d{1,3}\\)"
	numberRegex    = "\\d{1,3}"

    // Get all the matches we want in a single regex using grouping
	// Each match is a slice of strings:
	//   - match[0]: the entire match.
	//   - match[1]: will be "do()", "don't()", or (for a mul instruction) equal to the full string.
	//   - match[2] and match[3]: available only for a mul instruction.
	partTwoRegex   = `(do\(\)|don't\(\)|mul\(([0-9]{1,3}),([0-9]{1,3})\))`
)

// Scan an input, find valid pairs of numbers in a
// `mul(x,y)` clause, then find the product of all of them
// A number is a valid pair if it is:
// 1) inside a call to `mul()`, with *no spaces*
// 2) 1 to 3 digits
func solve() int {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer file.Close()

	grandTotal := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		// Find the occurrences of the `mul` function in the current line of text
		re := regexp.MustCompile(mainRegexMatch)
		matchedMuls := re.FindAllString(currentLine, -1)

		lineTotal := 0

		for _, match := range matchedMuls {

			// Multiply each number in the mul function
			mulTotal := 1
			re := regexp.MustCompile(numberRegex)
			matchedMultipliers := re.FindAllString(match, -1)
			for _, numInMul := range matchedMultipliers {
				num, err := strconv.Atoi(numInMul)
				if err != nil {
					fmt.Println(err)
					return -1
				}

				if num > 0 {
					mulTotal *= num
				}
			}

			lineTotal += mulTotal
		}

		grandTotal += lineTotal
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return grandTotal
}

// Just the same thing except we add a switch to temporarily disable
// the addition step. The switch is denoted by a `control clause`; either
// `do()` or `don't()` -- the latter is the 'off' switch
func solvePartTwo() int {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
        fmt.Println(err)
        return -1
	}
	text := string(data)

	re, err := regexp.Compile(partTwoRegex)
	if err != nil {
        fmt.Println(err)
        return -1
	}

	matches := re.FindAllStringSubmatch(text, -1)
	if matches == nil || len(matches) == 0 {
        fmt.Println("No matches found")
        return -1
    }

	enabled := true  // Initially, mul instructions are enabled.
	grandTotal := 0

	for _, match := range matches {
		switch match[1] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			// We have a `mul` instruction
			if !enabled {
				continue
			}

			// match[2] and match[3] should have the two numbers.
			a, err := strconv.Atoi(match[2])
			if err != nil {
                fmt.Println(err)
                return -1
			}
			b, err := strconv.Atoi(match[3])
			if err != nil {
                fmt.Println(err)
                return -1
			}
			product := a * b
			grandTotal += product
		}
	}

    return grandTotal
}