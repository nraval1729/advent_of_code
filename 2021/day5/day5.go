package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type point struct {
	X int
	Y int
}

func newPoint(x, y int) point {
	return point{
		X: x,
		Y: y,
	}
}

type lineSegment struct {
	P point
	Q point

	M int
	C int
}

func newLineSegment(x1, y1, x2, y2 int) lineSegment {
	var m, c int
	if x1 != x2 {
		m = (y1 - y2) / (x1 - x2)
		c = (x1*y2 - y1*x2) / (x1 - x2)
	} else {
		m, c = -1, -1
	}
	return lineSegment{
		P: newPoint(x1, y1),
		Q: newPoint(x2, y2),

		M: m,
		C: c,
	}
}

func (l lineSegment) getPointsBetween(doDiagonal bool) []point {
	var points []point

	if l.P.X == l.Q.X {
		// Line parallel to the Y-Axis

		y1, y2 := l.P.Y, l.Q.Y
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		for y := y1; y <= y2; y++ {
			points = append(points, newPoint(l.P.X, y))
		}
	} else if l.P.Y == l.Q.Y {
		// Line parallel to the X-Axis

		x1, x2 := l.P.X, l.Q.X
		if x1 > x2 {
			x1, x2 = x2, x1
		}

		for x := x1; x <= x2; x++ {
			points = append(points, newPoint(x, l.P.Y))
		}
	} else if doDiagonal{
		var xIncr, yIncr int
		if l.P.X < l.Q.X && l.P.Y < l.Q.Y {
			xIncr, yIncr = 1, 1
		} else if l.P.X > l.Q.X && l.P.Y > l.Q.Y {
			xIncr, yIncr = -1, -1
		} else if l.P.X < l.Q.X && l.P.Y > l.Q.Y {
			xIncr, yIncr = 1, -1
		} else {
			xIncr, yIncr = -1, 1
		}

		i, j := l.P.X, l.P.Y

		for i != l.Q.X && j != l.Q.Y {
			points = append(points, newPoint(i, j))
			i += xIncr
			j += yIncr
		}
		points = append(points, newPoint(l.Q.X, l.Q.Y))
	}
	return points
}

func main() {
	lineSegments, err := readAndParseInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution for part 1: %d\n", computePartOne(lineSegments))
	fmt.Printf("Solution for part 2: %d\n", computePartTwo(lineSegments))

}

func computePartOne(lineSegments []lineSegment) int {
	var numPoints int
	pointToCount := make(map[point]int)

	for _, lineSegment := range lineSegments {
		points := lineSegment.getPointsBetween(false)

		for _, point := range points {
			if _, exists := pointToCount[point]; exists {
				pointToCount[point] = pointToCount[point] + 1
			} else {
				pointToCount[point] = 1
			}
		}
	}

	for _, count := range pointToCount {
		if count > 1 {
			numPoints++
		}
	}

	return numPoints
}

func computePartTwo(lineSegments []lineSegment) int {
	var numPoints int
	pointToCount := make(map[point]int)

	for _, lineSegment := range lineSegments {
		points := lineSegment.getPointsBetween(true)

		for _, point := range points {
			if _, exists := pointToCount[point]; exists {
				pointToCount[point] = pointToCount[point] + 1
			} else {
				pointToCount[point] = 1
			}
		}
	}

	for _, count := range pointToCount {
		if count > 1 {
			numPoints++
		}
	}

	return numPoints
}

func readAndParseInput(inputFilename string) ([]lineSegment, error) {
	var lineSegments []lineSegment
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return nil, err
	}

	for _, pair := range strings.Split(string(content), "\n") {
		points := strings.Split(pair, " -> ")
		x1Str, y1Str := strings.Split(points[0], ",")[0], strings.Split(points[0], ",")[1]
		x2Str, y2Str := strings.Split(points[1], ",")[0], strings.Split(points[1], ",")[1]

		x1, _ := strconv.Atoi(x1Str)
		y1, _ := strconv.Atoi(y1Str)
		x2, _ := strconv.Atoi(x2Str)
		y2, _ := strconv.Atoi(y2Str)

		lineSegments = append(lineSegments, newLineSegment(x1, y1, x2, y2))
	}

	return lineSegments, nil
}
