package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	ages, err := readAndParseInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution for part 1: %d\n", computePartOne(ages))

	fmt.Printf("Solution for part 2: %d\n", computePartTwo(ages))

}

func computePartOne(ages []int) int {
	const iterations = 80

	agesCopy := make([]int, len(ages))
	copy(agesCopy, ages)
	return len(simulateAges(agesCopy, iterations))
}

func computePartTwo(ages []int) int {
	const iterations = 256

	agesCopy := make([]int, len(ages))
	copy(agesCopy, ages)
	return len(simulateAges(agesCopy, iterations))
}

func simulateAges(ages []int, iterations int) []int {
	var numIterations int
	var zeroCount int

	for numIterations < iterations {
		for idx, age := range ages {
			if age == 0 {
				ages[idx] = 6
				zeroCount++
			} else {
				ages[idx] = ages[idx] - 1
			}
		}

		for i := 0; i < zeroCount; i++ {
			ages = append(ages, 8)
		}

		numIterations++
		zeroCount = 0
	}
	return ages
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

	var ages []int

	for _, m := range strings.Split(string(content), ",") {
		age, _ := strconv.Atoi(m)
		ages = append(ages, age)
	}

	return ages, nil
}
