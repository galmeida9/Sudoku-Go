package main

import (
	"Sudoku-Go/src/sudoku"
)

func main() {
	var board = sudoku.CreateBoard()
	sudoku.PrintBoard(board)
}
