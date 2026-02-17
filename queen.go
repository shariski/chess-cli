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

	isDiagonal := dRow == dCol && dRow != 0
	isStraight := (dRow > 0 && dCol == 0) || (dRow == 0 && dCol > 0)

	if !isDiagonal && !isStraight {
		return false
	}

	if !board.IsPathClear(start, end) {
		return false
	}

	return true
}
