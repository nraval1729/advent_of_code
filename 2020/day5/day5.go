package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Problem: https://adventofcode.com/2020/day/5
// Input: https://adventofcode.com/2020/day/5/input

const LEN_OF_F_B = 7

func main() {
	boardingPasses := []string{
		"FBBFFBBLLL", "FFBFFFBRLL", "FFBBBBFRRL", "FBFBBBBRLL", "BFBBBBFLLR", "FFFBBBBLRR", "BFFFFFBLLL", "BBFFFBFRRL", "FFBFFFFLLR", "BFFFBBBRRL",
		"FBFBFFFLRL", "FFFBBFBLRR", "FBFBFBFLRR", "FBBBBFBRRL", "BFFBFFBRRR", "FBBBFBBRLL", "FBFFBFBRLR", "BBFBFFFLRL", "FFBFFFFRLR", "FFBBFBFRRR", "BFBBBFBLRR", "FFBBFFFLRL", "FBBBBFFRLR", "FBBBBBBRLR", "FFBFBFBLLL", "BBFBFBBLLL", "FFFFFBBRRL", "FBFFBFBRRR", "FFFBFFFRLL", "BFBFBFFLLL", "BFBFFBFLLL", "FFFFBFFRRL", "FBFFFFFRLR", "FBBFFBBLLR", "BFFFFFBLRL", "BFBFFFBLLR", "FBBBBBBLLL", "BBFBFBFLLL", "FFBFFBFLRR", "BBFFFBBRLL", "FFBFFFFLLL", "FBBFFBFLLL", "FFFBBBFRRR", "BFBBFBFRLL", "FBBFBFFRRL", "FBFBBFBLRR", "FFBBBBBRLR", "FFBFBFFLLL", "FBFFFBFRRL", "BFFBBFFRLL", "BFFBBFFLLR", "BBFFBFFRRL", "FBFBFBFLLR", "BBFFBFBLRL", "BFBFBFBLRR", "FFBBFBBRRL", "BFBBBFBRLR", "FBFFFFBLLR", "BFBFFBBLRR", "BBFBFBFRRR", "FFBFBBBRRL", "BFBFBBFRLR", "FBFBBBBRLR", "BFFFFBBRRR", "BBFFFBFLLR", "FBFFFFBRRL", "BBFFBFBLLL", "BFFFBBBLLL", "FFBFFFBLRR", "BFBFBBBRRR", "FBFFBFFRRR", "BFFBFFBLLL", "BFBBFFFLRL", "BBFBFFBRLR", "FBBFFFFLLL", "FBFFFFBLLL", "BBFFFFBRLR", "BBFBFFBRRL", "FBBFBFBRRR", "FBFFBFFLRR", "FFBBBFBLLR", "FFFFFBFRRR", "BFBBBBBLLR", "BBFBFFFRLL", "BFFBFBBLRR", "BFFBBFBRLL", "FBFFBFFRLR", "FBFBBFFLLR", "FFFFFBFRRL", "FFBFFBFRLR", "FBFBBBFLLL", "FFBFBBFLLR", "FBFFBBBRLR", "FFFBFBFLLR", "BFBBFBBLRL", "BFFFFBFLLL", "BFBFFBFLRL", "BBFFBFBLLR", "BFFFBFBRRL", "BBFBFFBLLR", "FFBBFFFRLL", "FFFFFBBRLR", "FFBFBFBRLL", "FBBBBFBLRR", "FFFBFFFRRL", "FFFFBFFLRL", "FFFFBFFRLR", "BFFFFFFRRR", "BBFFBFBRRL", "BBFBFFFLRR", "FFBBBFFLLR", "FBFBFBBRLR", "FBBFBBBLRR", "BFFBBBFRLL", "BFBBFBBLLL", "FBFFFBBLLR", "FBBBBFFRRL", "BFBBFFFLLR", "BFBBBFFLRL", "FFBFFBBRLL", "FBFBFFBRRR", "FBBFBBBLRL", "BFFFBFFLRL", "BFFBFFBLRR", "BFBFBBFLRL", "FBFBBFFRLL", "BFFFFBBRRL", "FFBBBBBLRL", "FBBFFBFRRL", "BBFFBFBRLL", "FBFFFFFLRL", "FFFBFBFRLR", "FFBBBBFLLR", "FFFBBBFRLR", "FFBBFFBRRR", "FFFBFFBLRL", "FBFBFBBLLL", "BFFFFBBLRR", "BBFFBBFLLR", "FFFFFFBRLL", "FBBBBBBLRR", "FFFFBFBLLR", "BBFFBBFRRL", "BFBBBBBLRR", "BFBBFBBLRR", "BBFFFFBRRR", "BFBBBFBLLR", "BBFFBBBRRR", "FBBFBFBRLL", "FFFFFBBLRR", "BBFBFBBLLR", "FBFFFBFRLR", "BBFFBFFLLR", "BBFBFFFRLR", "FFFFBFBLRL", "FFBBBBBLRR", "FBBBFFBLRR", "FBBFBBFRLL", "FBBFBFBRLR", "FFBFBBBLRL", "FBFFBBBLLL", "FFFBBBBRRR", "FFBFFBFLRL", "BFFFBBFRLL", "BFFFFBBLLR", "BFFFFFBRRL", "BFFBFBBRRL", "FFBFFFBRRL", "BFFBFFFLRL", "BFFBBBFLLL", "FFBBFFFLRR", "FBBBFFBRLL", "FFBBFFFRLR", "BFBFBFFRLL", "FFFFBBFRRR", "BFBFBBFRRL", "FBFBBBFLRL", "FFFBFBBRLL", "FFBBBFFLLL", "BBFFFBBLRL", "FBBBFFBRRR", "BFBFBFBLLL", "FFBFFBBRLR", "FFBFBBBRLR", "FFFFFFBRLR", "BFFBBFFLRL", "FFBFFBFRRR", "FFBFBFFRLL", "FBFBFFBLLR", "FBFBFBBRRR", "FFFBBFBRRR", "FBBBBFBLLR", "FFFFBFBRLR", "FFFBBBFRLL", "FBFFBFBLLL", "FBFBFFFLRR", "FFFBBFFLLL", "BFBFFFFLLL", "BBFFFBFLLL", "FFFFBFBRRL", "FFBBBFBRRL", "BFBFFBBRLR", "FFFBBFFLRL", "FFFBFBBLLR", "FFBBBFFRLR", "FBBBBFBLRL", "BFFFBBFLLL", "BFFBFBFLLR", "FBFFBBFLRR", "BFFBBBBLLR", "FBBBFFFRRR", "FBFFBFBLRL", "BFFFFFFRLR", "FBBBBFBRRR", "FFBBBBFLRL", "BBFBFFBLRL", "BFBFBBBLRR", "FFFFFBFLLR", "FBBFFBFLLR", "BFFFFBBRLL", "FFBFBBBLRR", "FFFFBFFLLR", "BBFFFFFRLR", "FBBBBFFRLL", "BFFFBFBLRR", "FBFBBBBLLR", "BFFBFFBRLL", "BFBFBFFLRR", "BFFFFBFLRR", "FBBFFBFLRR", "BFFBBFBLLR", "FBBFFFBLRL", "FBFBFBBRRL", "FBFFBFFRRL", "BFBBFBFLLR", "FFFFBFFLRR", "BFBFFFBLRL", "BFFBBBBRRR", "FFFFBBBLRR", "FFFBBFBLLL", "FFBBFBFLLR", "BFBBFBFLLL", "FBBFBFFLLR", "FFFBBFFRRR", "FFBBBBBRRR", "FFFBFFFRLR", "BFFBFBFLRR", "FBFFBFBRLL", "FFFFBFFLLL", "FFBBBFFRLL", "FBFBFFBLLL", "FBFBBBFLLR", "BFBBBFFRLL", "BFBBBBBRRL", "FFFFFFBRRL", "FBBFBBFLRL", "BFFBFBBRLL", "FFFFFBFLRL", "BFBFFFBLRR", "BBFFFBFLRR", "BFBFBBBLRL", "FBFFFFBRLR", "BFBFBBBLLL", "BFFFFFFRRL", "FBBFFFBRRR", "FBBFBFFLRL", "BBFFFFFLRL", "BFFFFBBLRL", "FBFBBFFRRR", "FBBFFFFLRR", "FFFBFFFLRR", "FBFBBFFRRL", "BFFBFFBRRL", "BFFBFBFRRL", "FFBBFBBLLR", "FBFFBFFRLL", "BFBFFBBLLL", "FFBBFFBLLL", "BBFFFBBRRL", "BFFFFFFLLL", "FFBBBBFRLL", "BFBBBBFRRR", "BFFBBBBRLR", "FBFFBBFRLL", "BBFBFFBLRR", "FBFFBBBLRR", "BBFBFFFRRL", "BFBFFFFLRR", "BFBFFBBRRR", "FBBFFFBLLL", "BFFBBBFRLR", "BFFFFBFRLL", "FBFBBBBRRL", "BBFFBBFLRL", "FBBBBFFLLR", "BFBBFFBRRR", "FBBFFBBLRR", "FBBBBBBRRL", "BBFFFBBLLR", "BFBBBBFLRL", "BFFBFFFRLR", "FFFFBFBRRR", "FFFFBFBRLL", "BFFFFFFRLL", "BFFFBBBRLL", "FFFFFBBLLL", "FBFFFBFLRR", "FBFBFBFLLL", "BFBBFFBRLL", "BFFFBBFRLR", "BFBBFFFRLL", "FFFBBFBLRL", "BBFFFFFLRR", "BBFFFBFRLR", "BFBBBBFLRR", "FBFFFBBLLL", "BBFFBFBRRR", "FBFBFBBLRR", "FBBFBFFRLR", "FBFBBBBLLL", "FFFBBFBLLR", "FFFFBBBRRL", "FFBBBBFRLR", "FFFBFFBRRL", "BFBBFBBLLR", "FFBFFFFLRL", "BFFBFBFRLR", "FFBBFBBLLL", "BFFBBBFLRL", "FFFFFBFRLL", "BFFBFFBRLR", "BFFFBFFRRL", "BFBFBBBRLL", "BFFFFFBLRR", "FBBFFBBRLL", "FFFFBBBLRL", "BBFFFFBLRR", "BFBFBBBLLR", "BFFFBFFRLR", "BFBBFBFRRR", "BFFBBBFLRR", "FFBFBBFRRR", "FBBBFBFRLR", "BFBFFBBRLL", "BBFFBFFRLL", "BFFFBFBLLR", "FBFBBFFLRL", "BFFFFBBLLL", "BFBBFBBRRL", "BFFFBFFLLL", "BFBBBFBLLL", "FFBBFBFLRR", "FFBBFBFRLR", "BFBBBFFRRR", "BBFBFFFLLL", "BFFFFFBRLL", "FBBFBBFRRR", "BFFFBBFLRR", "BFFBBBFRRL", "FFBFFFFRRL", "FBFFBFBLRR", "FFBFBFBRLR", "BBFFBFFLRR", "BFFBBFFLLL", "FFBBBFBRRR", "BBFFBFFRRR", "FBBFFFFRRR", "FBBBBFFLRR", "FBBFBBBRRR", "FFBFBFBLRL", "FBBFFFBLRR", "FBFFBFFLRL", "BFBBFFFRRR", "FBFBBBFRLR", "BFFFFFBRRR", "FBFBFFFRRR", "BBFFFBFLRL", "FFBFBFBLLR", "FBFFFBBRRL", "FBBBFFBLLL", "BFBFBFFRRR", "FBFFBBBRLL", "FFBFFBBLLR", "FFBBFBBRLL", "BFBFBBFRLL", "BFBBFFBLLL", "BFFBBBFRRR", "BFBFBFBRRL", "BFFFFFFLLR", "BFFFBBBRRR", "FBBBFBFRRR", "FFFBFFBRLL", "BFFFFBFLRL", "FBFBBFFRLR", "FBBBBBFLRL", "FFFBBBFRRL", "BFBFFFFLLR", "FBFFFBBLRR", "FBBFBBBRRL", "FFFBBBFLRR", "BFBFFBFRRL", "BFFFBBFRRL", "FBBBFFBLLR", "FFBBFBBLRL", "BBFBFBFRLR", "FFBFFBFLLL", "FFFFBFBLLL", "BBFFFFBRLL", "FBFFFBBRLR", "FFBBBFFRRR", "FFFBFFFRRR", "FFBFBBBLLR", "FFBFFFBRLR", "FFBBBFBLRR", "BFFBBBBLRL", "BFFFFBFRRR", "BBFBFBFRLL", "FBFBBBFLRR", "FFFFFBFLLL", "BFBBBFFLLR", "BFFBFBFRRR", "BFBBFFBLRR", "BBFFFBBRLR", "FBBBBBBLRL", "BFFBBBBRRL", "FFBFBFFRLR", "BFFFFFBLLR", "BBFFFFBRRL", "FFFBBFFLRR", "BFFFFBFRRL", "BBFFBFFLRL", "BBFBFBFLRR", "FBFBBBBLRR", "BFFBBFBRRR", "FFFBBBBRLL", "FBFBBBFRRL", "FFFFFFBRRR", "BFFFFFFLRL", "FFBBFFFLLR", "FBBFBFBLLL", "FBBFBBFRLR", "BFFBFFBLLR", "FBFBBFBRLL", "FBFBFBFLRL", "FFBFFFBLLR", "FBBFFBBLRL", "FFFBBBBLLR", "BFFBBBBLRR", "BFFBFBBLLR", "FFBBFFBRLR", "FBBBBFFLLL", "BFFFFFBRLR", "BFBBBBFRLL", "FFFBBFFRLR", "BFFFBFBLLL", "BBFFFFFRRL", "FBFFFFBRLL", "BFBFBFBLRL", "FBBFBBFLRR", "FFBFBBFLRR", "FFBFBFFLRR", "FFBFFFFRLL", "BFBFBFBRRR", "FBBBBFFLRL", "FFBFFFBLLL", "BFBBBFFRLR", "FBFBFFBRLL", "FBBBBBFRRL", "FBFFBFFLLL", "BBFFBBBLLR", "BFFFBFFRRR", "BBFFBBBRLL", "BBFFFBBLRR", "FBBFBBBRLR", "BFBFBBFRRR", "FBFFBBFLRL", "FFFFBBFLRR", "BFFFBFBRLL", "BBFFFFFRLL", "FFFBFBFRRR", "BFBBBBFLLL", "FFBBBFBRLR", "FFFFBBBRRR", "BFBFFBFRLR", "FBFFBBFLLL", "FFBFFBBLRR", "FBFBFFBLRR", "FFBFFBBRRL", "FFFBBBBRRL", "BFFBFBBLRL", "BBFFBBBLRR", "FFFBBFBRLL", "FFFBFBBLRL", "BFBFFFBRRL", "FBBFFFFLLR", "BBFBFFBLLL", "BFBFBBFLLR", "BFFBFBBRRR", "FBFFFFFLLR", "FFBBBFFLRL", "FBFFBBFRLR", "BFFFBFBLRL", "FBFBBBFRRR", "FFBFFBFRRL", "FBFBBFBRLR", "BFBFBFFRLR", "FBFFBBBRRR", "BBFBFFFLLR", "FBBFBBBLLL", "FFBBBFBLLL", "FFBFFFFRRR", "BFFBFBFLLL", "BFFBBBFLLR", "FBBBBBBRLL", "FBBFFBFRLL", "FFFBFFBLLR", "FFBBFFBLRL", "BFBBFBFLRR", "BFBFFFBLLL", "FBFFFBBRRR", "FFBFBFFLLR", "FBFFBBBRRL", "FBBBBFBLLL", "FFFFFBBLLR", "FBBFBFBLRL", "BFBBFBBRLL", "FFFFBFBLRR", "BBFFBBFLLL", "FBBBFBFRLL", "FBFBBFBLLL", "BBFBFBFLLR", "BFBFBFBRLR", "FFFFFBFRLR", "FFFFFBBRLL", "FBFBFBFRRR", "FBBBFBFLRL", "BBFFBBBLRL", "FFBBBBBRLL", "FBBBFBBRRR", "BFFFBBBLLR", "BFBBBBBRLL", "FFFFFBBLRL", "BBFFBFFRLR", "BFBFBFBRLL", "FFFFBBFLLL", "BFBFBFFRRL", "FFBBFFBLLR", "BFBBFFFLLL", "BFBBFFBRLR", "FBFBFFBRLR", "FBFFFBFRLL", "FBBFBFFLRR", "FBBFFBFRRR", "BFBFFFFRLL", "BFFFBBBLRL", "FBFBBBBRRR", "BBFFBBFLRR", "FBBBBFBRLL", "BBFFBBBRRL", "BFBBBFFLLL", "FBFFFFBLRL", "FBFFFFFRRL", "BFBBBFFLRR", "FBBFBFBLRR", "FBFBFBBLRL", "FFBFBFFRRL", "BFFBBFBRLR", "FFBFFBBLLL", "FBBBFBBLLL", "FBBBFBBRRL", "FBBBFBBLLR", "FBFFFBFLLL", "FBBBFFBRLR", "FBBFBBBRLL", "FBBFFFBLLR", "FFFBFBBRRR", "FBBFFFFLRL", "BFFFFBFRLR", "FBBBFFFLLL", "FFBBFBFLLL", "BBFFFFFRRR", "BBFFFBFRRR", "FFBFFBBLRL", "FBBBFFBLRL", "FFBBFBFRRL", "FBFFFBFLRL", "FBBBFFFLRL", "FFFFBBBRLR", "FFBBFBBRLR", "FBFBFFBRRL", "BFBFFBFRRR", "BFBFBFFLRL", "BFFBFFFLLL", "FBBFBFBRRL", "FFBBBFBLRL", "BFFFFBFLLR", "FBBBFBFLLL", "FBFFFBBLRL", "FBBFFFBRRL", "FBBBBBFLLR", "FBBBFBFRRL", "FFFBBFBRRL", "FBBBFFBRRL", "FBFBFBBRLL", "BFBBBBBLLL", "FBFBFFFLLR", "FFFBFFBRLR", "BBFFBFBLRR", "BFFFBFFLLR", "FBBFBBFRRL", "FFBBBBBLLL", "FBBFBFFRLL", "FFFBBBFLRL", "FFBBBBFLRR", "FBFFFFFRLL", "BBFFFFBLLR", "FFFFBBFRRL", "FFBBBBFRRR", "FBBBFBBLRL", "BFFBBFBLRL", "BFBBFFFRLR", "FFFBFBBRRL", "BBFBFBFRRL", "FFFBBBFLLR", "BFFFFBBRLR", "BFFBFFFLLR", "BBFBFFFRRR", "FFFFBBFRLR", "BFBBBFBRRR", "FBFBBBBLRL", "FFFFBBFLLR", "FBFBBBFRLL", "BFBBFBBRLR", "FFBBFFFRRR", "FFBBBFBRLL", "FBBFFBFLRL", "BFBFFFBRRR", "BFFFBBBRLR", "FFFFBFFRRR", "FFBBFFBLRR", "FBBFBFBLLR", "FFFBBBFLLL", "FBBBBBFLLL", "FFFFBBBLLL", "FFFBFFBLLL", "BFBBBBBRLR", "BFBFBBFLRR", "FBFBBFBRRR", "BFBBFFFLRR", "FFBBFFFLLL", "BFFFBBFLLR", "FFBBFBFRLL", "FBBBFFFRLL", "FFFBBFBRLR", "BFBFFBFLRR", "FBFBBFBLLR", "FBFFFFBRRR", "FFFBFBBLLL", "FBFBFFFRRL", "FFFBFBFLRR", "FFBBBBBRRL", "FFFBFFFLLR", "BFFFBFFRLL", "BFBBBBBLRL", "FBBBFBFLRR", "FFFFBBBLLR", "BFFFBFFLRR", "FBBBBBFLRR", "FBBBBBFRLR", "BFFBFFFRRR", "FFBFBBFLRL", "FFBFBBFRRL", "FFBFFBBRRR", "FFFFBBFRLL", "BBFBFBFLRL", "FBFBBFFLLL", "BBFFFFBLLL", "FBBBBFFRRR", "BFFBFBFRLL", "BFBBFBFRRL", "FFBFBBFLLL", "FBFFFBFRRR", "FFBFFBFLLR", "BFBBFFBRRL", "FBBFFBBRRR", "FFBFBBBRLL", "FBBBBBBLLR", "FBFBBFBRRL", "BFFBFFFRLL", "BBFFBBFRRR", "FBBBBBFRLL", "FBBFFFFRLR", "FBBFFBBRLR", "BFBFBFBLLR", "FBFFBFBRRL", "FFFBBFFRRL", "FFBFBBFRLR", "BFFBFBBRLR", "BFFFBFBRRR", "BFBBBFFRRL", "FBBBFBBLRR", "FFFBFBFRLL", "FFBBFFBRLL", "FFFBBBBRLR", "FBBBFBFLLR", "BFBFBBBRRL", "FBBBFFFRLR", "FBFFFFFLRR", "FFBFBFFRRR", "FFFBFBBLRR", "FBBFBFFLLL", "FFBFFBFRLL", "FFBBBFFLRR", "FFBFBFBRRR", "FFFBBBBLRL", "FFBFBFFLRL", "FBFBBFBLRL", "FFFBFBFLRL", "BFBFFBFLLR", "BBFFFBFRLL", "FFBBBBBLLR", "BBFFFBBLLL", "BBFBFFBRRR", "BFBFFFBRLL", "FFFBFBFRRL", "FFBFFFFLRR", "BBFFFFFLLL", "FBBBFFFLRR", "BFFFBBFRRR", "FBBFBBBLLR", "BFBFFFFRRL", "FFBBFFBRRL", "BFBFFBBLRL", "BBFFFBBRRR", "FFFBBFFRLL", "BFFBBFBRRL", "BFFFFFFLRR", "BFBBBBFRLR", "FBBFFBFRLR", "BFFBBFFRRL", "BFFBFFFRRL", "FFFFBBBRLL", "FBFBFFFRLR", "BBFFBBFRLR", "BBFFBBFRLL", "BFFFBBFLRL", "FBFFBFFLLR", "BFBBFBFLRL", "FBFBFBFRLL", "BFBFBBBRLR", "BFBFFBBLLR", "FBBBBFBRLR", "FBFFFFBLRR", "FBFFBBFRRR", "BBFFBFFLLL", "FFFBFBFLLL", "BFFBBFFRRR", "FBBFBFFRRR", "BBFFBBBRLR", "FFFBFFBLRR", "FBFBFBFRLR", "BFBBBBBRRR", "BFFBFBFLRL", "BBFFBBBLLL", "BBFFFFFLLR", "BFBFFBFRLL", "BFBBBBFRRL", "BFBBFFBLRL", "BFFBBFBLLL", "BFFBBBBLLL", "BFBFBBFLLL", "FBFBFFFRLL", "FBFBFBFRRL", "BFFBFBBLLL", "FFBBFBBRRR", "FFFBFFBRRR", "FBFFFBFLLR", "FBFFFFFRRR", "FBBFBBFLLL", "FFFBFBBRLR", "FBFFFFFLLL", "FBBFFFFRRL", "BFBBFFBLLR", "BBFBFBBLRL", "FBBFBBFLLR", "BFBFFFFLRL", "FFBBFBBLRR", "BFBBBFBLRL", "BFFBBFFRLR", "FFFFBFFRLL", "BFFBBFFLRR", "FBBBFBBRLR", "FBFBFFFLLL", "FFBBBFFRRL", "BFBBFBFRLR", "FBFFFBBRLL", "FFFBFFFLRL", "BFBBBFBRRL", "FBBFFFBRLR", "FFBBBBFLLL", "FBBBFFFLLR", "FBFFBBFLLR", "FFBFBBFRLL", "BBFFBFBRLR", "FBFFBBBLLR", "FFBFBFBRRL", "BFFBFFBLRL", "FBFFBFBLLR", "FFBFBFBLRR", "FFBBFBFLRL", "FFBBFFFRRL", "BFBFBFFLLR", "FBBBBBBRRR", "BFBFFFFRLR", "BFFBBBBRLL", "FFFFFBBRRR", "FFBFFFBRRR", "FFFBFFFLLL", "BFBFFFFRRR", "FBBBBBFRRR", "FBBBFFFRRL", "FBBFFFFRLL", "FFFFFBFLRR", "FFBFBBBLLL", "BFBFFBBRRL", "FFFFBBFLRL", "BFBBBFBRLL", "FFBFFFBLRL", "FBBFFFBRLL", "BFFBBFBLRR", "FFFBBFFLLR", "FBBFFBBRRL", "FFBFBBBRRR", "FBFBFFBLRL", "BFBBFBBRRR", "BBFFFFBLRL", "FBFBFBBLLR", "BFFFBBBLRR", "BFBFFFBRLR", "FFFBBBBLLL", "BFFBFFFLRR", "BFBBFFFRRL", "FBFFBBBLRL", "FBFBBFFLRR", "FBFFBBFRRL", "BBFBFFBRLL",
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