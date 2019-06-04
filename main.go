package main

import (
	"fmt"

	"github.com/ludikrous/TicTacGo/board"
)

//===============================================

func main() {
	board := board.New(3)
	fmt.Println(board)
	board.SetX(0, 0)
	board.SetX(1, 1)
	board.SetX(2, 2)
	fmt.Println(board)
	fmt.Println(board.CheckForWinner(0))
}
