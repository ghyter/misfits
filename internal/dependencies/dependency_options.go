package dependencies

import (
	"github.com/ghyter/misfits/internal/embeds"
	"github.com/ghyter/misfits/internal/game"
	"github.com/ghyter/misfits/internal/resources"
)

type Options struct {
	AssetManager    embeds.AssetManager
	FontManager     resources.FontManager
	GameManager     game.GameManager
	DefaultFont     string  // Font name
	DefaultFontSize float64 // Font size
}

type Option func(*Options)

func WithAssetManager(assetManager embeds.AssetManager) Option {
	return func(o *Options) {
		o.AssetManager = assetManager
	}
}

func WithFontManager(fontManager resources.FontManager) Option {
	return func(o *Options) {
		o.FontManager = fontManager
	}
}

func WithGameManager(gameManager game.GameManager) Option {
	return func(o *Options) {
		o.GameManager = gameManager
	}
}

func WithDefaultFont(fontName string) Option {
	return func(o *Options) {
		o.DefaultFont = fontName
	}
}

func WithDefaultFontSize(fontSize float64) Option {
	return func(o *Options) {
		o.DefaultFontSize = fontSize
	}
}
