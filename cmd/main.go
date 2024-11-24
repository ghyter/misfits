package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct{}

func (g *Game) Update() error {
	fmt.Println("Update loop")
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen with a background color
	screen.Fill(color.RGBA{R: 50, G: 50, B: 150, A: 255})

	// Draw a simple rectangle
	rect := ebiten.NewImage(100, 50)
	rect.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(350, 275)
	screen.DrawImage(rect, op)

	// Display "Hello, Ebiten!" text
	text.Draw(screen, "Hello, Ebiten!", basicfont.Face7x13, 350, 250, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	game := &Game{}
	ebiten.RunGame(game)
}
