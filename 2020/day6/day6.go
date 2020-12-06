package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Problem: https://adventofcode.com/2020/day/6
// Input: https://adventofcode.com/2020/day/6/input
func main() {
	groupAnswers, err := readGroupAnswers("input.txt")
	if err != nil {
		panic(fmt.Errorf("readGroupAnswers threw %v\n", err))
	}
	fmt.Println(computePartOne(groupAnswers))
	fmt.Println(computePartTwo(groupAnswers))
}

func computePartOne(groupAnswers []string) int {
	numAnswers := 0
	for _, groupAnswer := range groupAnswers {
		answersFromEveryone := strings.Split(groupAnswer, "\n")
		numAnswers += computeYesAnswersFromAnyone(answersFromEveryone)
	}
	return numAnswers
}

func computePartTwo(groupAnswers []string) int {
	numAnswers := 0
	for _, groupAnswer := range groupAnswers {
		answersFromEveryone := strings.Split(groupAnswer, "\n")
		numAnswers += computeYesAnswersFromEveryone(answersFromEveryone)
	}
	return numAnswers
}

func computeYesAnswersFromEveryone(answersFromEveryone []string) int {
	numYesAnswersByEveryone := 0
	answerToCount := make(map[string]int)

	for _, answersFromPerson :=range answersFromEveryone {
		answers := strings.Split(answersFromPerson, "")
		for _, answer := range answers {
			if _, ok := answerToCount[answer];  ok {
				answerToCount[answer]++
			} else {
				answerToCount[answer] = 1
			}
		}
	}
	for _, count := range answerToCount {
		if count == len(answersFromEveryone) {
			numYesAnswersByEveryone++
		}
	}
	return numYesAnswersByEveryone
}

func computeYesAnswersFromAnyone(answersFromEveryone []string) int {
	answeredYesMap := make(map[string]bool)

	for _, answersFromPerson :=range answersFromEveryone {
		answers := strings.Split(answersFromPerson, "")
		for _, answer := range answers {
			if _, ok := answeredYesMap[answer];  !ok {
				answeredYesMap[answer] = true
			}
		}
	}
	return len(answeredYesMap)
}

func readGroupAnswers(inputFileName string) ([]string, error) {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n\n"), nil
}
