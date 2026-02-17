package main

type Bishop struct {
	color    string
	hasMoved bool
}

func (p *Bishop) Symbol() string {
	if p.color == "white" {
		return "B"
	}

	return "b"
}

func (p *Bishop) Color() string {
	return p.color
}

func (p *Bishop) Move() {
	p.hasMoved = true
}

func (p *Bishop) IsValidMove(board *Board, start, end Position) bool {
	dRow := AbsInt(end.row - start.row)
	dCol := AbsInt(end.col - start.col)

	if dRow == dCol && dRow != 0 {
		return true
	}

	return false
}
