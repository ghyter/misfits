package states

import (
	"fmt"
	"image"
	"image/color"

	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type MenuState struct {
	options  []int
	selected int
	dm       *dependencies.DependencyManager
}

// NewMenuState initializes a new MenuState.
func NewMenuState(dm *dependencies.DependencyManager) *MenuState {
	return &MenuState{
		options:  []int{2, 3, 4}, // Number of players to choose
		selected: 0,              // Initially, the first option is selected
		dm:       dm,
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

	for i, option := range m.options {
		text := fmt.Sprintf("%d Player Game", option)
		y := 100 + i*50 // Vertical spacing between options

		// Highlight the selected option
		textColor := color.RGBA{R: 255, G: 255, B: 255, A: 255} // White for non-selected
		if i == m.selected {
			textColor = color.RGBA{R: 255, G: 255, B: 0, A: 255} // Yellow for selected
		}

		defaultfont, err := dependencies.Get[font.Face](m.dm)
		if err != nil {
			panic("Default Font not found")
		}

		drawer := &font.Drawer{
			Dst:  screen,
			Src:  image.NewUniform(textColor), // Color for the text
			Face: defaultfont,
			Dot:  fixed.P(200, y),
		}

		// Draw the text using the selected color
		drawer.DrawString(text)
	}
}
