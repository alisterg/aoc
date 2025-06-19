package aoc_2024_07

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PLUS     = 0
	MULTIPLY = 1
	CONCAT   = 2
)

// evaluateSeq evaluates the list of numbers with the given operators.
// It returns the result when operations are evaluated strictly left-to-right.
// ops represents one possible combination of operators for the sequence
func evaluateSeq(nums []int, ops []int) int {
	res := nums[0]
	for i, op := range ops {
		if i+1 >= len(nums) {
			break
		}

		switch op {
		case MULTIPLY:
			res *= nums[i+1]
		case PLUS:
			res += nums[i+1]
		case CONCAT:
			first := strconv.Itoa(res)
			second := strconv.Itoa(nums[i+1])
			res, _ = strconv.Atoi(first + second)
		}
	}

	return res
}

// generateOps generates all combinations of operators for the sequence length
func generateOps(operators []int, length int) [][]int {
	if length == 1 {
		var combinations [][]int
		for _, operator := range operators {
			combinations = append(combinations, []int{operator})
		}
		return combinations
	}

	var combinations [][]int
	for _, operator := range operators {
		subCombinations := generateOps(operators, length-1)
		for _, subCombination := range subCombinations {
			combination := append([]int{operator}, subCombination...)
			combinations = append(combinations, combination)
		}
	}

	return combinations
}

// parseNums parses one line of the input into:
//   - target number
//   - slice of operands
func parseNums(line string) (int, []int) {
	line = strings.TrimSpace(line)
	if line == "" {
		return -1, nil
	}
	// Expected format: "testValue: num num num ..."
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		fmt.Printf("Invalid input line: %s\n", line)
		return -1, nil
	}
	targetStr := strings.TrimSpace(parts[0])
	numsStr := strings.TrimSpace(parts[1])

	// Parse target value.
	target, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Printf("Error parsing target in line: %s\n", line)
		return -1, nil
	}

	// Parse the numbers.
	numStrs := strings.Fields(numsStr)
	nums := []int{}
	for _, s := range numStrs {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Error parsing number %s in line: %s\n", s, line)
			return -1, nil
		}
		nums = append(nums, num)
	}

	return target, nums
}

func solve(inp string) int {
	input := strings.Split(inp, "\n")
	totalCalibration := 0

	for _, line := range input {
		target, operands := parseNums(line)
		if target == -1 || operands == nil {
			continue
		}

		// If only one number exists, then check if it equals target.
		if len(operands) == 1 {
			if operands[0] == target {
				totalCalibration += target
			}
			continue
		}

		// Try every combination of operators.
		opsList := generateOps([]int{PLUS, MULTIPLY}, len(operands))
		for _, ops := range opsList {
			if evaluateSeq(operands, ops) == target {
				totalCalibration += target
				break
			}
		}
	}

	return totalCalibration
}

func solvePartTwo(inp string) int {
	input := strings.Split(inp, "\n")
	result := 0

	for _, line := range input {
		target, operands := parseNums(line)
		if target == -1 || operands == nil {
			continue
		}

		// If only one number exists, then check if it equals target.
		if len(operands) == 1 {
			if operands[0] == target {
				result += target
			}
			continue
		}

		// Try every combination of operators.
		opsList := generateOps([]int{PLUS, MULTIPLY, CONCAT}, len(operands)-1)
		for _, ops := range opsList {
			if evaluateSeq(operands, ops) == target {
				result += target
				break
			}
		}

	}

	return result
}
