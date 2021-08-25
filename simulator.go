package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type dice []int

func main() {

	fmt.Println("Antall runder: ", playRound())
}

func playRound() int {
	sum := 0
	round := 1
	for sum < 10000 {

		points, roll := rollOnce(rollDice(6))
		sum += points
		fmt.Printf("%v = %v pts  -> %v\n", sortDice(roll), points, sum)

		round++
	}
	return round
}

func rollOnce(roll dice) (int, dice) {
	poeng := getPoints(roll)

	return poeng, roll
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
	count, unique := countDistinct(roll)
	roll = sortDice(roll)
	return len(roll) == 6 && count == 3 && len(unique) == 3 && roll[0] == roll[1] && roll[1] != roll[2] && roll[2] == roll[3] && roll[3] != roll[4] && roll[4] == roll[5]
}

func isYatzy(roll dice) bool {
	roll = sortDice(roll)
	return roll[0] == roll[len(roll)-1]
}

func sortDice(roll dice) dice {
	sort.Slice(roll, func(i, j int) bool { return roll[i] < roll[j] })
	return roll
}

func countNonZero(roll dice) int {
	count := 0
	for _, r := range roll {
		if r != 0 {
			count++
		}
	}
	return count
}

func returnNonZero(roll dice) dice {
	nonZero := dice{}
	for _, r := range roll {
		if r != 0 {
			nonZero = append(nonZero, r)
		}
	}
	return nonZero
}

func countDistinct(roll dice) (int, dice) {
	// Return the number of distinct values as well as a dice array of those values
	unique := dice{}
	exists := false
	for _, v := range roll {
		for _, u := range unique {
			if v == u {
				exists = true
			}
		}
		if !exists {
			unique = append(unique, v)
		} else {
			exists = false
		}
	}
	return len(unique), unique
}

func getPoints(roll dice) int {
	points := 0
	if isYatzy(roll) {
		points = 10000
	} else if isStraight(roll) {
		points = 1500
	} else if isThreePairs(roll) {
		points = 1000
	} else {
		for i := 1; i <= len(roll); i++ {
			var occ = countOccurrence(i, roll)
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

func rollDice(numberToRoll int) dice {
	roll := dice{}
	for i := 0; i <= numberToRoll; i++ {
		roll = append(roll, rollDie())
	}
	return roll
}
