package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func CreateGameWindow() (*sdl.Window, *sdl.Renderer, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, nil, fmt.Errorf("Error initializing SDL: ", err)
	}

	window, err := sdl.CreateWindow(
		"Sudoku Go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return nil, nil, fmt.Errorf("Error initializing window: ", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, nil, fmt.Errorf("Error initializing renderer: ", err)
	}

	return window, renderer, nil
}

func InitialScreen(renderer *sdl.Renderer) {
	img, err := sdl.LoadBMP("resources/img/sudoku.bmp")
	if err != nil {
		fmt.Println("Error initializing sudoku logo image: ", err)
		return
	}
	defer img.Free()

	logoTex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("Error initializing sudoku logo texture: ", err)
		return
	}
	defer logoTex.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.Copy(logoTex,
			&sdl.Rect{X: 0, Y: 0, W: 512, H: 512},
			&sdl.Rect{X: 40, Y: 20, W: 128, H: 128})

		renderer.Present()
	}
}
