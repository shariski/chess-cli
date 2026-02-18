package main

import "testing"

func TestPawn_IsValidMove(t *testing.T) {

	t.Run("white pawn move forward one", func(t *testing.T) {
		board := &Board{}
		pawn := &Pawn{color: "white"}

		start := Position{row: 6, col: 0}
		end := Position{row: 5, col: 0}

		board.grid[6][0] = pawn

		if !pawn.IsValidMove(board, start, end) {
			t.Fatalf("expected valid move")
		}
	})

	t.Run("white pawn move forward two first move", func(t *testing.T) {
		board := &Board{}
		pawn := &Pawn{color: "white"}

		start := Position{row: 6, col: 0}
		end := Position{row: 4, col: 0}

		board.grid[6][0] = pawn

		if !pawn.IsValidMove(board, start, end) {
			t.Fatalf("expected valid double move")
		}
	})

	t.Run("white pawn cannot move two after moved", func(t *testing.T) {
		board := &Board{}
		pawn := &Pawn{color: "white", hasMoved: true}

		start := Position{row: 6, col: 0}
		end := Position{row: 4, col: 0}

		board.grid[6][0] = pawn

		if pawn.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid double move after moved")
		}
	})

	t.Run("pawn blocked forward", func(t *testing.T) {
		board := &Board{}
		pawn := &Pawn{color: "white"}

		start := Position{row: 6, col: 0}
		end := Position{row: 5, col: 0}

		board.grid[6][0] = pawn
		board.grid[5][0] = &Pawn{color: "black"} // blocking

		if pawn.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move due to blocking piece")
		}
	})

	t.Run("pawn capture diagonal", func(t *testing.T) {
		board := &Board{}
		pawn := &Pawn{color: "white"}

		start := Position{row: 6, col: 0}
		end := Position{row: 5, col: 1}

		board.grid[6][0] = pawn
		board.grid[5][1] = &Pawn{color: "black"}

		if !pawn.IsValidMove(board, start, end) {
			t.Fatalf("expected valid capture")
		}
	})

	t.Run("pawn cannot capture empty diagonal", func(t *testing.T) {
		board := &Board{}
		pawn := &Pawn{color: "white"}

		start := Position{row: 6, col: 0}
		end := Position{row: 5, col: 1}

		board.grid[6][0] = pawn

		if pawn.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid diagonal move")
		}
	})

	t.Run("pawn cannot move backward", func(t *testing.T) {
		board := &Board{}
		pawn := &Pawn{color: "white"}

		start := Position{row: 6, col: 0}
		end := Position{row: 7, col: 0}

		board.grid[6][0] = pawn

		if pawn.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid backward move")
		}
	})

	t.Run("pawn cannot capture teammate", func(t *testing.T) {
		board := &Board{}
		pawn := &Pawn{color: "white"}

		start := Position{row: 6, col: 0}
		end := Position{row: 5, col: 1}

		board.grid[6][0] = pawn
		board.grid[5][1] = &Pawn{color: "white"} // teammate

		if pawn.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move when capturing teammate")
		}
	})
}
