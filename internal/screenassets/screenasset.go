package screenassets

import "github.com/hajimehoshi/ebiten/v2"

type ScreenAsset interface {
	Update() error
	Draw(screen *ebiten.Image)
}
