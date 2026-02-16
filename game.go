package main

import "fmt"

func Start() {
	board := NewBoard()
	board.Print()

	err := board.Move(Position{row: 6, col: 0}, Position{row: 7, col: 0})
	if err != nil {
		fmt.Println(err)
	}

	board.Print()
}
