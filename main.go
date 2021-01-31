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
	gameObj, err := game.CreateGameWindow()

	if err != nil {
		fmt.Println("Error initializing window: ", err)
	}

	defer gameObj.Window.Destroy()
	defer gameObj.Renderer.Destroy()

	game.InitialScreen(gameObj)
}
