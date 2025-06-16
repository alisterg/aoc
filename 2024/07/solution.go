package aoc_2024_07

import (
	"fmt"
	"strconv"
	"strings"
)

// evaluateSeq evaluates the list of numbers with the given operators.
// It returns the result when operations are evaluated strictly left-to-right.
// ops is a slice of booleans with length = len(nums)-1, where false means '+', and true means '*'.
func evaluateSeq(nums []int, ops []bool) int {
	res := nums[0]
	for i, op := range ops {
		if op {
			res *= nums[i+1]
		} else {
			res += nums[i+1]
		}
	}

	return res
}

// generateOps generates all combinations of operators (as boolean slices) for n-1 positions.
func generateOps(n int) [][]bool {
	combinations := [][]bool{}
	total := 1 << (n - 1)
	for i := range total {
		ops := make([]bool, n-1)
		for j := range n - 1 {
			// bit j: false for addition, true for multiplication
			ops[j] = (i>>j)&1 == 1
		}
		combinations = append(combinations, ops)
	}
	return combinations
}

func solve(inp string) int {
	input := strings.Split(inp, "\n")
	totalCalibration := 0

	for _, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// Expected format: "testValue: num num num ..."
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			fmt.Printf("Invalid input line: %s\n", line)
			continue
		}
		targetStr := strings.TrimSpace(parts[0])
		numsStr := strings.TrimSpace(parts[1])

		// Parse target value.
		target, err := strconv.Atoi(targetStr)
		if err != nil {
			fmt.Printf("Error parsing target in line: %s\n", line)
			continue
		}

		// Parse the numbers.
		numStrs := strings.Fields(numsStr)
		nums := []int{}
		for _, s := range numStrs {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Error parsing number %s in line: %s\n", s, line)
				continue
			}
			nums = append(nums, num)
		}

		// If only one number exists, then check if it equals target.
		if len(nums) == 1 {
			if nums[0] == target {
				totalCalibration += target
			}
			continue
		}

		// Try every combination of operators.
		opsList := generateOps(len(nums))
		for _, ops := range opsList {
			if evaluateSeq(nums, ops) == target {
				totalCalibration += target
				break
			}
		}

	}

	return totalCalibration
}
