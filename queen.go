package main

type Queen struct {
	color    string
	hasMoved bool
}

func (p *Queen) Symbol() string {
	if p.color == "white" {
		return "Q"
	}

	return "q"
}

func (p *Queen) Color() string {
	return p.color
}

func (p *Queen) Move() {
	p.hasMoved = true
}

func (p *Queen) IsValidMove(board *Board, start, end Position) bool {
	dRow := AbsInt(end.row - start.row)
	dCol := AbsInt(end.col - start.col)

	if !(dRow == dCol && dRow != 0) {
		return false
	}

	if !((dRow > 0 && dCol == 0) || (dRow == 0 && dCol > 0)) {
		return false
	}

	if !board.IsPathClear(start, end) {
		return false
	}

	return true
}
