package main

import "testing"

func TestIsStraight(t *testing.T) {
	roll := dice{1, 2, 3, 4, 5, 6}
	roll2 := dice{1, 1, 2, 2, 3, 3}
	roll3 := dice{6, 5, 4, 3, 2, 1}
	roll4 := dice{6, 1, 3, 4, 5, 2}

	if !isStraight(roll) {
		t.Errorf("Input is a Straight so test should pass: %v", roll)
	}

	if isStraight(roll2) {
		t.Errorf("Input is not a Straight so test should pass: %v", roll2)
	}

	if !isStraight(roll3) {
		t.Errorf("Input is a Straight (in reverse) so test should pass: %v", roll3)
	}

	if !isStraight(roll4) {
		t.Errorf("Input is a Straight (unordered) so test should pass: %v", roll4)
	}
}

func TestIsYatzy(t *testing.T) {
	roll := dice{1, 1, 1, 1, 1, 1}
	roll2 := dice{4, 4, 4, 4, 4, 4}
	roll3 := dice{6, 5, 4, 3, 2, 1}

	if !isYatzy(roll) {
		t.Errorf("Input is Yatzy: %v", roll)
	}

	if !isYatzy(roll2) {
		t.Errorf("Input is Yatzy: %v", roll2)
	}

	if isYatzy(roll3) {
		t.Errorf("Input is not a Yatzy: %v", roll3)
	}
}
func TestCountDistinct(t *testing.T) {
	dist := 0
	roll2 := dice{1, 1, 2, 2, 3, 3}
	roll3 := dice{6, 5, 4, 3, 2, 1}
	roll4 := dice{6, 5, 4, 2, 2, 2}

	dist, _ = countDistinct(roll2)
	if dist != 3 {
		t.Errorf("Input is three pairs, test should pass. Distinct was %v", dist)
	}
	dist, _ = countDistinct(roll3)
	if dist != 6 {
		t.Errorf("Input is all unique values, test should pass. Distinct was %v", dist)
	}
	dist, _ = countDistinct(roll4)
	if dist != 4 {
		t.Errorf("Input has 4 unique values, test should pass. Distinct was %v", dist)
	}

}

func TestReturnNonZero(t *testing.T) {
	roll := dice{1, 1, 2, 2, 3, 3}
	roll2 := dice{0, 2, 0, 3, 0, 1}
	roll3 := dice{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if len(returnNonZero(roll)) != 6 {
		t.Errorf("No values are zero: %v", roll)
	}

	if len(returnNonZero(roll2)) != 3 {
		t.Errorf("Three values are zero: %v", roll)
	}

	if len(returnNonZero(roll3)) != 0 {
		t.Errorf("All values are zero: %v", roll)
	}
}

func TestCountNonZero(t *testing.T) {
	roll := dice{1, 1, 2, 2, 3, 3}
	roll2 := dice{0, 2, 0, 3, 0, 1}
	roll3 := dice{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if countNonZero(roll) != 6 {
		t.Errorf("No values are zero: %v", roll)
	}

	if countNonZero(roll2) != 3 {
		t.Errorf("Three values are zero: %v", roll)
	}

	if countNonZero(roll3) != 0 {
		t.Errorf("All values are zero: %v", roll)
	}
}
func TestIsThreePairs(t *testing.T) {
	roll := dice{1, 2, 3, 4, 5, 6}
	roll2 := dice{1, 1, 2, 2, 3, 3}
	roll3 := dice{4, 6, 5, 6, 5, 4}
	roll4 := dice{4, 4, 4, 4, 5, 3}

	if isThreePairs(roll) {
		t.Errorf("Input is not three pairs: %v", roll)
	}

	if !isThreePairs(roll2) {
		t.Errorf("Input is three pairs: %v", roll2)
	}

	if !isThreePairs(roll3) {
		t.Errorf("Input is three pairs: %v", roll3)
	}

	if isThreePairs(roll4) {
		t.Errorf("Input is not three different pairs: %v", roll4)
	}
}

func TestGetTotal(t *testing.T) {
	sum := 0

	roll := dice{1, 2, 3, 4, 5, 6}  // Straight - 1500 points
	roll2 := dice{1, 1, 2, 2, 3, 3} // Three pairs - 1000 points
	roll3 := dice{5, 4, 5, 1, 2, 3} // 200 points
	roll4 := dice{1, 1, 1, 2, 3, 5} // 1050 points
	roll5 := dice{5, 2, 4, 5, 2, 2} // 300 points

	sum = getPoints(roll)
	if sum != 1500 {
		t.Errorf("Sum should be 1500, was: %v", sum)
	}
	sum = getPoints(roll2)
	if sum != 1000 {
		t.Errorf("Sum should be 1000, was: %v", sum)
	}
	sum = getPoints(roll3)
	if sum != 200 {
		t.Errorf("Sum should be 200, was: %v", sum)
	}
	sum = getPoints(roll4)
	if sum != 1050 {
		t.Errorf("Sum should be 1050, was: %v", sum)
	}
	sum = getPoints(roll5)
	if sum != 300 {
		t.Errorf("Sum should be 300, was: %v", sum)
	}

}
