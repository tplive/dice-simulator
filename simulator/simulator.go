package simulator

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type dice []int

type GameData struct {
	round      int
	total      int
	whenToQuit int
	minToKeep  int
}

func (s *GameData) SetMinToKeep(k int) {
	s.minToKeep = k
}

func (s *GameData) GetRounds() int {
	return s.round
}

func (s *GameData) GetTotalPoints() int {
	return s.total
}

func (s *GameData) SetWhenToQuit(q int) {
	s.whenToQuit = q
}

func (s *GameData) AddPointsToTotal(p int) {
	s.total += p
}

func (s *GameData) ResetSimulator() {
	s.round = 1
	s.total = 0

}

func (s *GameData) PlayRounds() {

	s.round = 1
	whenToQuit := s.whenToQuit
	for s.total <= 10000 {
		if s.round >= 50 {
			fmt.Printf("No more than 50 rounds plz..\n")
			break
		}

		fmt.Printf("R%v ", s.round)
		s.playRound(whenToQuit)

		s.round++
	}
}

func (s *GameData) playRound(whenToQuit int) int {
	var rollPoints, accPoints, remaining int
	endRound := false
	var roll = dice{}
	remaining = 6

	for {
		rollPoints, roll, remaining = rollOnce(rollDice(remaining))

		// If we are already in the game, know when to quit
		if s.total >= 1000 {
			whenToQuit = s.whenToQuit
		}

		if rollPoints == 0 {
			// If we get no points, wipe the accumulated points and end the round
			accPoints = 0
			endRound = true
		} else if s.total == 0 && accPoints+rollPoints >= 1000 {
			// If we just got in this round, play it safe and finish up
			accPoints += rollPoints
			endRound = true
		} else if s.total >= 1000 && accPoints >= s.minToKeep && remaining <= whenToQuit {
			// if we are already in, have the minimum we want to keep, and we have fewer die remaining than we need to go on, then end the round
			accPoints += rollPoints
			endRound = true
		} else if rollPoints > 0 && remaining == 0 {
			// but if we got some points, and all dice yield points, we can keep the points and roll again!
			remaining = 6
			accPoints += rollPoints
		} else if rollPoints > 0 {
			accPoints += rollPoints
		}

		if endRound {
			// When we end the round, we add points (if any) to the total
			s.AddPointsToTotal(accPoints)
			fmt.Printf(" %v gir %vpts = %vpts", roll, rollPoints, accPoints)
			fmt.Printf(" => %vpts\n", s.total)
			break
		} else {
			// if we got some points, but aren't ready to quit yet, accumulate and keep going
			fmt.Printf(" %v gir %vpts = %vpts", roll, rollPoints, accPoints)
		}
	}

	return accPoints
}

func rollOnce(roll dice) (int, dice, int) {
	poeng, remaining := getPoints(roll)

	return poeng, roll, remaining
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
	equal := false
	if len(roll) < 6 {
		equal = false
	} else {
		roll = sortDice(roll)
		equal = true
		for i := range roll {
			if roll[0] != roll[i] {
				equal = false
				break
			}
		}
	}
	return equal
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

func getPoints(roll dice) (int, int) {
	points := 0
	remaining := len(roll)
	if isYatzy(roll) {
		points = 10000
		remaining = 0
	} else if isStraight(roll) {
		points = 1500
		remaining = 0
	} else if isThreePairs(roll) {
		points = 1000
		remaining = 0
	} else {
		for i := 1; i < 7; i++ {
			var occ = countOccurrence(i, roll)
			var pts = pointsForDice(i, occ)
			if pts > 0 {
				points += pts
				remaining -= occ
			}
		}
	}
	return points, remaining
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
		}
	} else if faceValue == 2 {
		if number == 3 {
			points += 200
		} else if number == 4 {
			points += 400
		} else if number == 5 {
			points += 800
		}
	} else if faceValue == 3 {
		if number == 3 {
			points += 300
		} else if number == 4 {
			points += 600
		} else if number == 5 {
			points += 1200
		}
	} else if faceValue == 4 {
		if number == 3 {
			points += 400
		} else if number == 4 {
			points += 800
		} else if number == 5 {
			points += 1600
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
		}
	} else if faceValue == 6 {
		if number == 3 {
			points += 600
		} else if number == 4 {
			points += 1200
		} else if number == 5 {
			points += 2400
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
	for i := 0; i < numberToRoll; i++ {
		roll = append(roll, rollDie())
	}
	return sortDice(roll)
}
