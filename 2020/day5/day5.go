package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Problem: https://adventofcode.com/2020/day/5
// Input: https://adventofcode.com/2020/day/5/input

const LEN_OF_F_B = 7

func main() {
	boardingPasses, err := readBoardingPasses("input.txt")
	if err != nil {
		panic(fmt.Errorf("readBoardingPasses threw %v\n", err))
	}

	fmt.Println(computePartOne(boardingPasses))
	fmt.Println(computePartTwo(boardingPasses))
}

func computePartOne(boardingPasses []string) int {
	maxSeatId := 0

	for _, boardingPass := range boardingPasses {
		currSeatId := computeSeatID(boardingPass)
		if currSeatId > maxSeatId {
			maxSeatId = currSeatId
		}
	}
	return maxSeatId
}

func computePartTwo(boardingPasses []string) int {
	seatIds := make([]int, len(boardingPasses))
	for _, boardingPass := range boardingPasses {
		seatIds = append(seatIds, computeSeatID(boardingPass))
	}
	sort.Ints(seatIds)
	for i := 0; i < len(seatIds)-1; i++ {
		if seatIds[i+1]-seatIds[i] == 2 {
			return seatIds[i] + 1
		}
	}
	return -1
}

func computeSeatID(boardingPass string) int {
	return (computeSeatRow(boardingPass) * 8) + computeSeatCol(boardingPass)
}

func computeSeatRow(boardingPass string) int {
	seatRow, _ := strconv.ParseInt(convertBoardingPassToBinary(boardingPass)[:LEN_OF_F_B], 2, 64)
	return int(seatRow)
}

func computeSeatCol(boardingPass string) int {
	seatCol, _ := strconv.ParseInt(convertBoardingPassToBinary(boardingPass)[LEN_OF_F_B:], 2, 64)
	return int(seatCol)
}

// After thinking about it, I realized that the boardingPass is just a binary number encoded as a string of B and R (1) and F and L (0)
func convertBoardingPassToBinary(boardingPass string) string {
	bReplaced := strings.ReplaceAll(boardingPass, "B", "1")
	rReplaced := strings.ReplaceAll(bReplaced, "R", "1")
	fReplaced := strings.ReplaceAll(rReplaced, "F", "0")

	return strings.ReplaceAll(fReplaced, "L", "0")
}

func readBoardingPasses(inputFileName string) ([]string, error) {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n"), nil
}
