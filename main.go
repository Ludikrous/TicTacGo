package main

import (
	"errors"
	"fmt"
	"log"
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
	fmt.Println("===================================")
	fmt.Println(" _______    ______         _____     \n/_  __(_)__/_  __/__ _____/ ___/__  \n / / / / __// / / _ `/ __/ (_ / _ \\ \n/_/ /_/\\__//_/  \\_,_/\\__/\\___/\\___/									   \n")
	fmt.Println("===================================")
	fmt.Println("Welcome to TicTacGo! Follow the prompts to play.\n\n")

	keepGoing := true

	for keepGoing {

		fmt.Println("===[ NEW GAME ]===================================")
		firstPlayer, player, computer := chooseOrder()
		fmt.Println(firstPlayer + " will go first as O")
		board := board.New(3)
		fmt.Println(board)

		playerDecider := 0
		if firstPlayer == "Computer" {
			playerDecider = 1
		}

		playerWon, computerWon := false, false

		for !board.BoardFull() {
			// check for a win
			if playerDecider%2 == 0 {
				// make player go
				playerTurn(board, player)
				playerWon, _ = board.CheckForWinner(player)
				if playerWon {
					gameEnded(GEplayerWon)
					keepGoing = askForRestart()
					break
				}
			} else {
				//make computer go
				computerTurn(board, computer)
				computerWon, _ = board.CheckForWinner(computer)
				if computerWon {
					gameEnded(GEcomputerWon)
					keepGoing = askForRestart()
					break
				}
			}
			fmt.Println(board)

			playerDecider++
			turnCount++
		}
		if !(playerWon || computerWon) {
			gameEnded(GEdrawGame)
			keepGoing = askForRestart()
		}
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

func computerTurn(board *board.Board, piece board.Piece) {
	err := errors.New("")

	for err != nil {

		seed := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(seed)
		toss := rng.Intn(9)

		row := toss % 3
		col := toss / 3

		err = board.Set(row, col, piece)
	}
}

func playerTurn(board *board.Board, piece board.Piece) {
	err := errors.New("")

	for err != nil {
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
		err = board.Set(row, col, piece)
		if err != nil {
			fmt.Println("Invalid input. Try again.")
		}
	}
}

func gameEnded(finalResult string) {
	fmt.Println(finalResult)
	fmt.Printf("The game lasted %v turns.", turnCount)
}

func askForRestart() bool {
	fmt.Print("\nDo you want to play again? (y/n) ")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	if response == "y" {
		fmt.Println("\n\n\n")
		return true
	} else if response == "n" {
		fmt.Println("\n\n\n")
		return false
	} else {
		fmt.Println("Please type 'y' or 'n' and then press enter:")
		return askForRestart()
	}
}
