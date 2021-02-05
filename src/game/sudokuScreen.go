package game

import (
	"Sudoku-Go/src/sudoku"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

type matrixCell struct {
	B *button
	// if it is a number that came with the puzzle
	Initial  bool
	Row, Col int
}

type saveStatus struct {
	Sol    [9][]int
	Cur    [][]int
	Org    [9][]int
	IsEasy bool
}

const (
	rowSize      = 9
	colSize      = 9
	boundarySize = 4
)

//-----------------------------------------------------------------------//
// global variables to store the UI elements, so no redrawing is needed
//-----------------------------------------------------------------------//

var matrixCells [rowSize][colSize]matrixCell
var boundaryX, boundaryY [boundarySize]button
var buttons [colSize]button
var clearButton, clearAllButton button

//-----------------------------------------------------------------------//
// save globally the current matrix and the selected cell
//-----------------------------------------------------------------------//

var grid [][]int
var originalGrid [9][]int
var selectedCell matrixCell

var won bool

func startNewSudokuGame(b *button) {
	switch b.Text.text {
	case "Easy":
		grid = sudoku.CreateGrid(0)
		isEasy = true
	case "Medium":
		grid = sudoku.CreateGrid(1)
		isEasy = false
	case "Hard":
		grid = sudoku.CreateGrid(2)
		isEasy = false
	}

	for i := 0; i < rowSize; i++ {
		originalGrid[i] = make([]int, len(grid[i]))
		copy(originalGrid[i], grid[i])
	}

	selectedCell = matrixCell{B: nil, Row: 0, Col: 0}
	won = false

	createEverything(false)
}

func renderSudokuScreen() {
	for {
		event := sdl.WaitEvent()
		switch event.(type) {
		case *sdl.MouseButtonEvent:
			processButtonEvents(event)
		case *sdl.QuitEvent:
			saveGame()
			closeGame()
			return
		}

		renderer.SetDrawColor(46, 42, 56, 255)
		renderer.Clear()

		drawMatrix()
		drawNumbInput()
		backButton.drawButton()
		clearButton.drawButton()
		clearAllButton.drawButton()

		renderer.Present()
	}
}

//----------------------------//
// Create UI elements
//----------------------------//

func createEverything(cont bool) {
	createMatrix()
	createBoundaries()
	createNumInput()
	createBackButton()
	createClearButtons()

	if cont {
		backButton.Fn = func(b *button) { InitialScreen() }
	} else {
		backButton.Fn = func(b *button) { saveGame(); chooseDifficulty(nil) }
	}

	renderSudokuScreen()
}

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

//----------------------------//
// Interaction methods
//----------------------------//

func changeColor(b *button) {
	// Check if color is the original one (grey)
	if b.Color.r == 214 {
		b.Color = struct {
			r uint8
			g uint8
			b uint8
			a uint8
		}{r: 240, g: 228, b: 81, a: 255}

		if selectedCell.B != nil {
			changeColor(selectedCell.B)
		}

		selectedCell.B = b

		if isEasy {
			setAvailableOpt(sudoku.GetImpossibleNum(grid, selectedCell.Row, selectedCell.Col))
		}
	} else {
		b.Color = struct {
			r uint8
			g uint8
			b uint8
			a uint8
		}{r: 214, g: 214, b: 214, a: 255}

		selectedCell.B = nil
	}
}

func changeCellNum(b *button) {
	num, _ := strconv.Atoi(b.Text.text)
	if selectedCell.B != nil && sudoku.CheckValue(grid, selectedCell.Row, selectedCell.Col, num) && isEasy || selectedCell.B != nil && !isEasy {
		selectedCell.B.Text = b.Text
		_, selectedCell.B.Text.textTex, _ = createText(b.Text.text, 24, 110, 110, 110, 255)
		changeColor(selectedCell.B)
		grid[selectedCell.Row][selectedCell.Col] = num
	}
}

func setAvailableOpt(opt []int) {
	resetOpt()
	for i := range opt {
		buttons[opt[i]-1].Color = struct {
			r uint8
			g uint8
			b uint8
			a uint8
		}{r: 125, g: 125, b: 125, a: 255}
	}
}

func resetOpt() {
	for i := range buttons {
		buttons[i].Color = struct {
			r uint8
			g uint8
			b uint8
			a uint8
		}{r: 214, g: 214, b: 214, a: 255}
	}
}

func clearNum(b *button) {
	if selectedCell.B != nil {
		selectedCell.B.Text.text = " "
		_, selectedCell.B.Text.textTex, _ = createText(" ", 24, 0, 0, 0, 0)
		grid[selectedCell.Row][selectedCell.Col] = 0
		changeColor(selectedCell.B)
	}
}

func clearGame(b *button) {
	for i := 0; i < rowSize; i++ {
		grid[i] = make([]int, len(originalGrid[i]))
		copy(grid[i], originalGrid[i])
	}

	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			selectedCell = matrixCells[row][col]
			if !selectedCell.Initial {
				clearNum(nil)
			}
		}
	}
}

//----------------------------//
// Process button events
//----------------------------//

func processButtonEvents(event sdl.Event) {
	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			if !matrixCells[row][col].Initial && matrixCells[row][col].B.processEvent(event) {
				selectedCell.Row = row
				selectedCell.Col = col
				break
			}

			if sudoku.CheckZeroes(grid) == 0 && sudoku.CheckSolution(grid) && !won {
				win()
			}
		}
	}

	for i := range buttons {
		buttons[i].processEvent(event)
	}

	backButton.processEvent(event)
	clearButton.processEvent(event)
	clearAllButton.processEvent(event)
}

//----------------------------//
// Save and Load game methods
//----------------------------//

func saveGame() {
	gridSolution := sudoku.GetSolution()

	saveGame := saveStatus{Sol: gridSolution, Cur: grid, Org: originalGrid, IsEasy: isEasy}

	json, err := json.MarshalIndent(saveGame, "", " ")
	if err != nil {
		fmt.Errorf("error converting game state to json: ", err)
	}

	ioutil.WriteFile("savegame.json", json, 0644)
}

func loadGame(b *button) {
	file, err := ioutil.ReadFile("savegame.json")
	if err != nil {
		fmt.Errorf("Error loading savegame: ", err)
		noSavegame(b)
		return
	}

	data := saveStatus{}
	_ = json.Unmarshal([]byte(file), &data)

	sudoku.ReadSudoku(data.Sol)
	grid = data.Cur
	originalGrid = data.Org
	isEasy = data.IsEasy

	createEverything(true)
}

//----------------------------//
// Winning screen
//----------------------------//

func win() {
	message := [][]string{{"Y", "O", "U"}, {"W", "O", "N"}, {" ", "!", " "}}

	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			matrixCells[row][col].B.Text.text = " "
			_, matrixCells[row][col].B.Text.textTex, _ = createText(" ", 24, 0, 0, 0, 0)
		}
	}

	for row := 3; row < 6; row++ {
		for col := 3; col < 6; col++ {
			*matrixCells[row][col].B = createButton(
				&sdl.Rect{X: matrixCells[row][col].B.Rect.X, Y: matrixCells[row][col].B.Rect.Y, W: 50, H: 50},
				&sdl.Color{R: 214, G: 214, B: 214, A: 255},
				&sdl.Color{R: 0, G: 0, B: 0, A: 0},
				message[row-3][col-3],
				24,
				nil)
		}
	}

	won = true
}
