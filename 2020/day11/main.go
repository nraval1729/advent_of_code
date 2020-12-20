package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Problem: https://adventofcode.com/2020/day/11
// Input: https://adventofcode.com/2020/day/11/input
func main() {
	seats, err := readAndParseInput("input.txt")
	if err != nil {
		panic(fmt.Errorf("readAndParseInput returned %v\n", err))
	}
	fmt.Println(computePartOne(seats))
	fmt.Println(computePartTwo(seats))
}

func computePartOne(seats [][]string) int {
	numOccupiedSeatsCurr, numOccupiedSeatsPrev := 0, 0
	for {
		seatsCopy := copySeats(seats)
		for ri, r := range seats {
			for ci, c := range r {
				if c == "L" && getNumAdjacentOccupiedSeats(ri, ci, seats) == 0 {
					seatsCopy[ri][ci] = "#"
					numOccupiedSeatsCurr++
				}
				if c == "#" && getNumAdjacentOccupiedSeats(ri, ci, seats) >= 4 {
					seatsCopy[ri][ci] = "L"
				}
			}
		}
		if numOccupiedSeatsCurr == numOccupiedSeatsPrev {
			return countOccupiedSeats(seatsCopy)
		}
		numOccupiedSeatsPrev = numOccupiedSeatsCurr
		numOccupiedSeatsCurr = 0
		seats = seatsCopy
	}
}

func copySeats(seats [][]string) [][]string {
	seatsCopy := make([][]string, len(seats))
	copy(seatsCopy, seats)
	for i := range seats {
		seatsCopy[i] = make([]string, len(seats[i]))
		copy(seatsCopy[i], seats[i])
	}
	return seatsCopy
}

func countOccupiedSeats(seats [][]string) int {
	count := 0

	for _, row := range seats {
		for _, col := range row {
			if col == "#" {
				count++
			}
		}
	}

	return count
}

func computePartTwo(seats [][]string) int {
	numOccupiedSeatsCurr, numOccupiedSeatsPrev := 0, 0
	for {
		seatsCopy := copySeats(seats)
		for ri, r := range seats {
			for ci, c := range r {
				if c == "L" && getNumAdjacentOccupiedSeatsDirectionally(ri, ci, seats) == 0 {
					seatsCopy[ri][ci] = "#"
					numOccupiedSeatsCurr++
				}
				if c == "#" && getNumAdjacentOccupiedSeatsDirectionally(ri, ci, seats) >= 5 {
					seatsCopy[ri][ci] = "L"
				}
			}
		}
		if numOccupiedSeatsCurr == numOccupiedSeatsPrev {
			return countOccupiedSeats(seatsCopy)
		}
		numOccupiedSeatsPrev = numOccupiedSeatsCurr
		numOccupiedSeatsCurr = 0
		seats = seatsCopy
	}
}

func printSeats(seats [][]string) {
	for _, row := range seats {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

func getNumAdjacentOccupiedSeats(ri, ci int, seats [][]string) int {
	numOccupiedSeats := 0
	rmax, cmax := len(seats), len(seats[0])

	if isWithinBounds(ri-1, ci, rmax, cmax) && seats[ri-1][ci] == "#" {
		numOccupiedSeats++
	}
	if isWithinBounds(ri+1, ci, rmax, cmax) && seats[ri+1][ci] == "#" {
		numOccupiedSeats++
	}
	if isWithinBounds(ri, ci-1, rmax, cmax) && seats[ri][ci-1] == "#" {
		numOccupiedSeats++
	}
	if isWithinBounds(ri, ci+1, rmax, cmax) && seats[ri][ci+1] == "#" {
		numOccupiedSeats++
	}
	if isWithinBounds(ri-1, ci-1, rmax, cmax) && seats[ri-1][ci-1] == "#" {
		numOccupiedSeats++
	}
	if isWithinBounds(ri-1, ci+1, rmax, cmax) && seats[ri-1][ci+1] == "#" {
		numOccupiedSeats++
	}
	if isWithinBounds(ri+1, ci+1, rmax, cmax) && seats[ri+1][ci+1] == "#" {
		numOccupiedSeats++
	}
	if isWithinBounds(ri+1, ci-1, rmax, cmax) && seats[ri+1][ci-1] == "#" {
		numOccupiedSeats++
	}
	return numOccupiedSeats

}

func getNumAdjacentOccupiedSeatsDirectionally(ri, ci int, seats [][]string) int {
	numOccupiedSeats := 0
	rmax, cmax := len(seats), len(seats[0])

	// up
	numOccupiedSeats += getNumOccupiedSeatsInDirection(ri, ci, rmax, cmax, -1, 0, seats)

	// down
	numOccupiedSeats += getNumOccupiedSeatsInDirection(ri, ci, rmax, cmax, 1, 0, seats)

	// right
	numOccupiedSeats += getNumOccupiedSeatsInDirection(ri, ci, rmax, cmax, 0, 1, seats)

	// left
	numOccupiedSeats += getNumOccupiedSeatsInDirection(ri, ci, rmax, cmax, 0, -1, seats)

	// up-left
	numOccupiedSeats += getNumOccupiedSeatsInDirection(ri, ci, rmax, cmax, -1, -1, seats)

	// up-right
	numOccupiedSeats += getNumOccupiedSeatsInDirection(ri, ci, rmax, cmax, -1, 1, seats)

	// down-right
	numOccupiedSeats += getNumOccupiedSeatsInDirection(ri, ci, rmax, cmax, 1, 1, seats)

	// down-left
	numOccupiedSeats += getNumOccupiedSeatsInDirection(ri, ci, rmax, cmax, 1, -1, seats)


	return numOccupiedSeats

}

func getNumOccupiedSeatsInDirection(ri, ci, rmax, cmax, delX, delY int, seats [][]string) int {
	count := 0

	for {
		ri += delX
		ci += delY

		if !isWithinBounds(ri, ci, rmax, cmax) || seats[ri][ci] == "L"{
			return count
		}
		if seats[ri][ci] == "#" {
			count++
			return count
		}
	}
}

func isWithinBounds(ri, ci, rmax, cmax int) bool {
	return ri >= 0 && ri < rmax && ci >= 0 && ci < cmax
}

func readAndParseInput(inputFilename string) ([][]string, error) {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, err
	}

	input := strings.Split(string(content), "\n")
	seats := make([][]string, len(input))

	for i, row := range input {
		seats[i] = strings.Split(row, "")
	}
	return seats, nil
}
