package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(rollDie(), rollDie(), rollDie(), rollDie(), rollDie(), rollDie())

}

func rollDie() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}
