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

func (p *Bishop) IsValidMove(start, end Position) bool {
	dCol := AbsInt(end.col - start.col)
	if dCol != 0 {
		return false
	}

	return true
}
