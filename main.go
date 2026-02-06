package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Recursive function to sum 4th powers of non-positive numbers
func sumPower4(nums []int, index int) int {
	if index >= len(nums) {
		return 0
	}
	if nums[index] <= 0 {
		return nums[index]*nums[index]*nums[index]*nums[index] + sumPower4(nums, index+1)
	}
	return sumPower4(nums, index+1)
}

// Recursive function to process test cases
func processTestCases(lines []string, index int, results []int) []int {
	if index >= len(lines) {
		return results
	}

	// Parse X
	X, err := strconv.Atoi(lines[index])
	if err != nil {
		results = append(results, -1)
		return processTestCases(lines, index+2, results)
	}

	// Check if next line exists
	if index+1 >= len(lines) {
		results = append(results, -1)
		return results
	}

	// Parse numbers
	numStrs := strings.Fields(lines[index+1])
	if len(numStrs) != X {
		results = append(results, -1)
	} else {
		nums := make([]int, X)
		for i := 0; i < X; i++ {
			nums[i], _ = strconv.Atoi(numStrs[i])
		}
		results = append(results, sumPower4(nums, 0))
	}

	// Process next test case
	return processTestCases(lines, index+2, results)
}

// Recursive function to print results
func printResults(results []int, index int) {
	if index >= len(results) {
		return
	}
	fmt.Println(results[index])
	printResults(results, index+1)
}

func main() {
	// Read everything from stdin at once
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	// Split by newlines
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) == 0 {
		return
	}

	// Parse number of test cases N
	N, err := strconv.Atoi(lines[0])
	if err != nil || N < 1 || N > 100 {
		return
	}

	// Process test cases recursively
	results := processTestCases(lines[1:], 0, []int{})

	// Print results recursively
	printResults(results, 0)
}
