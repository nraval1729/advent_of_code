package main

import (
	"fmt"
)

// Problem: https://adventofcode.com/2020/day/15
func main() {
	startingNumbers := []int{6, 19, 0, 5, 7, 13, 1}
	fmt.Println(computePartOne(startingNumbers, 2020))
	fmt.Println(computePartTwo(startingNumbers))
}

// This is a pretty elegant solution in my opinion.
// I had initially solved part 1 by storing all the indices where a number was previously spoken
// However, that is catastrophic for part 2 given its input size. The main bottleneck was the copy() function.
// So, I decided to only always store the LAST TWO indices where a number was previously spoken. VOILA! No copy() and runs FAST!
func computePartOne(nums []int, lenSpokenNumbers int) int {
	spokenNumbers := make([]int, lenSpokenNumbers)
	numberToSpokenBeforeIndices := make(map[int][]int)

	// Add all starting numbers to spokenNumbers first
	for i := 0; i < len(nums); i++ {
		spokenNumbers[i] = nums[i]
		numberToSpokenBeforeIndices[nums[i]] = []int{i}
	}

	// Main loop
	for i := len(nums); i < len(spokenNumbers); i++ {
		lastNumberSpoken := spokenNumbers[i-1]
		spokenBeforeIndices, ok := numberToSpokenBeforeIndices[lastNumberSpoken]
		var currSpokenNumber int

		if ok && len(spokenBeforeIndices) > 1 {
			// if spoken before
			currSpokenNumber = spokenBeforeIndices[1] - spokenBeforeIndices[0]
			spokenNumbers[i] = currSpokenNumber
		} else {
			//	not spoken before
			currSpokenNumber = 0
			spokenNumbers[i] = currSpokenNumber
		}
		// Update mapping for spokenNumber
		currSpokenNumberSpokenBeforeIndices, ok := numberToSpokenBeforeIndices[currSpokenNumber]
		if ok {
			// only append last two occurences
			numberToSpokenBeforeIndices[currSpokenNumber] = append(
				[]int{currSpokenNumberSpokenBeforeIndices[len(currSpokenNumberSpokenBeforeIndices)-1], i},
			)
		} else {
			numberToSpokenBeforeIndices[currSpokenNumber] = []int{i}
		}
	}
	return spokenNumbers[lenSpokenNumbers-1]
}

func computePartTwo(nums []int) int {
	return computePartOne(nums, 30000000)
}
