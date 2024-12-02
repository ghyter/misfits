package states

import (
	"fmt"
	"image"
	"image/color"

	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/embeds"
	"github.com/ghyter/misfits/internal/globals"
	"github.com/ghyter/misfits/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	_ "golang.org/x/image/webp"
)

type MenuState struct {
	options         []int
	selected        int
	dm              *dependencies.DependencyManager
	backgroundImage *ebiten.Image
	menuFont        font.Face
}

// NewMenuState initializes a new MenuState.
func NewMenuState(dm *dependencies.DependencyManager) *MenuState {

	assets, err := dependencies.Get[embeds.AssetManager](dm)
	if err != nil {
		panic("Assets Manager not found")
	}
	backgroundImage, err := assets.GetEbitenImage("images/titlecard.webp")
	if err != nil {
		panic("Background Image Not Found")
	}

	fontManager, _ := dependencies.Get[resources.FontManager](dm)
	menufont, err := fontManager.LoadFont(resources.DefaultFontName, 32)
	if err != nil {
		panic("Default Font not found")
	}

	return &MenuState{
		options:         []int{2, 3, 4}, // Number of players to choose
		selected:        0,              // Initially, the first option is selected
		dm:              dm,
		backgroundImage: backgroundImage,
		menuFont:        menufont,
	}
}

// Update handles input and state updates for the menu.
func (m *MenuState) Update() (int, bool) {
	// Handle navigation with arrow keys
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.selected--
		if m.selected < 0 {
			m.selected = len(m.options) - 1
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.selected++
		if m.selected >= len(m.options) {
			m.selected = 0
		}
	}

	// Confirm selection with Enter
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return m.options[m.selected], true // Return selected option and signal transition
	}

	return 0, false // No transition
}

// Draw renders the menu to the screen.
func (m *MenuState) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 255}) // Clear the screen with black

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(.45, .45)
	screen.DrawImage(m.backgroundImage, op)

	for i, option := range m.options {
		text := fmt.Sprintf("%d Player Game", option)

		x := 100
		y := 500 + i*50 // Vertical spacing between options

		// Highlight the selected option
		textColor := globals.InactiveText // White for non-selected
		if i == m.selected {
			textColor = globals.ActiveText
		}

		drawer := &font.Drawer{
			Dst:  screen,
			Src:  image.NewUniform(textColor), // Color for the text
			Face: m.menuFont,
			Dot:  fixed.P(x, y),
		}

		// Draw the text using the selected color
		drawer.DrawString(text)
	}
}
