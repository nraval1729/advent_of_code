package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Problem: https://adventofcode.com/2020/day/10
// Input: https://adventofcode.com/2020/day/10/input
func main() {
	joltages, err := readAndParseInput("input.txt")
	if err != nil {
		panic(fmt.Errorf("readAndParseInput returned %v\n", err))
	}
	fmt.Println(computePartOne(joltages))
	fmt.Println(computePartTwo(joltages))
}

func computePartOne(joltages []int) int {
	num3JoltDifferences, num1JoltDifferences := 0, 0

	sort.Ints(joltages)
	if joltages[0] == 1 {
		num1JoltDifferences++
	}
	if joltages[0] == 3 {
		num3JoltDifferences++
	}

	for i := 1; i < len(joltages); i++ {
		difference := joltages[i] - joltages[i-1]

		if difference == 1 {
			num1JoltDifferences++
		}
		if difference == 3 {
			num3JoltDifferences++
		}
	}
	// Built in adapter is always +3 than the highest adaper joltage
	num3JoltDifferences++

	return num1JoltDifferences * num3JoltDifferences
}

func computePartTwo(joltages []int) int {
	sort.Ints(joltages)
	numWaysToGetTo := make([]int, len(joltages))
	numWaysToGetTo[0] = 1

	if joltages[1] <= 3 {
		numWaysToGetTo[1] = numWaysToGetTo[0] + 1
	}
	if joltages[2] <= 3 {
		numWaysToGetTo[2] = numWaysToGetTo[1] + 1
	}
	if joltages[2]-joltages[1] <= 3 {
		numWaysToGetTo[2]++
	}

	for i := 3; i < len(joltages); i++ {
		numWaysToGetTo[i] = 0
		if i-1 >= 0 && joltages[i]-joltages[i-1] <= 3 {
			numWaysToGetTo[i] += numWaysToGetTo[i-1]
		}
		if i-2 >= 0 && joltages[i]-joltages[i-2] <= 3 {
			numWaysToGetTo[i] += numWaysToGetTo[i-2]
		}
		if i-3 >= 0 && joltages[i]-joltages[i-3] <= 3 {
			numWaysToGetTo[i] += numWaysToGetTo[i-3]
		}
	}
	return numWaysToGetTo[len(joltages)-1]
}

func readAndParseInput(inputFilename string) ([]int, error) {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, err
	}

	input := strings.Split(string(content), "\n")
	inputParsedToInts := make([]int, len(input))

	for i, c := range input {
		inputParsedToInts[i], _ = strconv.Atoi(c)
	}
	return inputParsedToInts, nil
}
