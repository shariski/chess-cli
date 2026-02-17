package main

type Pawn struct {
	color    string
	hasMoved bool
}

func (p *Pawn) Symbol() string {
	if p.color == "white" {
		return "P"
	}

	return "p"
}

func (p *Pawn) Color() string {
	return p.color
}

func (p *Pawn) Move() {
	p.hasMoved = true
}

func (p *Pawn) IsValidMove(board *Board, start, end Position) bool {
	direction := 1
	if p.color == "white" {
		direction = -1
	}

	dRow := end.row - start.row
	dCol := end.col - start.col

	// move forward one square
	if dRow == 1*direction && dCol == 0 && board.IsEmpty(end) {
		return true
	}

	// move forward two squares
	if !p.hasMoved && dRow == 2*direction && dCol == 0 && board.IsEmpty(end) && board.IsEmpty(Position{row: start.row + direction, col: start.col}) {
		return true
	}

	return false
}
