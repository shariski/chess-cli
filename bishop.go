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
