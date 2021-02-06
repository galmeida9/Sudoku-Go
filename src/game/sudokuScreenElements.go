package game

import (
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

//-----------------------------------------------------------------------//
// global variables to store the UI elements, so no redrawing is needed
//-----------------------------------------------------------------------//

var matrixCells [rowSize][colSize]matrixCell
var boundaryX, boundaryY [boundarySize]button
var buttons [colSize]button
var clearButton, clearAllButton button

//----------------------------//
// Create UI elements
//----------------------------//

func createMatrix() {
	startX, startY, inc := 35, 75, 60

	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			matrixCells[row][col].Row = row
			matrixCells[row][col].Col = col

			numb := strconv.Itoa(grid[row][col])
			color := &sdl.Color{R: 0, G: 0, B: 0, A: 0}

			if numb == "0" {
				numb = " "
				matrixCells[row][col].Initial = false
			} else if originalGrid[row][col] != grid[row][col] {
				matrixCells[row][col].Initial = false
				color = &sdl.Color{R: 110, G: 110, B: 110, A: 255}
			} else {
				matrixCells[row][col].Initial = true
			}

			cell := createButton(
				&sdl.Rect{X: int32(startX + col*inc), Y: int32(startY + row*inc), W: 50, H: 50},
				&sdl.Color{R: 214, G: 214, B: 214, A: 255},
				color,
				numb,
				24,
				changeColor)

			matrixCells[row][col].B = &cell
		}
	}
}

func createBoundaries() {
	startX, startY, inc := 35, 75, 180

	for i := 0; i < boundarySize; i++ {
		boundaryX[i] = createButton(
			&sdl.Rect{X: int32(startX + i*inc - 10), Y: int32(startY), W: 10, H: 60 * 9},
			&sdl.Color{R: 125, G: 125, B: 125, A: 255},
			&sdl.Color{R: 0, G: 0, B: 0, A: 0},
			"1",
			1,
			nil)

		boundaryY[i] = createButton(
			&sdl.Rect{X: int32(startX - 10), Y: int32(startY + i*inc - 10), W: 60*9 + 10, H: 10},
			&sdl.Color{R: 125, G: 125, B: 125, A: 255},
			&sdl.Color{R: 0, G: 0, B: 0, A: 0},
			"1",
			1,
			nil)
	}
}

func createNumInput() {
	startX, startY, inc := 35, 700, 60

	for i := range buttons {
		buttons[i] = createButton(
			&sdl.Rect{X: int32(startX + i*inc), Y: int32(startY), W: 50, H: 50},
			&sdl.Color{R: 214, G: 214, B: 214, A: 255},
			&sdl.Color{R: 0, G: 0, B: 0, A: 0},
			strconv.Itoa(i+1),
			24,
			changeCellNum)
	}
}

func createClearButtons() {
	clearButton = createButton(
		&sdl.Rect{X: 25, Y: 15, W: 100, H: 40},
		&sdl.Color{R: 125, G: 125, B: 125, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		"Clear",
		24,
		clearNum)

	clearAllButton = createButton(
		&sdl.Rect{X: 135, Y: 15, W: 100, H: 40},
		&sdl.Color{R: 125, G: 125, B: 125, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		"Clear All",
		24,
		clearGame)
}

//----------------------------//
// Draw UI elements
//----------------------------//

func drawMatrix() {
	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			matrixCells[row][col].B.drawButton()
		}
	}

	for i := 0; i < boundarySize; i++ {
		boundaryX[i].drawButton()
		boundaryY[i].drawButton()
	}
}

func drawNumbInput() {
	for i := range buttons {
		buttons[i].drawButton()
	}
}
