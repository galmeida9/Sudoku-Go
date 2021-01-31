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

	titleFont, titleTex, err := loadLogoText(gameObj)
	if err != nil {
		fmt.Println("Error initializing sudoku logo text", err)
		return
	}
	screen.titleTex = titleTex
	screen.titleFont = titleFont
	defer titleTex.Destroy()

	renderInitialScreen(gameObj, screen)
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

func loadLogoText(gameObj Game) (*ttf.Font, *sdl.Texture, error) {
	titleFont, err := ttf.OpenFont("resources/fonts/YuseiMagic-Regular.ttf", 64)
	if err != nil {
		return nil, nil, fmt.Errorf("Error initializing game font: ", err)
	}

	titleSurface, err := titleFont.RenderUTF8Blended("Sudoku Go", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return nil, nil, fmt.Errorf("Error initializing sudoku title logo surface: ", err)
	}

	titleTex, err := gameObj.Renderer.CreateTextureFromSurface(titleSurface)
	if err != nil {
		return nil, nil, fmt.Errorf("Error initializing sudoku title logo texture: ", err)
	}

	return titleFont, titleTex, nil
}

func renderInitialScreen(gameObj Game, screen initialScreen) {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
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

		gameObj.Renderer.Present()
	}
}
