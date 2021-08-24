package main

import (
	"fmt"
	"math/rand"
	"time"
)

type dice []int

func main() {

	roll := dice{rollDie(), rollDie(), rollDie(), rollDie(), rollDie(), rollDie()}

	fmt.Println(roll)

	fmt.Println("Totalt: ", getTotal(roll))
}

func isStraight(dice dice) bool {

	return (countOccurrence(1, dice) == 1 &&
		countOccurrence(2, dice) == 1 &&
		countOccurrence(3, dice) == 1 &&
		countOccurrence(4, dice) == 1 &&
		countOccurrence(5, dice) == 1 &&
		countOccurrence(6, dice) == 1)
}

func isThreePairs(roll dice) bool {
	unique := dice{}
	count := 0
	count, unique = countDistinct(roll)
	return len(roll) == 6 && count == 3 && len(unique) == 3

}

func containsDie(roll dice, faceValue int) bool {

	for _, v := range roll {
		if v == faceValue {
			return true
		}
	}
	return false

}

func countDistinct(roll dice) (int, dice) {
	unique := dice{}
	skip := false
	for _, v := range roll {
		for _, u := range unique {
			if v == u {
				skip = true
			}
		}
		if !skip {
			unique = append(unique, v)
		} else {
			skip = false
		}
	}
	return len(unique), unique
}

func getTotal(dice dice) int {
	var points int
	if isStraight(dice) {
		points = 1500
		fmt.Printf("Straight! %v\n", dice)
	} else if isThreePairs(dice) {
		points = 1000
		fmt.Printf("Tre par! %v\n", dice)
	} else {

		for i := 1; i < 7; i++ {
			var occ = countOccurrence(i, dice)
			var pts = pointsForDice(i, occ)
			if pts > 0 {
				points += pts
			}
		}
	}
	return points
}

func pointsForDice(faceValue int, number int) int {
	var points = 0
	if faceValue == 1 {
		if number == 1 {
			points += 100
		} else if number == 2 {
			points += 200
		} else if number == 3 {
			points += 1000
		} else if number == 4 {
			points += 2000
		} else if number == 5 {
			points += 4000
		} else if number == 6 {
			points += 8000
		}
	} else if faceValue == 2 {
		if number == 3 {
			points += 200
		} else if number == 4 {
			points += 400
		} else if number == 5 {
			points += 800
		} else if number == 6 {
			points += 1600
		}
	} else if faceValue == 3 {
		if number == 3 {
			points += 300
		} else if number == 4 {
			points += 600
		} else if number == 5 {
			points += 1200
		} else if number == 6 {
			points += 2400
		}
	} else if faceValue == 4 {
		if number == 3 {
			points += 400
		} else if number == 4 {
			points += 800
		} else if number == 5 {
			points += 1600
		} else if number == 6 {
			points += 3200
		}
	} else if faceValue == 5 {
		if number == 1 {
			points += 50
		} else if number == 2 {
			points += 100
		} else if number == 3 {
			points += 500
		} else if number == 4 {
			points += 1000
		} else if number == 5 {
			points += 2000
		} else if number == 6 {
			points += 4000
		}
	} else if faceValue == 6 {
		if number == 3 {
			points += 600
		} else if number == 4 {
			points += 1200
		} else if number == 5 {
			points += 2400
		} else if number == 6 {
			points += 4800
		}
	}
	return points
}

func countOccurrence(value int, dice dice) int {
	count := 0
	for _, num := range dice {
		if num == value {
			count++
		}
	}
	return count
}

func rollDie() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}
