package board

import (
	"bytes"
	"errors"
	"fmt"
)

//===============================================

const blank = 0
const o = 1
const x = 2

// NoWinner indicates that no one is winning
const NoWinner = 0

// OWinner indicates that O is winning
const OWinner = 1

// XWinner indicates that X is winning
const XWinner = 2

// Draw indicates that the came has ended in a draw
const Draw = 3

// A Board represents a Tic-Tac-Toe board.
type Board struct {
	board [][]Piece
}

//===============================================

// New returns a tic tac toe board of specified size
func New(size int) *Board {
	toReturn := Board{}
	toReturn.board = make([][]Piece, size)
	for i := 0; i < size; i++ {
		toReturn.board[i] = make([]Piece, size)
	}
	return &toReturn
}

// SetX sets the value at row, col on the board to x
func (board *Board) SetX(row, col int, p Piece) error {
	if board.board[row][col] != 0 {
		return errors.New("field is not blank")
	}

	board.board[row][col] = p
	return nil
}

// CheckForWinner parses the board to look
// for a winner
func (board Board) CheckForWinner(player Piece) (bool, error) {
	// check to make sure player is valid
	if player != x && player != o {
		return false, errors.New("invalid input")
	}

	// check row
	for _, row := range board.board {
		for loc, val := range row {
			if val != player {
				break
			}
			if loc == len(board.board)-1 {
				return true, nil
			}
		}
	}

	// check column
	for col := 0; col < len(board.board); col++ {
		for row := 0; row < len(board.board); row++ {
			if board.board[row][col] != player {
				break
			}
			if row == len(board.board)-1 {
				return true, nil
			}
		}
	}

	// check diagonal
	for pos := 0; pos < len(board.board); pos++ {
		if board.board[pos][pos] != player {
			break
		}
		if pos == len(board.board)-1 {
			return true, nil
		}
	}

	// check anti-diagonal
	for pos := 0; pos < len(board.board); pos++ {
		if board.board[len(board.board)-pos-1][pos] != player {
			break
		}
		if pos == len(board.board)-1 {
			return true, nil
		}
	}

	return false, nil
}

func (board Board) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("\n")
	for _, row := range board.board {
		for _, val := range row {
			buffer.WriteString(toPiece(val) + " ")
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

func (board Board) Print() {
	fmt.Println(board.board)
}

//===============================================

func (board Board) BoardFull() bool {
	// iterate through the board to look to an empty spot
	for _, row := range board.board {
		for _, value := range row {
			if value == blank {
				return false
			}
		}
	}
	// no empty spots were found
	return true
}

func toPiece(p Piece) string {
	switch {
	case p == blank:
		return "-"
	case p == o:
		return "O"
	case p == x:
		return "X"
	}
	return ""
}
