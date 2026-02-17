package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	board       *Board
	currentTurn string
}

func NewGame() *Game {
	return &Game{board: NewBoard(), currentTurn: "white"}
}

func (g *Game) Start() {
	for {
		g.board.Render()

		fmt.Printf("%s: ", g.current())

		// scan input position
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		start, end, err := g.parseInput(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// validate current piece having same color with current turn
		currentPiece := g.board.Get(*start)

		if currentPiece == nil {
			fmt.Println("square is empty")
			continue
		}

		if currentPiece.Color() != g.current() {
			fmt.Println("not your piece")
			continue
		}

		err = g.board.Move(*start, *end)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// movement happened, switch turn
		g.switchTurn()
	}
}

func (g *Game) switchTurn() {
	if g.currentTurn == "white" {
		g.currentTurn = "black"
	} else {
		g.currentTurn = "white"
	}
}

func (g *Game) current() string {
	return g.currentTurn
}

func (g *Game) parseInput(input string) (*Position, *Position, error) {
	inputParts := strings.Fields(input)

	if len(inputParts) != 2 {
		return nil, nil, fmt.Errorf("invalid input")
	}

	start, err := chessToIndex(inputParts[0])
	if err != nil {
		return nil, nil, err
	}

	end, err := chessToIndex(inputParts[1])
	if err != nil {
		return nil, nil, err
	}

	return start, end, nil
}

func chessToIndex(pos string) (*Position, error) {
	if len(pos) != 2 {
		return nil, fmt.Errorf("invalid position")
	}

	file := pos[0]
	rank := pos[1]

	if file < 'a' || file > 'h' || rank < '1' || rank > '8' {
		return nil, fmt.Errorf("invalid position")
	}

	col := int(file - 'a')
	row := 7 - int(rank-'1')

	return &Position{row: row, col: col}, nil
}
