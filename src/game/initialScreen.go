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
	startButton       Button
	continueButton    Button
}

func InitialScreen(gameObj Game) {
	screen := initialScreen{}

	logoTex, err := loadLogoImg(gameObj.Renderer)
	if err != nil {
		fmt.Println("Error initializing sudoku logo", err)
		return
	}
	screen.logoTex = logoTex
	defer logoTex.Destroy()

	titleFont, titleTex, err := createText(gameObj.Renderer, "Sudoku Go", 64, 255, 255, 255, 255)
	if err != nil {
		fmt.Println("Error initializing sudoku logo text", err)
		return
	}
	screen.titleTex = titleTex
	screen.titleFont = titleFont
	defer titleTex.Destroy()
	defer titleFont.Close()

	createButtons(gameObj.Renderer, &screen)

	renderInitialScreen(gameObj, screen)
}

func renderInitialScreen(gameObj Game, screen initialScreen) {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.MouseButtonEvent:
				screen.startButton.processEvent(event)
				screen.continueButton.processEvent(event)
			case *sdl.QuitEvent:
				return
			}
		}

		gameObj.Renderer.SetDrawColor(46, 42, 56, 255)
		gameObj.Renderer.Clear()

		gameObj.Renderer.Copy(screen.logoTex,
			&sdl.Rect{X: 0, Y: 0, W: 512, H: 512},
			&sdl.Rect{X: 70, Y: 100, W: 128, H: 128})

		textW, textH, _ := screen.titleFont.SizeUTF8("Sudoku Go")
		gameObj.Renderer.Copy(screen.titleTex, nil, &sdl.Rect{X: 210, Y: 100, W: int32(textW), H: int32(textH)})

		screen.startButton.drawButton(gameObj.Renderer)
		screen.continueButton.drawButton(gameObj.Renderer)

		gameObj.Renderer.Present()
	}
}

func loadLogoImg(renderer *sdl.Renderer) (*sdl.Texture, error) {
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

func createButtons(renderer *sdl.Renderer, screen *initialScreen) {
	screen.startButton = createButton(
		renderer,
		&sdl.Rect{X: 75, Y: 400, W: 450, H: 105},
		&sdl.Color{R: 240, G: 228, B: 81, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		"Start Game",
		56)

	screen.continueButton = createButton(
		renderer,
		&sdl.Rect{X: 75, Y: 550, W: 450, H: 105},
		&sdl.Color{R: 85, G: 217, B: 102, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		"Continue Game",
		56)
}
