package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type BagColourCountCombo struct {
	Colour string
	Count  int
}

// Problem: https://adventofcode.com/2020/day/7
// Input: https://adventofcode.com/2020/day/7/input
func main() {
	bagToContainedBags, bagToContainedBagsCount, err := createBagsGraph("input.txt")
	if err != nil {
		panic(fmt.Errorf("createBagsGraph threw %v\n", err))
	}
	fmt.Println(computePartOne(bagToContainedBags))
	fmt.Println(computePartTwo(bagToContainedBags, bagToContainedBagsCount))
}

func computePartOne(bagToContainedBags map[string][]string) int {
	myBagColour := "shiny gold"
	numBagsThatCanContainMyBag := 0

	for bagColour, _ := range bagToContainedBags {
		if dfs(myBagColour, bagColour, bagToContainedBags) {
			numBagsThatCanContainMyBag++
		}
	}
	return numBagsThatCanContainMyBag
}

func computePartTwo(bagToContainedBags map[string][]string, bagToContainedBagsCount map[string][]int) int {
	return countContainedBags(0, "shiny gold", bagToContainedBags, bagToContainedBagsCount)
}

func countContainedBags(currCount int, targetBagColour string, bagToContainedBags map[string][]string, bagToContainedBagsCount map[string][]int) int {
	if bagToContainedBags[targetBagColour] == nil {
		return currCount
	}
	for idx, containedBag := range bagToContainedBags[targetBagColour] {
		currCount += bagToContainedBagsCount[targetBagColour][idx] * countContainedBags(0, containedBag, bagToContainedBags, bagToContainedBagsCount)
	}
	for _, c := range bagToContainedBagsCount[targetBagColour] {
		currCount += c
	}
	return currCount
}

func dfs(targetBagColour string, bagColour string, bagToContainedBags map[string][]string) bool {
	ret := false

	for _, containedBagColour := range bagToContainedBags[bagColour] {
		if containedBagColour == targetBagColour {
			return true
		}
		ret = ret || dfs(targetBagColour, containedBagColour, bagToContainedBags)
	}
	return ret
}

func createBagsGraph(inputFileName string) (map[string][]string, map[string][]int, error) {
	bagToContainedBags := make(map[string][]string)
	bagToContainedBagsCount := make(map[string][]int)

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, nil, err
	}

	for _, bagRule := range strings.Split(string(content), "\n") {
		colour, containedBags, containedBagsCount := parseContainedBagsAndCount(bagRule)
		bagToContainedBags[colour] = containedBags
		bagToContainedBagsCount[colour] = containedBagsCount
	}
	return bagToContainedBags, bagToContainedBagsCount, nil
}

func parseContainedBagsAndCount(bagRule string) (string, []string, []int) {
	var colour string
	var containedBags []string
	var containedBagsCount []int

	splitOnSpace := strings.Split(bagRule, " ")
	colour = strings.Join(splitOnSpace[:2], " ")

	// This bag doesn't contain any other bags
	if len(splitOnSpace) == 7 {
		return colour, nil, nil
	}

	for i := 4; i < len(splitOnSpace); i += 4 {
		count, _ := strconv.Atoi(splitOnSpace[i])
		containedBagColour := strings.Join(splitOnSpace[i+1:i+3], " ")
		containedBags = append(containedBags, containedBagColour)
		containedBagsCount = append(containedBagsCount, count)
	}
	return colour, containedBags, containedBagsCount
}
