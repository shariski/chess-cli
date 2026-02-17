package main

import (
	"errors"
	"fmt"
	"strings"
)

type Piece interface {
	Symbol() string
	Color() string
	Move()
	IsValidMove(board *Board, start, end Position) bool
}

type Board struct {
	grid   [8][8]Piece
	rowIdx string
	colIdx string
}

type Position struct {
	row int
	col int
}

func NewBoard() *Board {
	board := new(Board)
	board.rowIdx = "87654321"
	board.colIdx = "abcdefgh"

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
	if !b.isValidPosition(start) {
		return errors.New("start position is not valid")
	}

	if !b.isValidPosition(end) {
		return errors.New("end position is not valid")
	}

	piece := b.Get(start)
	if piece == nil {
		return errors.New("square is empty")
	}

	// validate piece movement pattern
	if !piece.IsValidMove(b, start, end) {
		return errors.New("piece movement invalid")
	}

	// can't do friendly fire
	target := b.Get(end)
	if target != nil {
		if target.Color() == piece.Color() {
			return errors.New("cannot capture teammate")
		}
	}

	// TODO: is path clear?

	// update hasMoved to true
	piece.Move()

	b.Set(end, piece)
	b.Set(start, nil)

	return nil
}

func (b *Board) Render() {
	var sb strings.Builder

	for i, inner := range b.grid {
		fmt.Fprintf(&sb, "%c | ", b.rowIdx[i])

		for j, piece := range inner {
			if piece != nil {
				fmt.Fprintf(&sb, "%s  ", piece.Symbol())
			} else {
				if (i+j)%2 == 0 {
					sb.WriteString("-  ")
				} else {
					sb.WriteString("+  ")
				}
			}
		}

		sb.WriteString("\n")
	}

	sb.WriteString("  ")

	for range b.grid[0] {
		sb.WriteString("---")
	}

	sb.WriteString("\n")
	sb.WriteString("    ")

	for j := range b.grid[0] {
		fmt.Fprintf(&sb, "%c  ", b.colIdx[j])
	}

	sb.WriteString("\n\n")

	fmt.Print(sb.String())
}

// first letter is not capitalized (for internal user only)
func (b *Board) isValidPosition(pos Position) bool {
	if pos.row < 0 || pos.row > 7 {
		return false
	}

	if pos.col < 0 || pos.col > 7 {
		return false
	}

	return true
}

func (b *Board) Get(pos Position) Piece {
	return b.grid[pos.row][pos.col]
}

func (b *Board) Set(pos Position, p Piece) error {
	b.grid[pos.row][pos.col] = p

	return nil
}

func (b *Board) IsEmpty(pos Position) bool {
	piece := b.Get(pos)

	return piece == nil
}

func (b *Board) HasEnemyPiece(pos Position, color string) bool {
	piece := b.Get(pos)

	if piece == nil {
		return false
	}

	return piece.Color() != color
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
