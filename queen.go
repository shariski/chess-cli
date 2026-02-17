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
