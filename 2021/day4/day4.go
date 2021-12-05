package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type moves []int

type board [][]int

func main() {
	moves, boards, err := readAndParseInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution for part 1: %d\n", computePartOne(moves, boards))
	fmt.Printf("Solution for part 2: %d\n", computePartTwo(moves, boards))

}

func computePartOne(moves moves, boards []board) int {
	board, move := getBingoBoard(moves, boards, true)
	return sumUnmarked(board) * move
}

func getBingoBoard(moves moves, boards []board, first bool) (board, int) {
	var numWonBoards int
	wonBoards := make(map[int]bool)

	for _, move := range moves {
		for boardIdx, board := range boards {
			markBoard(board, move)
			if checkBingo(board) {
				if !wonBoards[boardIdx] {
					wonBoards[boardIdx] = true
					// Either return the first board (in case part 1) or return the last board (in case part 2)
					if (first && numWonBoards == 0) || numWonBoards == len(boards)-1 {
						return board, move
					} else {
						numWonBoards++
					}
				}
			}
		}
	}
	return nil, -1
}

func computePartTwo(moves moves, boards []board) int {
	board, move := getBingoBoard(moves, boards, false)
	return sumUnmarked(board) * move

}

func readAndParseInput(inputFilename string) (moves, []board, error) {
	var boards []board
	var moves moves
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return nil, nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, nil, err
	}

	for _, m := range strings.Split(strings.Split(string(content), "\n")[0], ",") {
		i, _ := strconv.Atoi(m)
		moves = append(moves, i)
	}

	var board board

	for _, row := range strings.Split(string(content), "\n")[1:] {
		if len(row) > 0 {
			//fmt.Println(row)
			var rowInts []int
			for _, m := range strings.Fields(row) {
				mInt, _ := strconv.Atoi(string(m))
				rowInts = append(rowInts, mInt)
			}
			board = append(board, rowInts)
		} else {
			boards = append(boards, board)
			board = nil
		}
	}

	boards = append(boards, board)
	return moves, boards[1:], err
}

func checkBingo(board board) bool {
	for _, row := range board {
		if sumSlice(row) == -5 {
			return true
		}
	}

	for ci := 0; ci < 5; ci++ {
		var column []int
		for ri := 0; ri < 5; ri++ {
			column = append(column, board[ri][ci])
		}

		if sumSlice(column) == -5 {
			return true
		}
	}

	return false
}

func markBoard(board board, move int) {
	for i, row := range board {
		for j, num := range row {
			if num == move {
				board[i][j] = -1
			}
		}
	}
}

func sumUnmarked(board board) int {
	var sum int

	for _, row := range board {
		for _, num := range row {
			if num != -1 {
				sum += num
			}
		}
	}

	return sum
}

func sumSlice(s []int) int {
	var sum int

	for _, n := range s {
		sum += n
	}

	return sum
}
