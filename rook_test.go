package main

import "testing"

func TestRook_IsValidMove(t *testing.T) {
	t.Run("black rook move forward", func(t *testing.T) {
		board := &Board{}
		rook := &Rook{color: "black"}

		board.grid[0][6] = rook

		start := Position{row: 0, col: 6}
		end := Position{row: 7, col: 6}

		if !rook.IsValidMove(board, start, end) {
			t.Fatalf("expected valid move")
		}
	})

	t.Run("rook move vertical", func(t *testing.T) {
		board := &Board{}
		rook := &Rook{color: "black"}

		board.grid[0][6] = rook

		start := Position{row: 0, col: 6}
		end := Position{row: 7, col: 6}

		if !rook.IsValidMove(board, start, end) {
			t.Fatalf("expected valid vertical move")
		}
	})

	t.Run("rook move horizontal", func(t *testing.T) {
		board := &Board{}
		rook := &Rook{color: "black"}

		board.grid[4][4] = rook

		start := Position{row: 4, col: 4}
		end := Position{row: 4, col: 0}

		if !rook.IsValidMove(board, start, end) {
			t.Fatalf("expected valid horizontal move")
		}
	})

	t.Run("rook cannot move diagonal", func(t *testing.T) {
		board := &Board{}
		rook := &Rook{color: "black"}

		board.grid[4][4] = rook

		start := Position{row: 4, col: 4}
		end := Position{row: 6, col: 6}

		if rook.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid diagonal move")
		}
	})

	t.Run("rook path blocked vertically", func(t *testing.T) {
		board := &Board{}
		rook := &Rook{color: "black"}

		board.grid[0][6] = rook
		board.grid[3][6] = &Pawn{color: "black"} // blocking

		start := Position{row: 0, col: 6}
		end := Position{row: 7, col: 6}

		if rook.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move due to blocking piece")
		}
	})

	t.Run("rook path blocked horizontally", func(t *testing.T) {
		board := &Board{}
		rook := &Rook{color: "black"}

		board.grid[4][4] = rook
		board.grid[4][2] = &Pawn{color: "black"} // blocking

		start := Position{row: 4, col: 4}
		end := Position{row: 4, col: 0}

		if rook.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move due to blocking piece")
		}
	})

	t.Run("rook cannot stay in same square", func(t *testing.T) {
		board := &Board{}
		rook := &Rook{color: "black"}

		board.grid[4][4] = rook

		start := Position{row: 4, col: 4}
		end := Position{row: 4, col: 4}

		if rook.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move to same square")
		}
	})
}
