package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/screenassets/states"
)

type Game interface {
	ebiten.Game
}

type MisfitGame struct {
	dm        *dependencies.DependencyManager
	gmOptions *GameManagerOptions
	options   *GameOptions
	state     states.GameState
	menu      *states.MenuState
}

func (g *MisfitGame) InitUI() error {

	g.menu = states.NewMenuState(g.dm)

	return nil
}

// Update implements Game.
func (g *MisfitGame) Update() error {

	switch g.state {
	case states.Menu:
		newstate, changed := g.menu.Update()

		if changed {
			g.state = states.GameState(newstate)
		}
	case states.Game:
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.state = states.Pause // Pause the game
		}
	case states.Pause:
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.state = states.Game // Resume gameplay
		}
	case states.GameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.state = states.Menu // Back to the menu
		}
	}
	return nil
}

func (g *MisfitGame) Draw(screen *ebiten.Image) {
	// Clear the screen with a background color
	screen.Fill(color.RGBA{R: 50, G: 50, B: 150, A: 255})
	switch g.state {
	case states.Menu:
		g.menu.Draw(screen)
	case states.Game:
		g.drawGame(screen)
	case states.Pause:
		g.drawPause(screen)
	case states.GameOver:
		g.drawGameOver(screen)
	}
}

func (g *MisfitGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.options.ScreenWidth, g.options.ScreenHeight
}

// drawGame renders the gameplay.
func (g *MisfitGame) drawGame(screen *ebiten.Image) {
	// Add your rendering logic here
}

// drawPause renders the pause screen.
func (g *MisfitGame) drawPause(screen *ebiten.Image) {
	// Add your rendering logic here
}

// drawGameOver renders the game over screen.
func (g *MisfitGame) drawGameOver(screen *ebiten.Image) {
	// Add your rendering logic here
}
