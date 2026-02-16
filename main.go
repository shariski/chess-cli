package main

import "fmt"

func main() {
	board := InitBoard()

	for _, inner := range board {
		for _, value := range inner {
			if value != nil {
				fmt.Printf("%s ", value.Symbol())
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}
