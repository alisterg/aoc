package aoc_2024_02

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func solve() int {
	matrix := loadInput()

	numSafeReports := 0
	for _, line := range matrix {
		if isReportSafe(line) {
			numSafeReports++
		}
	}

	return numSafeReports
}

func solvePartTwo() int {
	matrix := loadInput()

	numSafeReports := 0
	for _, line := range matrix {
		if isReportSafeWithTolerance(line) {
			numSafeReports++
		}
	}

	return numSafeReports
}

// Checks if the report is safe, with a twist:
// We can tolerate a single bad level; for example if removing a single level would make
// a report safe, the report is then counted as safe.
func isReportSafeWithTolerance(report []int) bool {
	if isReportSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		newReport := make([]int, len(report)-1)
		copy(newReport, report[:i])
		copy(newReport[i:], report[i+1:])

		if isReportSafe(newReport) {
			return true
		}
	}

	return false
}

// isReportSafe tells us whether a report (list of numbers) is `safe`
// `safe` means:
//
//	the levels are either *all increasing* or *all decreasing*
//	any two adjacent levels differ by *at least one* and *at most three*
func isReportSafe(report []int) bool {
	currentTrajectory := "none"
	previousTrajectory := "none"

	for i := 1; i < len(report); i++ {
		difference := math.Abs(float64(report[i] - report[i-1]))
		if difference < 1 || difference > 3 {
			return false
		}

		if report[i-1] < report[i] {
			currentTrajectory = "increasing"
		} else {
			currentTrajectory = "decreasing"
		}

		// We don't have a previous trajectory for the first element
		if i == 1 {
			previousTrajectory = currentTrajectory
			continue
		}

		if currentTrajectory != previousTrajectory {
			return false
		}

		previousTrajectory = currentTrajectory
	}

	return true
}

func loadInput() [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		var line []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			line = append(line, num)
		}
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return result
}
