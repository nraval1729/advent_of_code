package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Problem: https://adventofcode.com/2020/day/8
// Input: https://adventofcode.com/2020/day/8/input
func main() {
	instructions, err := readAndParseInput("input.txt")
	if err != nil {
		panic(fmt.Errorf("readAndParseInput returned %v\n", err))
	}
	_, partOne := computePartOne(instructions)
	fmt.Println(partOne)
	fmt.Println(computePartTwo(instructions))
}

func computePartOne(instructions []string) (bool, int) {
	// program counter
	pc := 0
	acc := 0
	executedInstructions := make(map[int]bool)

	for {
		if pc >= len(instructions) {
			return true, acc
		}
		if _, ok := executedInstructions[pc]; ok {
			return false, acc
		}
		executedInstructions[pc] = true

		operation, operand := parseInstruction(instructions[pc])
		if operation == "acc" {
			acc += operand
			pc++
		} else if operation == "jmp" {
			pc += operand
		} else {
			pc++
		}
	}
}

func computePartTwo(instructions []string) int {
	var candidatePrograms [][]string

	// Create a slice of all possible replacement candidate programs
	for i := 0; i < len(instructions); i++ {
		operation, _ := parseInstruction(instructions[i])
		if operation == "nop" {
			instructionsCopy := make([]string, len(instructions))
			copy(instructionsCopy, instructions)
			instructionsCopy[i] = strings.Replace(instructionsCopy[i], "nop", "jmp", 1)

			candidatePrograms = append(candidatePrograms, instructionsCopy)
		}
		if operation == "jmp" {
			instructionsCopy := make([]string, len(instructions))
			copy(instructionsCopy, instructions)
			instructionsCopy[i] = strings.Replace(instructionsCopy[i], "jmp", "nop", 1)

			candidatePrograms = append(candidatePrograms, instructionsCopy)
		}
	}

	// For each candidate program check to see if it terminates. If it does then return the acc
	for _, candidateProgram := range candidatePrograms {
		terminated, acc := computePartOne(candidateProgram)
		if terminated {
			return acc
		}
	}
	return -1
}

func parseInstruction(instruction string) (string, int) {
	splitOnSpace := strings.Split(instruction, " ")
	operation := splitOnSpace[0]
	sign := 1

	absOperand, _ := strconv.Atoi(splitOnSpace[1][1:])
	if strings.HasPrefix(splitOnSpace[1], "-") {
		sign = -1
	}
	operand := absOperand * sign

	return operation, operand
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

	return strings.Split(string(content), "\n"), nil
}
