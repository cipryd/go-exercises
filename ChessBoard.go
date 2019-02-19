package main

import "fmt"

func setBoard() [8][8]rune {
	var board [8][8]rune

	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}

	board[0][0] = 'r'
	board[0][7] = 'r'
	board[7][0] = 'R'
	board[7][7] = 'R'

	board[0][1] = 'n'
	board[0][6] = 'n'
	board[7][1] = 'N'
	board[7][6] = 'N'

	board[0][2] = 'b'
	board[0][5] = 'b'
	board[7][2] = 'B'
	board[7][5] = 'B'

	board[0][3] = 'q'
	board[7][3] = 'Q'

	board[0][4] = 'k'
	board[7][4] = 'K'

	return board
}

func printBoard(board [8][8]rune) {
	for line := range board {
		for i := 0; i <= 4*len(board); i++ {
			fmt.Printf("=")
		}
		fmt.Printf("\n| ")
		for column := range board[line] {
			fmt.Printf("%c | ", board[line][column])
		}
		fmt.Println()
	}
	for i := 0; i <= 4*len(board); i++ {
		fmt.Printf("=")
	}
	fmt.Println()
}

func printChessBoard() {
	board := setBoard()
	printBoard(board)
}
