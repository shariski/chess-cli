package main

import (
	"errors"
	"fmt"
)

type Piece interface {
	Symbol() string
	Color() string
}

type Board struct {
	grid [8][8]Piece
}

type Position struct {
	row int
	col int
}

func NewBoard() *Board {
	board := new(Board)

	// black pieces
	// non-pawn pieces
	board.grid[0][0] = &Rook{color: "black"}
	board.grid[0][1] = &Knight{color: "black"}
	board.grid[0][2] = &Bishop{color: "black"}
	board.grid[0][3] = &Queen{color: "black"}
	board.grid[0][4] = &King{color: "black"}
	board.grid[0][5] = &Bishop{color: "black"}
	board.grid[0][6] = &Knight{color: "black"}
	board.grid[0][7] = &Rook{color: "black"}

	// pawn pieces
	board.grid[1][0] = &Pawn{color: "black"}
	board.grid[1][1] = &Pawn{color: "black"}
	board.grid[1][2] = &Pawn{color: "black"}
	board.grid[1][3] = &Pawn{color: "black"}
	board.grid[1][4] = &Pawn{color: "black"}
	board.grid[1][5] = &Pawn{color: "black"}
	board.grid[1][6] = &Pawn{color: "black"}
	board.grid[1][7] = &Pawn{color: "black"}

	// white pieces
	// pawn pieces
	board.grid[6][0] = &Pawn{color: "white"}
	board.grid[6][1] = &Pawn{color: "white"}
	board.grid[6][2] = &Pawn{color: "white"}
	board.grid[6][3] = &Pawn{color: "white"}
	board.grid[6][4] = &Pawn{color: "white"}
	board.grid[6][5] = &Pawn{color: "white"}
	board.grid[6][6] = &Pawn{color: "white"}
	board.grid[6][7] = &Pawn{color: "white"}

	// non-pawn pieces
	board.grid[7][0] = &Rook{color: "white"}
	board.grid[7][1] = &Knight{color: "white"}
	board.grid[7][2] = &Bishop{color: "white"}
	board.grid[7][3] = &Queen{color: "white"}
	board.grid[7][4] = &King{color: "white"}
	board.grid[7][5] = &Bishop{color: "white"}
	board.grid[7][6] = &Knight{color: "white"}
	board.grid[7][7] = &Rook{color: "white"}

	return board
}

func (b *Board) Move(start, end Position) error {
	if !b.isValidPostion(start) {
		return errors.New("start position is not valid")
	}

	if !b.isValidPostion(end) {
		return errors.New("end position is not valid")
	}

	piece := b.grid[start.row][start.col]
	if piece == nil {
		return errors.New("square is empty")
	}

	// TODO: validate piece movement pattern

	captured := b.grid[end.row][end.col]
	if captured != nil {
		if captured.Color() == piece.Color() {
			return errors.New("cannot capture teammate")
		}
	}

	// TODO: hasMoved

	b.grid[start.row][start.col] = nil
	b.grid[end.row][end.col] = piece

	return nil
}

func (b *Board) Print() {
	for _, inner := range b.grid {
		for _, value := range inner {
			if value != nil {
				fmt.Printf("%s ", value.Symbol())
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// first letter is not capitalized (for internal user only)
func (b *Board) isValidPostion(pos Position) bool {
	if pos.row < 0 || pos.row > 7 {
		return false
	}

	if pos.col < 0 || pos.col > 7 {
		return false
	}

	return true
}
