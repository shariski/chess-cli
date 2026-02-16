package main

type Pawn struct {
	color string
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
