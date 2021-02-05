package game

import "github.com/veandco/go-sdl2/sdl"

var easy, medium, hard button
var isEasy bool

func chooseDifficulty(b *button) {
	easy = createDifficultyOpt("Easy", 75, 200)
	medium = createDifficultyOpt("Medium", 75, 350)
	hard = createDifficultyOpt("Hard", 75, 500)
	createBackButton()
	backButton.Fn = func(b *button) { InitialScreen() }

	renderDifficultyScreen()
}

func createDifficultyOpt(text string, x, y int32) button {
	return createButton(
		&sdl.Rect{X: x, Y: y, W: 450, H: 105},
		&sdl.Color{R: 214, G: 214, B: 214, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		text,
		56,
		startNewSudokuGame)
}

func renderDifficultyScreen() {
	for {
		event := sdl.WaitEvent()
		switch event.(type) {
		case *sdl.MouseButtonEvent:
			easy.processEvent(event)
			medium.processEvent(event)
			hard.processEvent(event)
			backButton.processEvent(event)
		case *sdl.QuitEvent:
			closeGame()
			return
		}

		renderer.SetDrawColor(46, 42, 56, 255)
		renderer.Clear()

		easy.drawButton()
		medium.drawButton()
		hard.drawButton()
		backButton.drawButton()

		renderer.Present()
	}
}
