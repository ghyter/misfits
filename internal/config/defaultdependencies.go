package config

import (
	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/embeds"
	"github.com/ghyter/misfits/internal/game"
	"github.com/ghyter/misfits/internal/resources"
)

func NewDefaultDependencies() (*dependencies.DependencyManager, error) {
	dm := dependencies.NewDependencyManager()

	dependencies.Register(dm, func() (embeds.AssetManager, error) {

		return embeds.NewDefaultAssetManager(dm)
	})

	dependencies.Register(dm, func() (embeds.AssetManager, error) {

		return embeds.NewDefaultAssetManager(dm)
	})

	dependencies.Register(dm, func() (resources.FontManager, error) {

		return resources.NewDefaultFontManager(dm)
	})

	dependencies.Register(dm, func() (game.GameManager, error) {

		return game.NewDefaultGameManager(dm)
	})

	return dm, nil
}
