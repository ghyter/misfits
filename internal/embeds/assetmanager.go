package embeds

import (
	"embed"
	"fmt"
)

//go:embed assets
var embeddedAssets embed.FS

type AssetManager interface {
	Get(path string) ([]byte, error)
}

type DefaultAssetManager struct {
	cache map[string][]byte
}

func NewDefaultAssetManager() AssetManager {
	return &DefaultAssetManager{cache: make(map[string][]byte)}
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
