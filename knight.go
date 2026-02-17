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

func (p *Knight) IsValidMove(start, end Position) bool {
	dCol := AbsInt(end.col - start.col)
	if dCol != 0 {
		return false
	}

	return true
}
