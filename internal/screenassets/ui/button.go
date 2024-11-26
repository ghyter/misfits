package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// InputHandler abstracts input handling for testing.
type InputHandler interface {
	CursorPosition() (int, int)
	IsMouseButtonPressed(button ebiten.MouseButton) bool
}

// DefaultInputHandler is the default implementation using Ebiten.
type DefaultInputHandler struct{}

func (d *DefaultInputHandler) CursorPosition() (int, int) {
	return ebiten.CursorPosition()
}

func (d *DefaultInputHandler) IsMouseButtonPressed(button ebiten.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(button)
}

// Button represents a simple UI button.
type Button struct {
	X, Y          int
	Width, Height int
	Label         string
	OnClick       func()
	isHovered     bool
	isClicked     bool
	Input         InputHandler
}

func NewButton(x, y, width, height int, label string, onClick func()) *Button {
	return &Button{
		X:       x,
		Y:       y,
		Width:   width,
		Height:  height,
		Label:   label,
		OnClick: onClick,
		Input:   &DefaultInputHandler{},
	}
}

func (b *Button) Update() {
	cursorX, cursorY := b.Input.CursorPosition()
	b.isHovered = cursorX >= b.X && cursorX <= b.X+b.Width && cursorY >= b.Y && cursorY <= b.Y+b.Height

	if b.isHovered && b.Input.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		b.isClicked = true
	}
	if b.isClicked && !b.Input.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		b.isClicked = false
		if b.isHovered && b.OnClick != nil {
			b.OnClick()
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	bgColor := color.RGBA{200, 200, 200, 255}
	if b.isHovered {
		bgColor = color.RGBA{150, 150, 150, 255}
	}
	if b.isClicked {
		bgColor = color.RGBA{100, 100, 100, 255}
	}
	rect := ebiten.NewImage(b.Width, b.Height)
	rect.Fill(bgColor)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(rect, op)
}
