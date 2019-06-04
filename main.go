package main

import (
	"fmt"
	"math/rand"
	"time"
)

//===============================================
const o = 1
const x = 2

var turnCount = 0

func main() {
	fmt.Println("Welcome to TicTacGo! Follow the prompts to play.")
	firstPlayer := chooseOrder()
	fmt.Println(firstPlayer + " will go first as O")

}

//===============================================

func chooseOrder() (firstPlayer string) {
	seed := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(seed)
	toss := rng.Intn(2)

	switch {
	case toss == 0:
		return "Computer"
	case toss == 1:
		return "Player"
	}
	return "" // this should never happen
}
