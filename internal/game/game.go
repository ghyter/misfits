package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"golang.org/x/image/font"

	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/screenassets"
)

type Game interface {
	ebiten.Game
}

type GameState int

const (
	StateMenu GameState = iota // Menu state of the game
	StateGame                  // Gameplay state
)

type MisfitGame struct {
	dm        *dependencies.DependencyManager
	gmOptions *GameManagerOptions
	options   *GameOptions
	state     GameState
	font      font.Face
	menu      screenassets.MenuScreen
}

func (g *MisfitGame) InitUI() error {
	g.menu = *screenassets.NewMenuScreen()
	return nil
}

// Update implements Game.
func (g *MisfitGame) Update() error {

	switch g.state {
	case StateMenu:
		// Handle input to select player count
		if ebiten.IsKeyPressed(ebiten.Key2) {
			g.options.NumPlayers = 2
			g.state = StateGame
		} else if ebiten.IsKeyPressed(ebiten.Key3) {
			g.options.NumPlayers = 3
			g.state = StateGame
		} else if ebiten.IsKeyPressed(ebiten.Key4) {
			g.options.NumPlayers = 4
			g.state = StateGame
		}
	case StateGame:
		fmt.Printf("Game running with %d players\n", g.options.NumPlayers)
	}
	return nil
}

func (g *MisfitGame) Draw(screen *ebiten.Image) {
	// Clear the screen with a background color
	screen.Fill(color.RGBA{R: 50, G: 50, B: 150, A: 255})
	switch g.state {
	case StateMenu:
		g.menu.Draw(screen)
	case StateGame:

	}

}

func (g *MisfitGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.options.ScreenWidth, g.options.ScreenHeight
}
