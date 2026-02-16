package main

type Rook struct {
	color string
}

func (p *Rook) Symbol() string {
	if p.color == "white" {
		return "R"
	}

	return "r"
}

func (p *Rook) Color() string {
	return p.color
}
