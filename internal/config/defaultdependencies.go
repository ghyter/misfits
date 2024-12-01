package config

import (
	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/embeds"
	"github.com/ghyter/misfits/internal/game"
	"github.com/ghyter/misfits/internal/resources"
	"golang.org/x/image/font"
)

func NewDefaultDependencies() (*dependencies.DependencyManager, error) {
	dm := dependencies.NewDependencyManager()

	dependencies.Register(dm, func(dm *dependencies.DependencyManager) (embeds.AssetManager, error) {

		return embeds.NewDefaultAssetManager(dm)
	})

	dependencies.Register(dm, func(dm *dependencies.DependencyManager) (resources.FontManager, error) {

		return resources.NewDefaultFontManager(dm)
	})

	dependencies.Register(dm, func(dm *dependencies.DependencyManager) (game.GameManager, error) {

		return game.NewDefaultGameManager(dm)
	})

	dependencies.Register(dm, func(dm *dependencies.DependencyManager) (font.Face, error) {
		fontManager, err := dependencies.Get[resources.FontManager](dm)
		if err != nil {
			return nil, err
		}
		return fontManager.LoadFont(resources.DefaultFontName, 12)
	})

	return dm, nil
}
