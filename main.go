package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ludikrous/TicTacGo/board"
)

//===============================================

var turnCount = 0

func main() {
	fmt.Println("Welcome to TicTacGo! Follow the prompts to play.")
	firstPlayer, player, computer := chooseOrder()
	fmt.Println(firstPlayer + " will go first as O")
	board := board.New(3)
	for board.BoardFull() {
		// todo
	}
}

//===============================================

func chooseOrder() (firstPlayer string, player, computer board.Piece) {
	seed := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(seed)
	toss := rng.Intn(2)

	switch {
	case toss == 0:
		return "Computer", board.X, board.O
	case toss == 1:
		return "Player", board.O, board.X
	}
	return "", board.Blank, board.Blank // this should never happen
}

func computerTurn(board *board.Board, piece int) {

}

func playerTurn(board *board.Board, piece board.Piece) {
	// get row from user
	fmt.Print("Row you wish to place a piece in: ")
	var row int
	if _, err := fmt.Scanf("%d", &row); err != nil {
		fmt.Println("Invalid input! Try again.")
		playerTurn(board, piece)
	}
	// get col from user
	fmt.Print("Column you wish to place a piece in: ")
	var col int
	if _, err := fmt.Scanf("%d", &col); err != nil {
		fmt.Println("Invalid input! Try again.")
		playerTurn(board, piece)
	}

	// place user's piece in the board
	board.Set(row, col, piece)
}
