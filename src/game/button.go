package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	Rect  *sdl.Rect
	Color struct {
		r, g, b, a uint8
	}
	Pressed bool
	Text    struct {
		w, h    int
		textTex *sdl.Texture
	}
}

func createButton(renderer *sdl.Renderer, rect *sdl.Rect, color, textColor *sdl.Color, text string, textSize int) Button {
	textFont, textTex, _ := createText(renderer, text, textSize, textColor.R, textColor.G, textColor.B, textColor.A)
	textW, textH, _ := textFont.SizeUTF8("Sudoku Go")

	return Button{Rect: rect, Color: struct {
		r uint8
		g uint8
		b uint8
		a uint8
	}{r: color.R, g: color.G, b: color.B, a: color.A}, Pressed: false, Text: struct {
		w       int
		h       int
		textTex *sdl.Texture
	}{w: textW, h: textH, textTex: textTex}}
}

func (b *Button) drawButton(renderer *sdl.Renderer) bool {
	renderer.SetDrawColor(b.Color.r, b.Color.g, b.Color.b, b.Color.a)
	renderer.FillRect(b.Rect)
	renderer.Copy(b.Text.textTex, nil, &sdl.Rect{
		X: b.Rect.X + (b.Rect.W-int32(b.Text.w))/2,
		Y: b.Rect.Y + (b.Rect.H-int32(b.Text.h))/2,
		W: int32(b.Text.w),
		H: int32(b.Text.h)})

	// if button press detected - reset it so it wouldn't trigger twice
	if b.Pressed {
		fmt.Println("pressed")
		b.Pressed = false
		return true
	}

	return false
}

func (b *Button) processEvent(event sdl.Event) {
	x, y, state := sdl.GetMouseState()

	if state == sdl.Button(sdl.BUTTON_LEFT) && x >= b.Rect.X &&
		x <= (b.Rect.X+b.Rect.W) && y >= b.Rect.Y &&
		y <= (b.Rect.Y+b.Rect.H) {
		b.Pressed = true
	}
}
