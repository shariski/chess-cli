package main

type Bishop struct {
	color string
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
