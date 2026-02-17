package main

type Rook struct {
	color    string
	hasMoved bool
}

func (p *Rook) Symbol() string {
	if p.color == "white" {
		return "R"
	}

	return "r"
}

func (p *Rook) Color() string {
	return p.color
}

func (p *Rook) Move() {
	p.hasMoved = true
}

func (p *Rook) IsValidMove(board *Board, start, end Position) bool {
	dCol := AbsInt(end.col - start.col)
	if dCol != 0 {
		return false
	}

	return true
}
