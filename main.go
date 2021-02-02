package main

import (
	"Sudoku-Go/src/game"
	"fmt"
)

func main() {
	if err := game.CreateGameWindow(); err != nil {
		fmt.Println("Error initializing window: ", err)
	}

	game.InitialScreen()

}
