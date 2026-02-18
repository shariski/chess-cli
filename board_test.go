package main

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	tests := []struct {
		row, col int
		piece    any
		symbol   string
		color    string
	}{
		// BLACK BACK RANK
		{0, 0, &Rook{}, "r", "black"},
		{0, 1, &Knight{}, "n", "black"},
		{0, 2, &Bishop{}, "b", "black"},
		{0, 3, &Queen{}, "q", "black"},
		{0, 4, &King{}, "k", "black"},
		{0, 5, &Bishop{}, "b", "black"},
		{0, 6, &Knight{}, "n", "black"},
		{0, 7, &Rook{}, "r", "black"},

		// BLACK PAWNS
		{1, 0, &Pawn{}, "p", "black"},
		{1, 1, &Pawn{}, "p", "black"},
		{1, 2, &Pawn{}, "p", "black"},
		{1, 3, &Pawn{}, "p", "black"},
		{1, 4, &Pawn{}, "p", "black"},
		{1, 5, &Pawn{}, "p", "black"},
		{1, 6, &Pawn{}, "p", "black"},
		{1, 7, &Pawn{}, "p", "black"},

		// WHITE PAWNS
		{6, 0, &Pawn{}, "P", "white"},
		{6, 1, &Pawn{}, "P", "white"},
		{6, 2, &Pawn{}, "P", "white"},
		{6, 3, &Pawn{}, "P", "white"},
		{6, 4, &Pawn{}, "P", "white"},
		{6, 5, &Pawn{}, "P", "white"},
		{6, 6, &Pawn{}, "P", "white"},
		{6, 7, &Pawn{}, "P", "white"},

		// WHITE BACK RANK
		{7, 0, &Rook{}, "R", "white"},
		{7, 1, &Knight{}, "N", "white"},
		{7, 2, &Bishop{}, "B", "white"},
		{7, 3, &Queen{}, "Q", "white"},
		{7, 4, &King{}, "K", "white"},
		{7, 5, &Bishop{}, "B", "white"},
		{7, 6, &Knight{}, "N", "white"},
		{7, 7, &Rook{}, "R", "white"},
	}

	for _, tt := range tests {
		p := board.grid[tt.row][tt.col]

		if p == nil {
			t.Fatalf("expected piece at [%d][%d], got nil", tt.row, tt.col)
		}

		if reflect.TypeOf(p) != reflect.TypeOf(tt.piece) {
			t.Fatalf("wrong type at [%d][%d]: got %T",
				tt.row, tt.col, p)
		}

		if p.Symbol() != tt.symbol {
			t.Fatalf("wrong symbol at [%d][%d]: got %s",
				tt.row, tt.col, p.Symbol())
		}

		if p.Color() != tt.color {
			t.Fatalf("wrong color at [%d][%d]: got %s",
				tt.row, tt.col, p.Color())
		}
	}

	// Validate for nil squares
	for row := 2; row <= 5; row++ {
		for col := 0; col < 8; col++ {
			if board.grid[row][col] != nil {
				t.Fatalf("expected nil at [%d][%d], got %T",
					row, col, board.grid[row][col])
			}
		}
	}

	// Validate board labels
	rowLabel := "87654321"
	if board.rowIdx != rowLabel {
		t.Fatalf("expected: %s, got: %s", rowLabel, board.rowIdx)
	}

	colLabel := "abcdefgh"
	if board.colIdx != colLabel {
		t.Fatalf("expected: %s, got: %s", colLabel, board.colIdx)
	}
}

func TestBoard_Move(t *testing.T) {

	t.Run("cannot move from empty square", func(t *testing.T) {
		board := &Board{}
		game := &Game{}

		err := board.Move(game, Position{0, 0}, Position{0, 1})

		if err == nil {
			t.Fatalf("expected error when moving from empty square")
		}
	})

	t.Run("cannot capture teammate", func(t *testing.T) {
		board := &Board{}
		game := &Game{}

		king := &King{color: "white"}
		pawn := &Pawn{color: "white"}

		board.grid[3][3] = king
		board.grid[4][4] = pawn

		err := board.Move(game,
			Position{4, 4},
			Position{3, 3},
		)

		if err == nil {
			t.Fatalf("expected error when capturing teammate")
		}
	})

	t.Run("successful move updates board", func(t *testing.T) {
		board := &Board{}
		game := &Game{}

		rook := &Rook{color: "white"}

		board.grid[4][4] = rook

		err := board.Move(game,
			Position{4, 4},
			Position{4, 7},
		)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if board.grid[4][4] != nil {
			t.Fatalf("expected start square to be empty")
		}

		if board.grid[4][7] != rook {
			t.Fatalf("expected rook to move to end square")
		}
	})

	t.Run("capture enemy piece", func(t *testing.T) {
		board := &Board{}
		game := &Game{}

		rook := &Rook{color: "white"}
		pawn := &Pawn{color: "black"}

		board.grid[4][4] = rook
		board.grid[4][7] = pawn

		err := board.Move(game,
			Position{4, 4},
			Position{4, 7},
		)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if board.grid[4][7] != rook {
			t.Fatalf("expected rook to capture pawn")
		}
	})

	t.Run("capture king sets winner", func(t *testing.T) {
		board := &Board{}
		game := &Game{}

		rook := &Rook{color: "white"}
		king := &King{color: "black"}

		board.grid[4][4] = rook
		board.grid[4][7] = king

		err := board.Move(game,
			Position{4, 4},
			Position{4, 7},
		)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if game.winner != "white" {
			t.Fatalf("expected winner to be white")
		}
	})

	t.Run("pawn move updates hasMoved", func(t *testing.T) {
		board := &Board{}
		game := &Game{}

		pawn := &Pawn{color: "white"}

		board.grid[6][0] = pawn

		err := board.Move(game,
			Position{6, 0},
			Position{5, 0},
		)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !pawn.hasMoved {
			t.Fatalf("expected pawn.hasMoved to be true")
		}
	})
}

func TestBoard_Move_CaptureKing(t *testing.T) {
	board := &Board{}
	game := &Game{}

	rook := &Rook{color: "white"}
	blackKing := &King{color: "black"}

	board.grid[4][4] = rook
	board.grid[4][7] = blackKing

	err := board.Move(game,
		Position{4, 4},
		Position{4, 7},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if game.winner != "white" {
		t.Fatalf("expected winner to be white, got %s", game.winner)
	}

	if !game.gameOver {
		t.Fatalf("expected gameOver to be true")
	}
}
