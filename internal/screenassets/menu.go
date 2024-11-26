package screenassets

import (
	"fmt"

	"github.com/ghyter/misfits/internal/screenassets/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type MenuScreen struct {
	button *ui.Button
}

func NewMenuScreen() *MenuScreen {
	button := ui.NewButton(
		5, 150, 50, 80, "Start Game",
		func() {
			fmt.Println("Button clicked: Start Game!")
		},
	)
	return &MenuScreen{
		button: button,
	}
}

func (m *MenuScreen) Update() error {
	m.button.Update()
	return nil
}

func (m *MenuScreen) Draw(screen *ebiten.Image) {
	m.button.Draw(screen)
}

func (m *MenuScreen) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}
