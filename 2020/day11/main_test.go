package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetNumOccupiedSeatsDirectionally(t *testing.T) {
	var seats = [][]string{
		{".", ".", ".", ".", ".", ".", ".", "#", "."},
		{".", ".", ".", "#", ".", ".", ".", ".", "."},
		{".", "#", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "#", "L", ".", ".", ".", ".", "#"},
		{".", ".", ".", ".", "#", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"#", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", "#", ".", ".", ".", ".", "."},
	}
	assert.Equal(t, getNumAdjacentOccupiedSeatsDirectionally(4, 3, seats), 8, "getNumAdjacentOccupiedSeatsDirectionally(4, 3, seats) != 8")

	var seats1 = [][]string {
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "L", ".", "L", ".", "#", ".", "#", ".", "#", ".", "#", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}

	assert.Equal(t, getNumAdjacentOccupiedSeatsDirectionally(1, 1, seats1), 0, "getNumAdjacentOccupiedSeatsDirectionally(1, 1, seats1) != 0")

	var seats2 = [][]string {
		{".", "#", "#", ".", "#", "#", "."},
		{"#", ".", "#", ".", "#", ".", "#"},
		{"#", "#", ".", ".", ".", "#", "#"},
		{".", ".", ".", "L", ".", ".", "."},
		{"#", "#", ".", ".", ".", "#", "#"},
		{"#", ".", "#", ".", "#", ".", "#"},
		{".", "#", "#", ".", "#", "#", "."},
	}

	assert.Equal(t, getNumAdjacentOccupiedSeatsDirectionally(3, 3, seats2), 0, "getNumAdjacentOccupiedSeatsDirectionally(3, 3, seats2) != 0")
}