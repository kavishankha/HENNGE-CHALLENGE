package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//step 1 scan fist line get the N value(this give you hoe much line to scan after)
//step 2 scan until N*2 lines (use recursive insted loop to scan)
//step 3 after N*2 scan stops (blank inputs not counted)
//step 4 store it in lice/arry (X 1 raw Y second raw)
//step 5 use recursive funtions insted of for loop for calculation
//each 2 set frist raw is x second raw is set of Yn
//check count X match with each count Yn
//if maching that get power of each Yn in raw get a sum and push sum to lice
// else push to arry -1
//step 6 print the arry line by line
//only use int32 to outpuut
// handle erorrs

func parseXLine(line string) (int, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return 0, fmt.Errorf("X line is empty")
	}

	parts := strings.Fields(line)
	if len(parts) != 1 {
		return 0, fmt.Errorf("X line should have only one value, found: %d values", len(parts))
	}

	xValue, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("X line contains non-integer value: %s", parts[0])
	}

	if xValue <= 0 || xValue > 100 {
		return 0, fmt.Errorf("X value must be in range 0 < X <= 100, got: %d", xValue)
	}

	return xValue, nil
}

func parseYValues(parts []string, idx int, values []int) error {
	if idx >= len(parts) {
		return nil
	}

	val, err := strconv.Atoi(parts[idx])
	if err != nil {
		return fmt.Errorf("Y line contains non-integer value: %s", parts[idx])
	}
	if val < -100 || val > 100 {
		return fmt.Errorf("Yn value must be in range -100 <= Yn <= 100, got: %d", val)
	}
	values[idx] = val

	return parseYValues(parts, idx+1, values)
}

func parseYLine(line string) ([]int, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, fmt.Errorf("Y line is empty")
	}

	parts := strings.Fields(line)
	values := make([]int, len(parts))

	err := parseYValues(parts, 0, values)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func readLines(scanner *bufio.Scanner, lineCount int, maxLines int, data map[int][]int) error {
	if lineCount >= maxLines {
		return nil
	}

	if !scanner.Scan() {
		if scanner.Err() != nil {
			return fmt.Errorf("scanner error: %v", scanner.Err())
		}
		return fmt.Errorf("unexpected end of input")
	}

	line := scanner.Text()

	if lineCount%2 == 0 {
		xValue, err := parseXLine(line)
		if err != nil {
			return fmt.Errorf("line %d: %v", lineCount+1, err)
		}
		data[lineCount] = []int{xValue}
	} else {
		yValues, err := parseYLine(line)
		if err != nil {
			return fmt.Errorf("line %d: %v", lineCount+1, err)
		}
		data[lineCount] = yValues
	}

	return readLines(scanner, lineCount+1, maxLines, data)
}

func calculatePowerSum(yValues []int, idx int, sum int) (int, error) {
	if idx >= len(yValues) {
		return sum, nil
	}

	var value int
	if yValues[idx] < 0 {
		// Only negative values get power of 4 and added to sum
		value = int(math.Pow(float64(yValues[idx]), 4))

		if sum > 0 && value > 0 && sum > (1<<31-1)-value {
			fmt.Println("Warning: Potential overflow detected")
		}

		return calculatePowerSum(yValues, idx+1, sum+value)
	} else {
		// Positive values (and zero) are completely ignored
		return calculatePowerSum(yValues, idx+1, sum)
	}
}

func processSets(setIdx int, totalSets int, data map[int][]int, results map[int]int) error {
	if setIdx >= totalSets {
		return nil
	}

	xValue := data[setIdx*2][0]
	yValues := data[setIdx*2+1]

	if xValue != len(yValues) {
		results[setIdx] = -1
	} else {
		sum, err := calculatePowerSum(yValues, 0, 0)
		if err != nil {
			return fmt.Errorf("error calculating power sum for set %d: %v", setIdx, err)
		}
		results[setIdx] = sum
	}

	return processSets(setIdx+1, totalSets, data, results)
}

func printResults(idx int, total int, results map[int]int) {
	if idx >= total {
		return
	}
	fmt.Println(results[idx])
	printResults(idx+1, total, results)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	data := make(map[int][]int)
	results := make(map[int]int)

	fmt.Println("enter the input")

	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error reading N:", err)
		return
	}

	if n < 1 || n > 100 {
		fmt.Println("Error: N must be in range 1 <= N <= 100")
		return
	}

	fmt.Scanln()

	err = readLines(scanner, 0, n*2, data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = processSets(0, n, data, results)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printResults(0, n, results)
}
