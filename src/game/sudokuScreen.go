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

func changeColor(b *button) {
	showHighlight(grid[selectedCell.Row][selectedCell.Col])

	if selectedCell.Initial {
		// if an initial puzzle cell is selected then no input is valid
		setAvailableOpt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	} else if isEasy {
		// if it is on easy mode to help we show the only valid options
		setAvailableOpt(sudoku.GetImpossibleNum(grid, selectedCell.Row, selectedCell.Col))
	} else {
		setAvailableOpt([]int{})
	}
}

func showHighlight(value int) {
	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			if grid[row][col] == value && row != selectedCell.Row && col != selectedCell.Col && value != 0 {
				matrixCells[row][col].B.Color = struct {
					r uint8
					g uint8
					b uint8
					a uint8
				}{r: 224, g: 172, b: 27, a: 255}
			} else if row == selectedCell.Row && col == selectedCell.Col {
				matrixCells[row][col].B.Color = struct {
					r uint8
					g uint8
					b uint8
					a uint8
				}{r: 240, g: 228, b: 81, a: 255}
			} else {
				matrixCells[row][col].B.Color = struct {
					r uint8
					g uint8
					b uint8
					a uint8
				}{r: 214, g: 214, b: 214, a: 255}
			}
		}
	}
}

func changeCellNum(b *button) {
	num, _ := strconv.Atoi(b.Text.text)
	if !selectedCell.Initial && selectedCell.B != nil && sudoku.CheckValue(grid, selectedCell.Row, selectedCell.Col, num) && isEasy ||
		!selectedCell.Initial && selectedCell.B != nil && !isEasy {
		selectedCell.B.Text = b.Text
		_, selectedCell.B.Text.textTex, _ = createText(b.Text.text, 24, 110, 110, 110, 255)
		grid[selectedCell.Row][selectedCell.Col] = num
		showHighlight(num)
	}
}

// setAvailableOpt shows the valid options given and array of the invalid ones
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
		showHighlight(0)
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
			if matrixCells[row][col].B.processEvent(event) {
				selectedCell.B = matrixCells[row][col].B
				selectedCell.Initial = matrixCells[row][col].Initial
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
		fmt.Println("error converting game state to json: ", err)
	}

	ioutil.WriteFile("savegame.json", json, 0644)
}

func loadGame(b *button) {
	file, err := ioutil.ReadFile("savegame.json")
	if err != nil {
		fmt.Println("Error loading savegame: ", err)
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
