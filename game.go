package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		}

		err = g.board.Move(*start, *end)
		if err != nil {
			fmt.Println(err)
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

	startChar := strings.Split(inputParts[0], ",")
	endChar := strings.Split(inputParts[1], ",")

	startRow, err := strconv.Atoi(startChar[0])
	if err != nil {
		return nil, nil, err
	}

	startCol, err := strconv.Atoi(startChar[1])
	if err != nil {
		return nil, nil, err
	}

	endRow, err := strconv.Atoi(endChar[0])
	if err != nil {
		return nil, nil, err
	}

	endCol, err := strconv.Atoi(endChar[1])
	if err != nil {
		return nil, nil, err
	}

	start := &Position{row: startRow, col: startCol}
	end := &Position{row: endRow, col: endCol}

	return start, end, nil
}
