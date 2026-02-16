package main

type King struct {
	color string
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
