package main

type King struct {
	color    string
	hasMoved bool
}

func (p *King) Symbol() string {
	if p.color == "white" {
		return "K"
	}

	return "k"
}

func (p *King) Color() string {
	return p.color
}

func (p *King) Move() {
	p.hasMoved = true
}
