package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type signalEntry struct {
	Patterns          []string
	OutputValueDigits []string
}

func main() {
	heightMap, err := readAndParseInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution for part 1: %d\n", computePartOne(heightMap))

	fmt.Printf("Solution for part 2: %d\n", computePartTwo(heightMap))

}

func computePartOne(heightMap [][]int) int {
	var riskLevel int

	for ri, row := range heightMap {
		for ci, _ := range row {
			if isLowPoint(ri, ci, heightMap) {
				riskLevel = riskLevel + 1 + heightMap[ri][ci]
			}
		}
	}
	return riskLevel
}

type point struct {
	ri, ci, val int
}
type basin struct {
	points map[point]bool
}

func (b *basin) size() int {
	return len(b.points)
}

func (b *basin) String() string {
	return fmt.Sprintf("%+v\n", b.points)
}

func newBasin() *basin {
	return &basin{
		points : make(map[point]bool),
	}
}

func (b *basin) containsPoint(p point) bool {
	_, exists := b.points[p]

	return exists
}

func (b *basin) addPoint(p point) {
	b.points[p] = true
}

func computePartTwo(heightMap [][]int) int {
	var basins []*basin
	seen := make(map[point]bool)
	const border = 9

	for i, row := range heightMap {
		for j, _ := range row {
			if _, exists := seen[point{i, j, heightMap[i][j]}]; !exists && heightMap[i][j] != border {
				// Start search
				currBasin := newBasin()
				currStack := []point{{i, j, heightMap[i][j]}}
				for len(currStack) != 0 {
					currPoint := currStack[0]
					currBasin.addPoint(currPoint)
					seen[currPoint] = true
					currStack = currStack[1:]

					for _, neighbourPoint := range getNeighbours(currPoint.ri, currPoint.ci, heightMap) {
						if neighbourPoint.val != border && !currBasin.containsPoint(neighbourPoint) {
							currBasin.addPoint(neighbourPoint)
							seen[neighbourPoint] = true
							currStack = append(currStack, neighbourPoint)
						}
					}
				}

				basins = append(basins, currBasin)
			}
		}
	}
	var basinSizes []int

	for _, basin := range basins {
		basinSizes = append(basinSizes, basin.size())
	}

	sort.Ints(basinSizes)

	largest3Basins := basinSizes[len(basinSizes)-3:]
	return largest3Basins[0] * largest3Basins[1] * largest3Basins[2]
}

func isLowPoint(ri, ci int, heightMap [][]int) bool {
	for _, neighbour := range getNeighbours(ri, ci, heightMap) {
		if neighbour.val <= heightMap[ri][ci] {
			return false
		}
	}

	return true
}

func getNeighbours(ri, ci int, heightMap [][]int) []point {
	var neighbours []point

	if 0 <= ri+1 && ri+1 < len(heightMap) {
		neighbours = append(neighbours, point{ri+1, ci, heightMap[ri+1][ci]})
	}
	if 0 <= ri-1 && ri-1 < len(heightMap) {
		neighbours = append(neighbours, point{ri-1, ci, heightMap[ri-1][ci]})
	}
	if 0 <= ci+1 && ci+1 < len(heightMap[0]) {
		neighbours = append(neighbours, point{ri, ci+1, heightMap[ri][ci+1]})
	}
	if 0 <= ci-1 && ci-1 < len(heightMap[0]) {
		neighbours = append(neighbours, point{ri, ci-1, heightMap[ri][ci-1]})
	}

	return neighbours
}

func readAndParseInput(inputFilename string) ([][]int, error) {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, err
	}

	var heightMap [][]int

	for _, row := range strings.Split(string(content), "\n") {
		var rowMap []int
		for _, h := range strings.Split(row, "") {
			height, _ := strconv.Atoi(h)
			rowMap = append(rowMap, height)
		}

		heightMap = append(heightMap, rowMap)
	}

	return heightMap, nil

}
