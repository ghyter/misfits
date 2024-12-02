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
	playing   *states.PlayingState
	paused    *states.PausedState
}

func (g *MisfitGame) InitUI() error {

	g.menu = states.NewMenuState(g.dm)

	return nil
}

// Update implements Game.
func (g *MisfitGame) Update() error {

	switch g.state {
	case states.Menu:
		numbPlayers, changed := g.menu.Update()
		if changed {
			g.options.NumPlayers = numbPlayers
			g.state = states.Playing
			g.playing = states.NewPlayingState(g.dm, numbPlayers)

		}
	case states.Playing:
		g.playing.Update()
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.state = states.Pause // Pause the game
			g.paused = states.NewPausedState(g.dm)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			g.state = states.Menu // Pause the game
		}
	case states.Pause:
		g.paused.Update()
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.state = states.Playing // Resume gameplay
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
	case states.Playing:
		g.playing.Draw(screen)
	case states.Pause:
		g.paused.Draw(screen)
	case states.GameOver:
		screen.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255})
		g.drawGameOver(screen)
	}
}

func (g *MisfitGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.options.ScreenWidth, g.options.ScreenHeight
}

// drawGameOver renders the game over screen.
func (g *MisfitGame) drawGameOver(screen *ebiten.Image) {
	// Add your rendering logic here
}
