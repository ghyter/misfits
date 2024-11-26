package game

import (
	"errors"

	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameManagerOption func(*GameManagerOptions)

type GameManagerOptions struct {
	TextWriter  func(dst *ebiten.Image, text string, x, y int)
	FontManager resources.FontManager
	dm          *dependencies.DependencyManager
}

func NewGameManagerOptions(dm *dependencies.DependencyManager, opts ...GameManagerOption) (*GameManagerOptions, error) {
	options := &GameManagerOptions{
		dm: dm,
	}
	for _, opt := range opts {
		opt(options)
	}

	if options.FontManager == nil {
		fm, err := dependencies.Get[resources.FontManager](dm)
		if err != nil {
			return nil, errors.New("FontManager is Required")
		}
		options.FontManager = fm
	}

	return options, nil
}

func WithTextWriter(tw func(dst *ebiten.Image, text string, x, y int)) GameManagerOption {
	return func(o *GameManagerOptions) {
		o.TextWriter = tw
	}
}

func WithFontManager(fm resources.FontManager) GameManagerOption {
	return func(o *GameManagerOptions) {
		o.FontManager = fm
	}
}
