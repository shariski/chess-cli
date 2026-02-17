package main

type Knight struct {
	color    string
	hasMoved bool
}

func (p *Knight) Symbol() string {
	if p.color == "white" {
		return "N"
	}

	return "n"
}

func (p *Knight) Color() string {
	return p.color
}

func (p *Knight) Move() {
	p.hasMoved = true
}

func (p *Knight) IsValidMove(board *Board, start, end Position) bool {
	dRow := AbsInt(end.row - start.row)
	dCol := AbsInt(end.col - start.col)

	if (dRow == 1 && dCol == 2) || (dRow == 2 && dCol == 1) {
		return true
	}

	return false
}
