package main

import "testing"

func TestQueen_IsValidMove(t *testing.T) {
	t.Run("white queen move forward", func(t *testing.T) {
		board := &Board{}
		queen := &Queen{color: "white"}

		board.grid[7][0] = queen

		start := Position{row: 7, col: 0}
		end := Position{row: 0, col: 0}

		if !queen.IsValidMove(board, start, end) {
			t.Fatalf("expected valid move")
		}
	})

	t.Run("queen move vertical", func(t *testing.T) {
		board := &Board{}
		queen := &Queen{color: "white"}

		board.grid[7][0] = queen

		start := Position{row: 7, col: 0}
		end := Position{row: 0, col: 0}

		if !queen.IsValidMove(board, start, end) {
			t.Fatalf("expected valid vertical move")
		}
	})

	t.Run("queen move horizontal", func(t *testing.T) {
		board := &Board{}
		queen := &Queen{color: "white"}

		board.grid[4][4] = queen

		start := Position{row: 4, col: 4}
		end := Position{row: 4, col: 0}

		if !queen.IsValidMove(board, start, end) {
			t.Fatalf("expected valid horizontal move")
		}
	})

	t.Run("queen move diagonal", func(t *testing.T) {
		board := &Board{}
		queen := &Queen{color: "white"}

		board.grid[4][4] = queen

		start := Position{row: 4, col: 4}
		end := Position{row: 1, col: 7}

		if !queen.IsValidMove(board, start, end) {
			t.Fatalf("expected valid diagonal move")
		}
	})

	t.Run("queen cannot move invalid pattern", func(t *testing.T) {
		board := &Board{}
		queen := &Queen{color: "white"}

		board.grid[4][4] = queen

		start := Position{row: 4, col: 4}
		end := Position{row: 6, col: 7}

		if queen.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move pattern")
		}
	})

	t.Run("queen path blocked vertically", func(t *testing.T) {
		board := &Board{}
		queen := &Queen{color: "white"}

		board.grid[7][0] = queen
		board.grid[4][0] = &Pawn{color: "white"} // blocking

		start := Position{row: 7, col: 0}
		end := Position{row: 0, col: 0}

		if queen.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move due to blocking piece")
		}
	})

	t.Run("queen path blocked diagonally", func(t *testing.T) {
		board := &Board{}
		queen := &Queen{color: "white"}

		board.grid[4][4] = queen
		board.grid[3][5] = &Pawn{color: "white"} // blocking diagonal

		start := Position{row: 4, col: 4}
		end := Position{row: 1, col: 7}

		if queen.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid diagonal move due to blocking piece")
		}
	})

	t.Run("queen cannot stay in same square", func(t *testing.T) {
		board := &Board{}
		queen := &Queen{color: "white"}

		board.grid[4][4] = queen

		start := Position{row: 4, col: 4}
		end := Position{row: 4, col: 4}

		if queen.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move to same square")
		}
	})
}
