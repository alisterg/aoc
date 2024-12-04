package aoc_2024_01

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve() int {
	listOne, listTwo := loadInput()

	sort.Ints(listOne)
	sort.Ints(listTwo)

	result := 0

	for i := 0; i < len(listOne); i++ {
		absoluteValue := math.Abs(float64(listOne[i] - listTwo[i]))
		result += int(absoluteValue)
	}

	return result
}

func solveSimilarityScore() int {
	listOne, listTwo := loadInput()

	score := 0
	for i := 0; i < len(listOne); i++ {
		occurrences := 0
		for j := 0; j < len(listTwo); j++ {
			if listOne[i] == listTwo[j] {
				occurrences++
			}
		}

		score += listOne[i] * occurrences
	}

	return score
}

func loadInput() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()

	var col1, col2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			fmt.Println("Invalid line format:", scanner.Text())
			continue
		}

		num1, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Invalid number:", fields[0])
			continue
		}

		num2, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Invalid number:", fields[1])
			continue
		}

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return col1, col2
}
