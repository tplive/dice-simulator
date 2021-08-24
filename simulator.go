package main

import (
	"fmt"
	"math/rand"
	"time"
)

type dice []int

func main() {

	var roll dice
	var points int
	roll = dice{rollDie(), rollDie(), rollDie(), rollDie(), rollDie(), rollDie()}

	fmt.Println(roll)
	for i := 1; i < 7; i++ {
		var occ = countOccurrence(i, roll)
		var pts = pointsForDice(i, occ)
		if pts > 0 {
			fmt.Println("Antall ", i, ": ", occ, " gir ", pts, " poeng")
			points += pts
		}
	}
	fmt.Println("Totalt: ", points)
	//fmt.Println("Antall enere: ", countOccurrence(1, roll), " gir ", pointsForDice(1, countOccurrence(1, roll)), " poeng")
	// fmt.Println("Antall toere: ", countOccurrence(2, roll), " gir ", pointsForDice(2, countOccurrence(2, roll)), " poeng")
	// fmt.Println("Antall treere: ", countOccurrence(3, roll))
	// fmt.Println("Antall firere: ", countOccurrence(4, roll))
	// fmt.Println("Antall femmere: ", countOccurrence(5, roll))
	// fmt.Println("Antall seksere: ", countOccurrence(6, roll))
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
