package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Problem: https://adventofcode.com/2020/day/9
// Input: https://adventofcode.com/2020/day/9/input
func main() {
	numbers, err := readAndParseInput("input.txt")
	if err != nil {
		panic(fmt.Errorf("readAndParseInput returned %v\n", err))
	}
	//fmt.Println(numbers)
	fmt.Println(computePartOne(numbers, 25))
	fmt.Println(computePartTwo(numbers))
}

func computePartOne(numbers []int, preambleLen int) int {
	preambleStart := 0
	preambleStop := preambleLen

	for i := preambleStop; i < len(numbers); i++ {
		if !isValidNumber(numbers[i], preambleStart, preambleStop, numbers) {
			return numbers[i]
		}
		preambleStart++
		preambleStop++
	}
	return -1
}

func computePartTwo(numbers []int) int {
	invalidNumber := computePartOne(numbers, 25)
	for i := 0; i < len(numbers)-1; i++ {
		currSum := numbers[i]
		for j := i+1; j < len(numbers); j++ {
			if currSum == invalidNumber {
				contiguousSequence := numbers[i:j]
				sort.Ints(contiguousSequence)

				return contiguousSequence[0] + contiguousSequence[len(contiguousSequence)-1]
			}
			if currSum > invalidNumber {
				break
			}
			currSum += numbers[j]
		}
	}
	return -1
}

func isValidNumber(n int, preambleStart, preambleStop int, numbers []int) bool {
	for i := preambleStart; i < preambleStop-1; i++ {
		for j := preambleStart + 1; j < preambleStop; j++ {
			if numbers[i]+numbers[j] == n {
				return true
			}
		}
	}
	return false
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
