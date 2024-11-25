package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game interface {
	ebiten.Game
}

type MisfitGame struct {
	drawText func(dst *ebiten.Image, text string, x, y int)
	options  *GameOptions
}

// Update implements Game.
func (g *MisfitGame) Update() error {
	fmt.Printf("Update loop for %d player(s)\n", g.options.NumPlayers)
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
	g.drawText(screen, "Hello, Ebiten!", 25, 25)
}

func (g *MisfitGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.options.ScreenWidth, g.options.ScreenHeight
}
