package game

import (
	"Sudoku-Go/src/sudoku"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func startNewSudokuGame() {
	a := sudoku.CreateGrid()
	fmt.Println("done")
	fmt.Println(a)
	renderSudokuScreen()
}

func renderSudokuScreen() {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				closeGame()
				return
			}
		}

		renderer.SetDrawColor(46, 42, 56, 255)
		renderer.Clear()

		renderer.Present()
	}
}
