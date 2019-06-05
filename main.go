package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ludikrous/TicTacGo/board"
)

//===============================================

const (
	GEplayerWon   string = "The player won!"
	GEcomputerWon string = "The computer won!"
	GEdrawGame    string = "The game ended in a draw."
)

var turnCount = 0

func main() {
	fmt.Println("Welcome to TicTacGo! Follow the prompts to play.")
	firstPlayer, player, computer := chooseOrder()
	fmt.Println(firstPlayer + " will go first as O")
	board := board.New(3)

	playerDecider := 0
	if firstPlayer == "Computer" {
		playerDecider = 1
	}

	for board.BoardFull() {
		// check for a win
		if playerDecider%2 == 0 {
			// make player go
			playerTurn(board, player)
			playerWon, _ := board.CheckForWinner(player)
			if playerWon {
				gameEnded(GEplayerWon)
			}
		} else {
			//make computer go
			computerTurn(board, computer)
			computerWon, _ := board.CheckForWinner(computer)
			if computerWon {
				gameEnded(GEcomputerWon)
			}
		}
		playerDecider++
		turnCount++
	}
	gameEnded(GEdrawGame)
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

func computerTurn(board *board.Board, piece board.Piece) {

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

func gameEnded(finalResult string) {
	fmt.Println(finalResult)
	fmt.Println("The game lasted " + string(turnCount) + " turns.")
}
