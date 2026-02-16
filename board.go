package main

type Piece interface {
	Symbol() string
	Color() string
}

func InitBoard() [8][8]Piece {
	var board [8][8]Piece

	board[0][0] = &Rook{color: "black"}
	board[0][1] = &Knight{color: "black"}
	board[0][2] = &Bishop{color: "black"}
	board[0][3] = &Queen{color: "black"}
	board[0][4] = &King{color: "black"}
	board[0][5] = &Bishop{color: "black"}
	board[0][6] = &Knight{color: "black"}
	board[0][7] = &Rook{color: "black"}

	return board
}
