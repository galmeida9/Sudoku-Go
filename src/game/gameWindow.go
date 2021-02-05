package game

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

var window *sdl.Window
var renderer *sdl.Renderer

var backButton button

// CreateGameWindow creates the window for the game
func CreateGameWindow() error {
	if err := initializeDependencies(); err != nil {
		return fmt.Errorf("Error initializing Dependencies: %q", err)
	}

	var err error
	window, err = sdl.CreateWindow(
		"Sudoku Go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return fmt.Errorf("Error initializing window: %q", err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return fmt.Errorf("Error initializing renderer: %q", err)
	}

	if err = createWindowIcon(); err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func initializeDependencies() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("Error initializing SDL: %q", err)
	}

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Error initializing TTF: %q", err)
	}

	if err := img.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("Error initializing SDL: %q", err)
	}

	return nil
}

func createText(text string, size int, r, g, b, a uint8) (*ttf.Font, *sdl.Texture, error) {
	textFont, err := ttf.OpenFont("resources/fonts/YuseiMagic-Regular.ttf", size)
	if err != nil {
		return nil, nil, fmt.Errorf("Error initializing game font: %q", err)
	}

	titleSurface, err := textFont.RenderUTF8Blended(text, sdl.Color{R: r, G: g, B: b, A: a})
	if err != nil {
		return nil, nil, fmt.Errorf("Error initializing sudoku title logo surface: %q", err)
	}

	textTex, err := renderer.CreateTextureFromSurface(titleSurface)
	if err != nil {
		return nil, nil, fmt.Errorf("Error initializing sudoku title logo texture: %q", err)
	}

	return textFont, textTex, nil
}

func createWindowIcon() error {
	img, err := img.Load("resources/img/sudoku.png")
	if err != nil {
		return fmt.Errorf("Error setting window icon: %q", err)
	}
	defer img.Free()

	window.SetIcon(img)
	return nil
}

func closeGame() {
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
	os.Exit(0)
}

func createBackButton() {
	backButton = createButton(
		&sdl.Rect{X: 475, Y: 15, W: 100, H: 40},
		&sdl.Color{R: 125, G: 125, B: 125, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		"Back",
		24,
		nil)
}
