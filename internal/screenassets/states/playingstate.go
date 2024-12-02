package states

import (
	"fmt"
	"image"

	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/globals"
	"github.com/ghyter/misfits/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type PlayingState struct {
	numplayers int
	playerfont font.Face
}

func NewPlayingState(dm *dependencies.DependencyManager, players int) *PlayingState {

	fontManager, _ := dependencies.Get[resources.FontManager](dm)
	playerfont, err := fontManager.LoadFont(resources.DefaultFontName, 16)
	if err != nil {
		panic("Default Font not found")
	}

	return &PlayingState{
		playerfont: playerfont,
		numplayers: players,
	}
}

func (p *PlayingState) Update() bool {

	return false
}

func (p *PlayingState) Draw(screen *ebiten.Image) {
	// Highlight the selected option
	textColor := globals.InactiveText // White for non-selected
	drawer := &font.Drawer{
		Dst:  screen,
		Src:  image.NewUniform(textColor), // Color for the text
		Face: p.playerfont,
		Dot:  fixed.P(10, 20),
	}

	drawer.DrawString(fmt.Sprintf("%d Player Game", p.numplayers))
}
