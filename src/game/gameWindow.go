package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

type Game struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
}

func CreateGameWindow() (Game, error) {
	if err := initializeDependencies(); err != nil {
		return Game{}, fmt.Errorf("Error initializing Dependencies: ", err)
	}

	window, err := sdl.CreateWindow(
		"Sudoku Go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return Game{}, fmt.Errorf("Error initializing window: ", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return Game{}, fmt.Errorf("Error initializing renderer: ", err)
	}

	return Game{Window: window, Renderer: renderer}, nil
}

func initializeDependencies() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("Error initializing SDL: ", err)
	}

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Error initializing TTF: ", err)
	}

	if err := img.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("Error initializing SDL: ", err)
	}

	return nil
}
