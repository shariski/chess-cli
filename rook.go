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
	dRow := AbsInt(end.row - start.row)
	dCol := AbsInt(end.col - start.col)

	if !((dRow > 0 && dCol == 0) || (dRow == 0 && dCol > 0)) {
		return false
	}

	if !board.IsPathClear(start, end) {
		return false
	}

	return true
}
