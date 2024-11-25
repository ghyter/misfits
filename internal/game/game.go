package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameManager interface {
	NewGame() Game
}

type DefaultGameManager struct {
	drawText func(dst *ebiten.Image, text string, x, y int)
}

func NewDefaultGameManager(textWriter func(dst *ebiten.Image, text string, x, y int)) (GameManager, error) {

	return &DefaultGameManager{
		drawText: textWriter,
	}, nil
}

func (gm *DefaultGameManager) NewGame() Game {

	return &MisfitGame{
		drawText: gm.drawText,
	}
}

type Game interface {
	ebiten.Game
}

type MisfitGame struct {
	drawText func(dst *ebiten.Image, text string, x, y int)
}

// Update implements Game.
func (g *MisfitGame) Update() error {
	fmt.Println("Update loop")
	return nil
}

func (g *MisfitGame) Draw(screen *ebiten.Image) {
	// Clear the screen with a background color
	screen.Fill(color.RGBA{R: 50, G: 50, B: 150, A: 255})

	// Draw a simple rectangle
	rect := ebiten.NewImage(100, 50)
	rect.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(350, 275)
	screen.DrawImage(rect, op)

	// Display "Hello, Ebiten!" text

	g.drawText(screen, "Hello, Ebiten! With the new Font", 25, 25)
}

func (g *MisfitGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
