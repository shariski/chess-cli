package main

import "testing"

func TestBishop_IsValidMove(t *testing.T) {
	t.Run("white bishop move forward diagonally", func(t *testing.T) {
		board := &Board{}
		bishop := &Bishop{color: "white"}

		board.grid[7][2] = bishop

		start := Position{row: 7, col: 2}
		end := Position{row: 4, col: 5}

		if !bishop.IsValidMove(board, start, end) {
			t.Fatalf("expected valid move")
		}
	})

	t.Run("white bishop move forward diagonally incorrect", func(t *testing.T) {
		board := &Board{}
		bishop := &Bishop{color: "white"}

		board.grid[7][2] = bishop

		start := Position{row: 7, col: 2}
		end := Position{row: 4, col: 6}

		if bishop.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move")
		}
	})

	t.Run("bishop move backward diagonal", func(t *testing.T) {
		board := &Board{}
		bishop := &Bishop{color: "white"}

		board.grid[4][4] = bishop

		start := Position{row: 4, col: 4}
		end := Position{row: 6, col: 6}

		if !bishop.IsValidMove(board, start, end) {
			t.Fatalf("expected valid backward diagonal move")
		}
	})

	t.Run("bishop cannot move straight", func(t *testing.T) {
		board := &Board{}
		bishop := &Bishop{color: "white"}

		board.grid[4][4] = bishop

		start := Position{row: 4, col: 4}
		end := Position{row: 4, col: 7}

		if bishop.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid straight move")
		}
	})

	t.Run("bishop path blocked", func(t *testing.T) {
		board := &Board{}
		bishop := &Bishop{color: "white"}

		board.grid[4][4] = bishop
		board.grid[3][5] = &Pawn{color: "white"} // blocking square

		start := Position{row: 4, col: 4}
		end := Position{row: 2, col: 6}

		if bishop.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move due to blocking piece")
		}
	})

	t.Run("bishop capture enemy", func(t *testing.T) {
		board := &Board{}
		bishop := &Bishop{color: "white"}

		board.grid[4][4] = bishop
		board.grid[2][6] = &Pawn{color: "black"}

		start := Position{row: 4, col: 4}
		end := Position{row: 2, col: 6}

		if !bishop.IsValidMove(board, start, end) {
			t.Fatalf("expected valid capture")
		}
	})

	t.Run("bishop cannot stay in same square", func(t *testing.T) {
		board := &Board{}
		bishop := &Bishop{color: "white"}

		board.grid[4][4] = bishop

		start := Position{row: 4, col: 4}
		end := Position{row: 4, col: 4}

		if bishop.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move to same square")
		}
	})
}
