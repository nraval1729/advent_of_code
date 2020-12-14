package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Problem: https://adventofcode.com/2020/day/14
// Input: https://adventofcode.com/2020/day/14/input
func main() {
	program, err := readAndParseInput("input.txt")
	if err != nil {
		panic(fmt.Errorf("readAndParseInput returned %v\n", err))
	}
	fmt.Println(computePartOne(program))
	fmt.Println(computePartTwo(program))
}

func computePartOne(program []string) int {
	var mask string
	memAddressToValue := make(map[int]string)

	for _, instruction := range program {
		lhs, rhs := parseProgramInstruction(instruction)
		if isMaskInstruction(lhs) {
			mask = rhs
		} else {
			memAddress, value := parseMemoryInstruction(lhs, rhs)
			maskedValueBinary := applyMask(mask, convertToBinary(value))

			memAddressToValue[memAddress] = maskedValueBinary
		}
	}

	return sumAllMemoryAddresses(memAddressToValue)
}

func sumAllMemoryAddresses(memAddressToValue map[int]string) int {
	sum := 0
	for _, v := range memAddressToValue {
		d, _ := strconv.ParseInt(v, 2, 64)
		sum += int(d)
	}
	return sum
}

func computePartTwo(program []string) int {
	var mask string
	memAddressToValue := make(map[int]string)

	for _, instruction := range program {
		lhs, rhs := parseProgramInstruction(instruction)
		if isMaskInstruction(lhs) {
			mask = rhs
		} else {
			memAddress, value := parseMemoryInstruction(lhs, rhs)
			valueBinary := convertToBinary(value)
			maskedMemoryAddress := applyMemoryAddressDecoderMask(mask, convertToBinary(memAddress))

			for _, memAddress := range generateMemoryAddressesFrom(maskedMemoryAddress) {
				memAddressToValue[memAddress] = valueBinary
			}
		}
	}

	return sumAllMemoryAddresses(memAddressToValue)
}

func parseProgramInstruction(instruction string) (string, string) {
	s := strings.Split(instruction, " = ")
	return s[0], s[1]
}

func isMaskInstruction(lhs string) bool {
	return lhs == "mask"
}

func parseMemoryInstruction(lhs string, rhs string) (int, int) {
	s := strings.Split(lhs, "[")
	ss := strings.Split(s[1], "]")
	d, _ := strconv.Atoi(ss[0])

	v, _ := strconv.Atoi(rhs)
	return d, v
}

func convertToBinary(d int) string {
	return fmt.Sprintf("%036b", d)
}

func applyMask(mask, v string) string {
	vCopy := strings.Split(v, "")
	for i := 0; i < len(mask); i++ {
		if string(mask[i]) != "X" {
			vCopy[i] = string(mask[i])
		}
	}
	return strings.Join(vCopy, "")
}

func applyMemoryAddressDecoderMask(mask, v string) string {
	vCopy := strings.Split(v, "")
	for i := 0; i < len(mask); i++ {
		if string(mask[i]) != "0" {
			vCopy[i] = string(mask[i])
		}
	}
	return strings.Join(vCopy, "")
}

func generateMemoryAddressesFrom(maskedMemoryAddress string) []int {
	var generatedAddresses []int
	for _, c := range generateMemoryAddressCombinations(maskedMemoryAddress) {
		d, _ := strconv.ParseInt(strings.Join(c, ""), 2, 64)
		generatedAddresses = append(generatedAddresses, int(d))
	}
	return generatedAddresses
}

func generateMemoryAddressCombinations(a string) [][]string {
	var allMemoryAddresses [][]string
	if string(a[0]) == "X" {
		allMemoryAddresses = append(allMemoryAddresses, []string{"1"})
		allMemoryAddresses = append(allMemoryAddresses, []string{"0"})
	} else {
		allMemoryAddresses = append(allMemoryAddresses, []string{string(a[0])})
	}
	for i := 1; i < len(a); i++ {
		if string(a[i]) == "X" {
			// Copy all slices to another 2D slice
			var ss [][]string
			for _, s := range allMemoryAddresses {
				sCopy := make([]string, len(s))
				copy(sCopy, s)
				ss = append(ss, sCopy)
			}
			// first append 0
			for idx, s := range allMemoryAddresses {
				allMemoryAddresses[idx] = append(s, "0")
			}

			for _, f := range ss {
				allMemoryAddresses = append(allMemoryAddresses, append(f, "1"))
			}
		} else {
			for idx, s := range allMemoryAddresses {
				allMemoryAddresses[idx] = append(s, string(a[i]))
			}
		}
	}
	return allMemoryAddresses
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
