package main

type Queen struct {
	color string
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
