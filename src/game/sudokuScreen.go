package game

import (
	"Sudoku-Go/src/sudoku"
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

type matrixCell struct {
	b *button
	// if it is a number that came with the puzzle
	initial  bool
	row, col int
}

// global variables to store the UI elements, so no redrawing is needed
var matrixCells [9][9]matrixCell
var boundaryX, boundaryY [4]button
var buttons [9]button

// save globally the current matrix and the selected cell
var grid [][]int
var selectedCell matrixCell

func startNewSudokuGame(b *button) {
	grid = sudoku.CreateGrid()
	selectedCell = matrixCell{b: nil, row: 0, col: 0}
	createMatrix()
	createBoundaries()
	createNumInput()
	renderSudokuScreen()
}

func renderSudokuScreen() {
	for {
		event := sdl.WaitEvent()
		switch event.(type) {
		case *sdl.MouseButtonEvent:
			processButtonEvents(event)
		case *sdl.QuitEvent:
			closeGame()
			return
		}

		renderer.SetDrawColor(46, 42, 56, 255)
		renderer.Clear()

		drawMatrix()
		drawNumbInput()

		renderer.Present()
	}
}

func createMatrix() {
	startX, startY, inc := 35, 65, 60

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			matrixCells[row][col].row = row
			matrixCells[row][col].col = col

			numb := strconv.Itoa(grid[row][col])
			if numb == "0" {
				numb = " "
				matrixCells[row][col].initial = false
			} else {
				matrixCells[row][col].initial = true
			}

			cell := createButton(
				&sdl.Rect{X: int32(startX + col*inc), Y: int32(startY + row*inc), W: 50, H: 50},
				&sdl.Color{R: 214, G: 214, B: 214, A: 255},
				&sdl.Color{R: 0, G: 0, B: 0, A: 0},
				numb,
				24,
				changeColor)

			matrixCells[row][col].b = &cell
		}
	}
}

func createBoundaries() {
	startX, startY, inc := 35, 65, 180

	for i := 0; i < 4; i++ {
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

func drawMatrix() {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			matrixCells[row][col].b.drawButton()
		}
	}

	for i := 0; i < 4; i++ {
		boundaryX[i].drawButton()
		boundaryY[i].drawButton()
	}
}

func createNumInput() {
	startX, startY, inc := 35, 700, 60

	for i := 0; i < 9; i++ {
		buttons[i] = createButton(
			&sdl.Rect{X: int32(startX + i*inc), Y: int32(startY), W: 50, H: 50},
			&sdl.Color{R: 214, G: 214, B: 214, A: 255},
			&sdl.Color{R: 0, G: 0, B: 0, A: 0},
			strconv.Itoa(i+1),
			24,
			changeCellNum)
	}
}

func drawNumbInput() {
	for i := 0; i < 9; i++ {
		buttons[i].drawButton()
	}
}

func changeColor(b *button) {
	// Check if color is the original one (grey)
	if b.Color.r == 214 {
		b.Color = struct {
			r uint8
			g uint8
			b uint8
			a uint8
		}{r: 240, g: 228, b: 81, a: 255}

		if selectedCell.b != nil {
			changeColor(selectedCell.b)
		}

		selectedCell.b = b

		setAvailableOpt(sudoku.GetImpossibleNum(grid, selectedCell.row, selectedCell.col))
	} else {
		b.Color = struct {
			r uint8
			g uint8
			b uint8
			a uint8
		}{r: 214, g: 214, b: 214, a: 255}

		selectedCell.b = nil
	}
}

func changeCellNum(b *button) {
	num, _ := strconv.Atoi(b.Text.text)
	if selectedCell.b != nil && sudoku.CheckValue(grid, selectedCell.row, selectedCell.col, num) {
		selectedCell.b.Text = b.Text
		_, selectedCell.b.Text.textTex, _ = createText(b.Text.text, 24, 110, 110, 110, 255)
		changeColor(selectedCell.b)
		grid[selectedCell.row][selectedCell.col] = num
	}
}

func setAvailableOpt(opt []int) {
	resetOpt()
	for i := 0; i < len(opt); i++ {
		buttons[opt[i]-1].Color = struct {
			r uint8
			g uint8
			b uint8
			a uint8
		}{r: 125, g: 125, b: 125, a: 255}
	}
}

func resetOpt() {
	for i := 0; i < len(buttons); i++ {
		buttons[i].Color = struct {
			r uint8
			g uint8
			b uint8
			a uint8
		}{r: 214, g: 214, b: 214, a: 255}
	}
}

func processButtonEvents(event sdl.Event) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if !matrixCells[row][col].initial {
				if matrixCells[row][col].b.processEvent(event) {
					selectedCell.row = row
					selectedCell.col = col
					break
				}
			}

			if checkCounter() == 0 && sudoku.CheckSolution(grid) {
				fmt.Println("YOU HAVE WON!")
			}
		}
	}

	for i := 0; i < 9; i++ {
		buttons[i].processEvent(event)
	}
}

func checkCounter() int {
	counter := 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				counter++
			}
		}
	}

	return counter
}
