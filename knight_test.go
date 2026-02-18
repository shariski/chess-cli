package main

import "testing"

func TestKnight_IsValidMove(t *testing.T) {
	t.Run("white knight move forward from base", func(t *testing.T) {
		board := &Board{}
		knight := &Knight{color: "white"}

		start := Position{row: 7, col: 6}
		end := Position{row: 5, col: 5}

		board.grid[7][6] = knight

		if !knight.IsValidMove(board, start, end) {
			t.Fatalf("expected valid move")
		}
	})

	t.Run("white knight move forward with piece in front of him", func(t *testing.T) {
		board := &Board{}
		knight := &Knight{color: "white"}
		pawn := &Pawn{color: "white"}

		start := Position{row: 7, col: 6}
		end := Position{row: 5, col: 5}

		board.grid[7][6] = knight
		board.grid[6][6] = pawn // white pawn in front of him

		if !knight.IsValidMove(board, start, end) {
			t.Fatalf("expected valid move")
		}

	})

	t.Run("black knight capture enemy's king", func(t *testing.T) {
		board := &Board{}
		knight := &Knight{color: "black"}
		king := &King{color: "white"}

		board.grid[5][5] = knight
		board.grid[7][4] = king

		start := Position{row: 5, col: 5}
		end := Position{row: 7, col: 4}

		// there's no winning condition here
		// winning condition is determined in board and game layer
		if !knight.IsValidMove(board, start, end) {
			t.Fatalf("expected valid move and capture")
		}
	})

	t.Run("black knight do straight slide move", func(t *testing.T) {
		board := &Board{}
		knight := &Knight{color: "white"}

		board.grid[0][6] = knight

		start := Position{row: 0, col: 6}
		end := Position{row: 7, col: 6}

		if knight.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move, cannot do slide straight")
		}
	})
}
