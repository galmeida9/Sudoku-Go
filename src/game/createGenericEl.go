package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// createText creates text
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

// createBackButton creates a generic back button
func createBackButton() {
	backButton = createButton(
		&sdl.Rect{X: 475, Y: 15, W: 100, H: 40},
		&sdl.Color{R: 125, G: 125, B: 125, A: 255},
		&sdl.Color{R: 0, G: 0, B: 0, A: 0},
		"Back",
		24,
		nil)
}
