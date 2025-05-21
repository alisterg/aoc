package aoc_2024_05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

// loadInput loads the file as a []string
func loadInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return result
}

// split the file into: 1) ordering rules & 2) page updates
func splitInput(input []string) ([]string, []string) {
	var orderingRules []string
	var pageUpdates []string

	isOrderingRules := true

	for _, line := range input {
		if line == "" {
			isOrderingRules = false
			continue
		}

		if isOrderingRules {
			orderingRules = append(orderingRules, line)
		} else {
			pageUpdates = append(pageUpdates, line)
		}
	}

	return orderingRules, pageUpdates
}

// sort according to the ordering rules
// if the sorted array is the same as the provided `update` then it is valid
func sortPageUpdate(orderingRules []string, update string) string {
	parts := strings.Split(update, ",")

	sort.SliceStable(parts, func(i, j int) bool {
		for _, rule := range orderingRules {
			ruleParts := strings.Split(rule, "|")

			if (parts[i] == ruleParts[0] && parts[j] == ruleParts[1]) {
				return true
			}

			if (parts[j] == ruleParts[0] && parts[i] == ruleParts[1]) {
				return false
			}

		}

		return parts[i] < parts[j]
	})

	return strings.Join(parts, ",")
}

// find the centre page value for summing
// nothing was mentioned about all pageUpdates being of odd size, what
// if it's even? I will ignore for now
func getMiddlePageValue(update string) int {
	arr := strings.Split(update, ",")
	middlePos := int(len(arr) / 2)
	result, _ := strconv.Atoi(arr[middlePos])
	return result
}

// check each page update and if valid, sum the middle page number
func solve() int {
	input := loadInput()
	orderingRules, pageUpdates := splitInput(input)

	total := 0

	for _, update := range pageUpdates {
		sortedUpdate := sortPageUpdate(orderingRules, update)
		if sortedUpdate == update {
			total += getMiddlePageValue(update)
		}
	}

	return total
}

// check each page update and if valid, sum the middle page number
func solvePartTwo() int {
	input := loadInput()
	orderingRules, pageUpdates := splitInput(input)

	total := 0

	for _, update := range pageUpdates {
		sortedUpdate := sortPageUpdate(orderingRules, update)
		if sortedUpdate != update {
			total += getMiddlePageValue(sortedUpdate)
		}
	}

	return total
}
