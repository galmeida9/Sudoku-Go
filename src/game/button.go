package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type funcDef func(b *button)

type button struct {
	Rect  *sdl.Rect
	Color struct {
		r, g, b, a uint8
	}
	Pressed bool
	Text    struct {
		w, h    int
		textTex *sdl.Texture
		text    string
	}
	Fn funcDef
}

func createButton(rect *sdl.Rect, color, textColor *sdl.Color, text string, textSize int, fn funcDef) button {
	textFont, textTex, _ := createText(text, textSize, textColor.R, textColor.G, textColor.B, textColor.A)
	textW, textH, _ := textFont.SizeUTF8(text)

	return button{Rect: rect, Color: struct {
		r uint8
		g uint8
		b uint8
		a uint8
	}{r: color.R, g: color.G, b: color.B, a: color.A}, Pressed: false, Text: struct {
		w       int
		h       int
		textTex *sdl.Texture
		text    string
	}{w: textW, h: textH, textTex: textTex, text: text}, Fn: fn}
}

func (b *button) drawButton() bool {
	renderer.SetDrawColor(b.Color.r, b.Color.g, b.Color.b, b.Color.a)
	renderer.FillRect(b.Rect)
	renderer.Copy(b.Text.textTex, nil, &sdl.Rect{
		X: b.Rect.X + (b.Rect.W-int32(b.Text.w))/2,
		Y: b.Rect.Y + (b.Rect.H-int32(b.Text.h))/2,
		W: int32(b.Text.w),
		H: int32(b.Text.h)})

	// if button press detected - reset it so it wouldn't trigger twice
	if b.Pressed {
		b.Pressed = false
		if b.Fn != nil {
			b.Fn(b)
		}
		return true
	}

	return false
}

func (b *button) processEvent(event sdl.Event) bool {
	x, y, state := sdl.GetMouseState()

	if state == sdl.Button(sdl.BUTTON_LEFT) && x >= b.Rect.X &&
		x <= (b.Rect.X+b.Rect.W) && y >= b.Rect.Y &&
		y <= (b.Rect.Y+b.Rect.H) {
		b.Pressed = true
		return true
	}

	return false
}
