package main

import "testing"

// Test if isStraight works
func TestIsStraight(t *testing.T) {
	roll := dice{1, 2, 3, 4, 5, 6}

	if !isStraight(roll) {
		t.Errorf("Input is a Straight so test should return true: %v", roll)
	}

	roll2 := dice{1, 1, 2, 2, 3, 3}

	if isStraight(roll2) {
		t.Errorf("Input is not a Straight so test should return true: %v", roll2)
	}

	roll3 := dice{6, 5, 4, 3, 2, 1}
	if !isStraight(roll3) {
		t.Errorf("Input is a Straight so test should return true: %v", roll3)
	}
}
