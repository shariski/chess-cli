package main

type Knight struct {
	color string
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
