package main

import (
	"fmt"

	"go.qvidahl.no/simulator/simulator"
)

func main() {

	var s simulator.GameData

	s.SetWhenToQuit(2)
	//s.SetMinToKeep(300)
	//s.SetAggressive(1)
	s.PlayRounds()

	fmt.Printf("Antall runder: %v, sum: %v", s.GetRounds(), s.GetTotalPoints())
}
