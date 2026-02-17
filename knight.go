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
