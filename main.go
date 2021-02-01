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

	// a := [1]int{2}
	// test(&a)
	// fmt.Println(a)
}

// func test(a *[1]int) {
// 	a[0] = 3
// }
