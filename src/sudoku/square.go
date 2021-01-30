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

	filledNumb := 4
	for filledNumb > 0 {
		column := rand.Intn(3)
		row := rand.Intn(3)

		if square.values[row][column] == 0 {
			square.values[row][column] = rand.Intn(9) + 1
			filledNumb--
		}
	}

	return square
}

func (s Square) PrintLine(row int) (line string) {
	for i := 0; i < s.columns; i++ {
		if s.values[row][i] == 0 {
			line += "   "
		} else {
			line += " " + strconv.Itoa(s.values[row][i]) + " "
		}

		if i < s.columns-1 {
			line += "|"
		}
	}

	return
}
