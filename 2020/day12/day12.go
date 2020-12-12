package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// Problem: https://adventofcode.com/2020/day/12
// Input: https://adventofcode.com/2020/day/12/input
func main() {
	instructions, err := readAndParseInput("input.txt")
	if err != nil {
		panic(fmt.Errorf("readAndParseInput returned %v\n", err))
	}
	fmt.Println(computePartOne(instructions))
	fmt.Println(computePartTwo(instructions))
}

func computePartOne(instructions []string) int {
	// For simplicity assume starting position is the origin
	currX, currY := 0, 0

	// starting direction is east
	dir := "E"

	for _, instruction := range instructions {
		action, value := parseInstruction(instruction)
		if action == "F" {
			dirX, dirY := getCurrentDirectionVector(dir, 1)
			currX += dirX * value
			currY += dirY * value
		} else if action == "R" || action == "L" {
			dir = getUpdatedDirection(dir, action, value)
		} else {
			deltaX, deltaY := parseDirection(action, value)
			currX += deltaX
			currY += deltaY
		}
	}
	return int(math.Abs(float64(currX))) + int(math.Abs(float64(currY)))
}

func computePartTwo(instructions []string) int {
	// For simplicity assume starting position is the origin
	currX, currY := 0, 0
	wpX, wpY := 10, 1

	for _, instruction := range instructions {
		action, value := parseInstruction(instruction)
		if action == "F" {
			currX += wpX*value
			currY += wpY*value
		} else if action == "R" || action == "L" {
			wpX, wpY = rotateWaypoint(wpX, wpY, action, value)
		} else {
			deltaX, deltaY := parseDirection(action, value)
			wpX += deltaX
			wpY += deltaY
		}
	}
	return int(math.Abs(float64(currX))) + int(math.Abs(float64(currY)))
}

func rotateWaypoint(wpX, wpY int, action string, value int) (int, int) {
	value = value % 360

	// Rotating D on left is the same as rotating 360-D on right
	if action == "R" {
		value = 360 - value
	}

	if value == 0 {
		return wpX, wpY
	} else if value == 90 {
		return -wpY, wpX
	} else if value == 180 {
		return -wpX, -wpY
	} else {
		return wpY, -wpX
	}
}

func parseInstruction(instruction string) (string, int) {
	action := string(instruction[0])
	value, _ := strconv.Atoi(instruction[1:])

	return action, value
}

func getCurrentDirectionVector(dir string, amplitude int) (int, int) {
	switch dir {
	case "E":
		return 1, 0
	case "S":
		return 0, -1
	case "W":
		return -1, 0
	case "N":
		return 0, 1
	default:
		panic("fuckkkkk")
	}
}

func parseDirection(action string, value int) (int, int) {
	switch action {
	case "N":
		return 0, value
	case "S":
		return 0, -value
	case "E":
		return value, 0
	case "W":
		return -value, 0
	default:
		panic("fuck!")
	}
}

func getUpdatedDirection(dir, action string, value int) string {
	dirToDegree := map[string]int{
		"E": 0, "N": 90, "W": 180, "S": 270,
	}
	degreeToDir := map[int]string{
		0: "E", 90: "N", 180: "W", 270:"S",
	}
	value = value % 360

	// Rotating D on left is the same as rotating 360-D on right
	if action == "R" {
		value = 360 - value
	}

	value = value % 360
	updatedDegree := value + dirToDegree[dir]
	return degreeToDir[updatedDegree%360]
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
