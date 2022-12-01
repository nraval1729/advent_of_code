package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lines, err := readAndParseInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution for part 1: %d\n", computePartOne(lines))
	fmt.Printf("Solution for part 2: %d\n", computePartTwo(lines))
}

func computePartOne(lines []string) int {
	var syntaxErrorScore int

	openingBraces := []string{"{", "(", "[", "<"}
	openingBraceToClosingBrace := map[string]string{
		"{": "}", "(": ")", "[": "]", "<": ">",
	}
	closingBraceToErrorScore := map[string]int {
		")": 3, "]": 57, "}": 1197, ">": 25137,
	}

	for _, line := range lines {
		var stack []string

		for _, s := range line {
			if contains(openingBraces, string(s)) {
				stack = append(stack, string(s))
			} else {
				popped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if openingBraceToClosingBrace[popped] != string(s) {
					syntaxErrorScore += closingBraceToErrorScore[string(s)]
				}
			}
		}
	}

	return syntaxErrorScore
}

func computePartTwo(lines []string) int {
	return 1
}

func contains(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}

	return false
}

func readAndParseInput(inputFilename string) ([]string, error) {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, err
	}

	var lines []string

	for _, line := range strings.Split(string(content), "\n") {
		lines = append(lines, line)
	}

	return lines, nil

}
