package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetUpdatedDirection(t *testing.T) {
	assert.Equal(t, getUpdatedDirection("E", "R", 90), "S", "getUpdatedDirection(E, R90) != S")
	assert.Equal(t, getUpdatedDirection("E", "R", 180), "W", "getUpdatedDirection(E, R180) != W")
	assert.Equal(t, getUpdatedDirection("E", "R", 270), "N", "getUpdatedDirection(E, R270) != N")
	assert.Equal(t, getUpdatedDirection("E", "R", 360), "E", "getUpdatedDirection(E, R360) != E")
	assert.Equal(t, getUpdatedDirection("E", "L", 90), "N", "getUpdatedDirection(E, L90) != N")
	assert.Equal(t, getUpdatedDirection("E", "L", 180), "W", "getUpdatedDirection(E, L180) != W")
	assert.Equal(t, getUpdatedDirection("E", "L", 270), "S", "getUpdatedDirection(E, L270) != N")
	assert.Equal(t, getUpdatedDirection("E", "L", 360), "E", "getUpdatedDirection(E, L360) != E")
}