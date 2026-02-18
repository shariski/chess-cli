package main

import "testing"

func TestKing_IsValidMove(t *testing.T) {

	t.Run("king move one step vertical", func(t *testing.T) {
		board := &Board{}
		king := &King{color: "white"}

		board.grid[4][4] = king

		start := Position{row: 4, col: 4}
		end := Position{row: 3, col: 4}

		if !king.IsValidMove(board, start, end) {
			t.Fatalf("expected valid vertical move")
		}
	})

	t.Run("king move one step horizontal", func(t *testing.T) {
		board := &Board{}
		king := &King{color: "white"}

		board.grid[4][4] = king

		start := Position{row: 4, col: 4}
		end := Position{row: 4, col: 5}

		if !king.IsValidMove(board, start, end) {
			t.Fatalf("expected valid horizontal move")
		}
	})

	t.Run("king move one step diagonal", func(t *testing.T) {
		board := &Board{}
		king := &King{color: "white"}

		board.grid[4][4] = king

		start := Position{row: 4, col: 4}
		end := Position{row: 3, col: 5}

		if !king.IsValidMove(board, start, end) {
			t.Fatalf("expected valid diagonal move")
		}
	})

	t.Run("king cannot move two squares vertically", func(t *testing.T) {
		board := &Board{}
		king := &King{color: "white"}

		board.grid[4][4] = king

		start := Position{row: 4, col: 4}
		end := Position{row: 2, col: 4}

		if king.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move (too far vertically)")
		}
	})

	t.Run("king cannot move like knight", func(t *testing.T) {
		board := &Board{}
		king := &King{color: "white"}

		board.grid[4][4] = king

		start := Position{row: 4, col: 4}
		end := Position{row: 6, col: 5}

		if king.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid knight-like move")
		}
	})

	t.Run("king cannot stay in same square", func(t *testing.T) {
		board := &Board{}
		king := &King{color: "white"}

		board.grid[4][4] = king

		start := Position{row: 4, col: 4}
		end := Position{row: 4, col: 4}

		if king.IsValidMove(board, start, end) {
			t.Fatalf("expected invalid move to same square")
		}
	})
}
