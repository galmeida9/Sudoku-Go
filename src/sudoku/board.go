package sudoku

import "fmt"

type Board struct {
	Squares [3][3]Square
}

func CreateBoard() (b Board) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.Squares[i][j] = CreateSquare()
		}
	}

	return
}

func PrintBoard(b Board) {
	for i := 0; i < 3; i++ {
		fmt.Println("========================================")
		for j := 0; j < 3; j++ {
			var line string = "||"
			for g := 0; g < 3; g++ {
				line += PrintLine(b.Squares[i][j], g) + "||"
			}
			fmt.Println(line)
		}
	}

	fmt.Println("========================================")
}
