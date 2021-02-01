package game

import (
	"Sudoku-Go/src/sudoku"
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

var grid [][]int
var matrixCells [9][9]button
var selectedCell *button
var buttons [9]button

func startNewSudokuGame(b *button) {
	grid = sudoku.CreateGrid()
	fmt.Println(grid)
	createMatrix()
	createNumInput()
	renderSudokuScreen()
}

func renderSudokuScreen() {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.MouseButtonEvent:
				processButtonEvents(event)
			case *sdl.QuitEvent:
				closeGame()
				return
			}
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
			numb := strconv.Itoa(grid[row][col])
			if numb == "0" {
				numb = " "
			}

			matrixCells[row][col] = createButton(
				&sdl.Rect{X: int32(startX + col*inc), Y: int32(startY + row*inc), W: 50, H: 50},
				&sdl.Color{R: 214, G: 214, B: 214, A: 255},
				&sdl.Color{R: 0, G: 0, B: 0, A: 0},
				numb,
				24,
				changeColor)
		}
	}
}

func drawMatrix() {
	startX, startY, inc := 35, 65, 60

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			//draw divisions
			if col%3 == 0 {
				division := createButton(
					&sdl.Rect{X: int32(startX + col*inc - 10), Y: int32(startY + row*inc), W: 10, H: 60},
					&sdl.Color{R: 125, G: 125, B: 125, A: 255},
					&sdl.Color{R: 0, G: 0, B: 0, A: 0},
					"1",
					1,
					nil)
				division.drawButton()
			}

			if row%3 == 0 {
				division := createButton(
					&sdl.Rect{X: int32(startX + col*inc - 10), Y: int32(startY + row*inc - 10), W: 60, H: 10},
					&sdl.Color{R: 125, G: 125, B: 125, A: 255},
					&sdl.Color{R: 0, G: 0, B: 0, A: 0},
					"1",
					1,
					nil)
				division.drawButton()
			}

			matrixCells[row][col].drawButton()
		}
	}

	division := createButton(
		&sdl.Rect{X: int32(startX + 9*inc - 10), Y: int32(startY - 10), W: 10, H: 60 * 9},
		&sdl.Color{R: 125, G: 125, B: 125, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		"1",
		1,
		nil)
	division.drawButton()

	division = createButton(
		&sdl.Rect{X: int32(startX - 10), Y: int32(startY + 9*inc - 10), W: 60*9 + 10, H: 10},
		&sdl.Color{R: 125, G: 125, B: 125, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		"1",
		1,
		nil)
	division.drawButton()
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

		if selectedCell != nil {
			changeColor(selectedCell)
		}

		selectedCell = b
	} else {
		b.Color = struct {
			r uint8
			g uint8
			b uint8
			a uint8
		}{r: 214, g: 214, b: 214, a: 255}

		selectedCell = nil
	}
}

func changeCellNum(b *button) {
	if selectedCell != nil {
		selectedCell.Text = b.Text
		_, selectedCell.Text.textTex, _ = createText(b.Text.text, 24, 110, 110, 110, 255)
		changeColor(selectedCell)
	}
}

func processButtonEvents(event sdl.Event) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			matrixCells[row][col].processEvent(event)
		}
	}

	for i := 0; i < 9; i++ {
		buttons[i].processEvent(event)
	}
}
