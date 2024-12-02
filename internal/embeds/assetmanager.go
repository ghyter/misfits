package embeds

import (
	"bytes"
	"embed"
	"fmt"
	"image"

	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/hajimehoshi/ebiten/v2"
	_ "golang.org/x/image/webp"
)

//go:embed assets
var embeddedAssets embed.FS

type AssetManager interface {
	Get(path string) ([]byte, error)
	GetEbitenImage(path string) (*ebiten.Image, error)
}

type DefaultAssetManager struct {
	cache map[string][]byte
	dm    *dependencies.DependencyManager
}

func NewDefaultAssetManager(dm *dependencies.DependencyManager) (AssetManager, error) {

	return &DefaultAssetManager{cache: make(map[string][]byte),
		dm: dm,
	}, nil
}

func (a *DefaultAssetManager) Get(path string) ([]byte, error) {
	if data, ok := a.cache[path]; ok {
		return data, nil
	}
	data, err := embeddedAssets.ReadFile("assets/" + path)
	if err != nil {
		return nil, fmt.Errorf("failed to load asset %s: %w", path, err)
	}
	a.cache[path] = data
	return data, nil

}

func (a *DefaultAssetManager) GetEbitenImage(path string) (*ebiten.Image, error) {

	rawimage, err := a.Get(path)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(rawimage)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil

}
