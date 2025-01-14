package aoc_2024_03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	inputFile      = "input.txt"
	mainRegexMatch = "mul\\(\\d{1,3},\\d{1,3}\\)"
	numberRegex    = "\\d{1,3}"
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
