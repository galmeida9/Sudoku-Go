package main

import (
	"Sudoku-Go/src/game"
	"fmt"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	window, renderer, err := game.CreateGameWindow()

	if err != nil {
		fmt.Println("Error initializing window: ", err)
	}

	defer window.Destroy()
	defer renderer.Destroy()

	game.InitialScreen(renderer)
}
