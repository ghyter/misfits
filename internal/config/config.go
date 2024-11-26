package config

import (
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

// EbitenConfig holds all the configuration for the Ebiten game window.
type EbitenConfig struct {
	WindowWidth        int
	WindowHeight       int
	WindowTitle        string
	WindowResizingMode ebiten.WindowResizingModeType
	UI                 *ebitenui.UI
}

// GetEbitenConfig initializes the default Ebiten configuration.
func GetEbitenConfig() *EbitenConfig {
	return &EbitenConfig{
		WindowWidth:        900,
		WindowHeight:       800,
		WindowTitle:        "Ebiten Misfit Friends",
		WindowResizingMode: ebiten.WindowResizingModeEnabled,
	}
}

// Apply applies the Ebiten configuration to the game engine.
func (c *EbitenConfig) Apply() {
	ebiten.SetWindowSize(c.WindowWidth, c.WindowHeight)
	ebiten.SetWindowTitle(c.WindowTitle)
	ebiten.SetWindowResizingMode(c.WindowResizingMode)

}
