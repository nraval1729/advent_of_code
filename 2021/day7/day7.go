package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	positions, err := readAndParseInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution for part 1: %d\n", computePartOne(positions))

	fmt.Printf("Solution for part 2: %d\n", computePartTwo(positions))

}

func computePartOne(positions []int) int {
	var currSum int
	minSum := math.Inf(0)

	for i := 0; i < len(positions); i++ {
		currSum = 0
		for j := 0; j < len(positions); j++ {
			if positions[j] > positions[i] {
				currSum += positions[j] - positions[i]
			} else {
				currSum += positions[i] - positions[j]
			}
		}

		if float64(currSum) < minSum {
			minSum = float64(currSum)
		}
	}

	return int(minSum)
}

func computePartTwo(positions []int) int {
	var currSum int

	minSum := math.Inf(0)

	for i := 0; i <= max(positions); i++ {
		currSum = 0
		for _, position := range positions {
			if position > i {
				currSum += sumN(position - i)
			} else {
				currSum += sumN(i - position)
			}
		}

		if float64(currSum) < minSum {
			minSum = float64(currSum)
		}
	}

	return int(minSum)
}

func sumN(n int) int {
	var sum int

	for i:=1; i<=n; i++ {
		sum += i
	}

	return sum
}

func max(ints []int) int{
	var max int

	for _, i := range ints {
		if i > max {
			max = i
		}
	}

	return max
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

	var positions []int

	for _, m := range strings.Split(string(content), ",") {
		age, _ := strconv.Atoi(m)
		positions = append(positions, age)
	}

	return positions, nil
}
