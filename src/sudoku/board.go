package sudoku

import "fmt"

type Board struct {
	row     int
	columns int
	Squares [3][3]Square
}

func CreateBoard() (b Board) {
	b.row = 3
	b.columns = 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.Squares[i][j] = CreateSquare()
		}
	}

	return
}

func (b Board) PrintBoard() {
	for i := 0; i < 3; i++ {
		fmt.Println("========================================")
		for j := 0; j < 3; j++ {
			line := "||"
			for g := 0; g < 3; g++ {
				line += b.Squares[i][j].PrintLine(g) + "||"
			}
			fmt.Println(line)
		}
	}

	fmt.Println("========================================")
}
