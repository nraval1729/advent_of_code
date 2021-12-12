package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type signalEntry struct {
	Patterns          []string
	OutputValueDigits []string
}

func main() {
	signalEntries, err := readAndParseInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution for part 1: %d\n", computePartOne(signalEntries))

	fmt.Printf("Solution for part 2: %d\n", computePartTwo(signalEntries))

}

func computePartOne(signalEntries []signalEntry) int {
	var count int

	for _, entry := range signalEntries {
		for _, digit := range entry.OutputValueDigits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				count++
			}
		}
	}

	return count
}

func computePartTwo(signalEntries []signalEntry) int {
	return 1
}

func readAndParseInput(inputFilename string) ([]signalEntry, error) {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, err
	}

	var signalEntries []signalEntry

	for _, entryString := range strings.Split(string(content), "\n") {
		patterns, outputValueDigits := strings.Split(entryString, " | ")[0], strings.Split(entryString, " | ")[1]

		entry := signalEntry{
			Patterns:          strings.Split(patterns, " "),
			OutputValueDigits: strings.Split(outputValueDigits, " "),
		}

		signalEntries = append(signalEntries, entry)
	}

	return signalEntries, nil

}
