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
