package main

import (
	"fmt"

	"go.qvidahl.no/simulator/simulator"
)

func main() {
	//iterate()
	playOne()
}

func playOne() {
	//s.SetAggressive(0.8)
	//s.IgnoreTripleTwos(0.5)
	var s simulator.GameData

	s.SetMinToKeep(0)
	s.SetWhenToQuit(3)
	s.PlayRounds()
	fmt.Printf("Antall runder: %v, sum: %v", s.GetRounds(), s.GetTotalPoints())
}

func iterate() {
	//s.SetAggressive(0.8)
	//s.IgnoreTripleTwos(0.5)

	avgRoundsToWin := 0.0
	totalRounds := 0
	counter := 0
	games := 300

	for counter < games {
		var s simulator.GameData

		s.SetMinToKeep(0)
		s.SetWhenToQuit(3)
		s.PlayRounds()
		totalRounds += s.GetRounds()
		counter++
	}
	avgRoundsToWin = float64(totalRounds) / float64(counter)

	fmt.Printf("Gjennomsnittlig antall runder for å nå 10000, kjører %v spill: %.2f", games, avgRoundsToWin)
	//fmt.Printf("Antall runder: %v, sum: %v", s.GetRounds(), s.GetTotalPoints())
}
