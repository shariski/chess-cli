package main

type King struct {
	color    string
	hasMoved bool
}

func (p *King) Symbol() string {
	if p.color == "white" {
		return "K"
	}

	return "k"
}

func (p *King) Color() string {
	return p.color
}

func (p *King) Move() {
	p.hasMoved = true
}

func (p *King) IsValidMove(board *Board, start Position, end Position) bool {
	dRow := AbsInt(end.row - start.row)
	dCol := AbsInt(end.col - start.col)

	if dRow <= 1 && dCol <= 1 && (dRow+dCol) != 0 {
		return true
	}

	return false
}
