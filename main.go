package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ludikrous/TicTacGo/board"
)

//===============================================
const o = 1
const x = 2

var turnCount = 0

func main() {
	fmt.Println("Welcome to TicTacGo! Follow the prompts to play.")
	firstPlayer := chooseOrder()
	fmt.Println(firstPlayer + " will go first as O")
	board := board.New(3)
	for board.BoardFull() {
		// todo
	}
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

func computerTurn(board *board.Board) {

}

func playerTurn(board *board.Board, piece int) {
	// get row from user
	fmt.Print("Row you wish to place a piece in: ")
	var row int
	if _, err := fmt.Scanf("%d", &row); err != nil {
		return
	}
	// get col from user
	fmt.Print("Column you wish to place a piece in: ")
	var col int
	if _, err := fmt.Scanf("%d", &col); err != nil {
		return
	}

}
