package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type initialScreen struct {
	logoTex, titleTex *sdl.Texture
	titleFont         *ttf.Font
	startButton       button
	continueButton    button
}

// InitialScreen creates the main screen
func InitialScreen() {
	screen := initialScreen{}

	logoTex, err := loadLogoImg()
	if err != nil {
		fmt.Println("Error initializing sudoku logo", err)
		return
	}
	screen.logoTex = logoTex
	defer logoTex.Destroy()

	titleFont, titleTex, err := createText("Sudoku Go", 64, 255, 255, 255, 255)
	if err != nil {
		fmt.Println("Error initializing sudoku logo text", err)
		return
	}
	screen.titleTex = titleTex
	screen.titleFont = titleFont
	defer titleTex.Destroy()
	defer titleFont.Close()

	createButtons(&screen)

	renderInitialScreen(screen)
}

func renderInitialScreen(screen initialScreen) {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.MouseButtonEvent:
				if screen.startButton.processEvent(event) {
					startNewSudokuGame()
				}
				screen.continueButton.processEvent(event)
			case *sdl.QuitEvent:
				closeGame()
				return
			}
		}

		renderer.SetDrawColor(46, 42, 56, 255)
		renderer.Clear()

		renderer.Copy(screen.logoTex,
			&sdl.Rect{X: 0, Y: 0, W: 512, H: 512},
			&sdl.Rect{X: 70, Y: 100, W: 128, H: 128})

		textW, textH, _ := screen.titleFont.SizeUTF8("Sudoku Go")
		renderer.Copy(screen.titleTex, nil, &sdl.Rect{X: 210, Y: 100, W: int32(textW), H: int32(textH)})

		screen.startButton.drawButton()
		screen.continueButton.drawButton()

		renderer.Present()
	}
}

func loadLogoImg() (*sdl.Texture, error) {
	img, err := img.Load("resources/img/sudoku.png")
	if err != nil {
		return nil, fmt.Errorf("Error initializing sudoku logo image: ", err)
	}
	defer img.Free()

	logoTex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("Error initializing sudoku logo texture: ", err)
	}

	return logoTex, nil
}

func createButtons(screen *initialScreen) {
	screen.startButton = createButton(
		&sdl.Rect{X: 75, Y: 400, W: 450, H: 105},
		&sdl.Color{R: 240, G: 228, B: 81, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		"Start Game",
		56)

	screen.continueButton = createButton(
		&sdl.Rect{X: 75, Y: 550, W: 450, H: 105},
		&sdl.Color{R: 85, G: 217, B: 102, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		"Continue Game",
		56)
}
