package sudoku

import (
	"math/rand"
	"strconv"
)

type Square struct {
	row     int
	columns int
	values  [][]int
}

func CreateSquare() Square {
	square := Square{
		row:     3,
		columns: 3,
		values:  [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
	}

	var filledNumb int = 4
	for filledNumb > 0 {
		var column int = rand.Intn(3)
		var row int = rand.Intn(3)

		if square.values[row][column] == 0 {
			square.values[row][column] = rand.Intn(9) + 1
			filledNumb--
		}
	}

	return square
}

func PrintLine(square Square, row int) (line string) {
	for i := 0; i < square.columns; i++ {
		if square.values[row][i] == 0 {
			line += "   "
		} else {
			line += " " + strconv.Itoa(square.values[row][i]) + " "
		}

		if i < square.columns-1 {
			line += "|"
		}
	}

	return
}
