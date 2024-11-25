package dependencies

import (
	"image"
	"image/color"

	"github.com/ghyter/misfits/internal/embeds"
	"github.com/ghyter/misfits/internal/game"
	"github.com/ghyter/misfits/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// DependencyManager interface provides access to all dependencies in the system.
type DependencyManager interface {
	GetAssetManager() embeds.AssetManager
	GetFontManager() resources.FontManager
	GetGameManager() game.GameManager
}

// DefaultDependencyManager is the concrete implementation of Dependencies.
type DefaultDependencyManager struct {
	AssetsManager embeds.AssetManager
	FontManager   resources.FontManager
	DefaultFont   font.Face
	GameManager   game.GameManager
}

// NewDependencies initializes all dependencies for the game.
func NewDependencies(opts ...Option) (*DefaultDependencyManager, error) {
	// Initialize options with defaults
	options := &Options{
		DefaultFont:     resources.DefaultFontName,
		DefaultFontSize: 16.0,
	}

	// Apply the functional options to override defaults
	for _, opt := range opts {
		opt(options)
	}

	// Initialize AssetManager if not provided
	assetManager := options.AssetManager
	if assetManager == nil {
		assetManager = embeds.NewDefaultAssetManager()
	}

	// Initialize FontManager if not provided
	fontManager := options.FontManager
	if fontManager == nil {
		fontManager = resources.NewDefaultFontManager(assetManager)
	}

	// Load the default font
	defaultFont, err := fontManager.LoadFont(options.DefaultFont, options.DefaultFontSize)
	if err != nil {
		return nil, err
	}

	// Initialize GameManager if not provided
	gameManager := options.GameManager
	if gameManager == nil {
		gameManager, err = game.NewDefaultGameManager(func(dst *ebiten.Image, text string, x, y int) {
			drawer := &font.Drawer{
				Dst:  dst,
				Src:  image.NewUniform(color.White),
				Face: defaultFont,
				Dot:  fixed.P(x, y),
			}
			drawer.DrawString(text)
		})
		if err != nil {
			return nil, err
		}
	}

	return &DefaultDependencyManager{
		AssetsManager: assetManager,
		FontManager:   fontManager,
		DefaultFont:   defaultFont,
		GameManager:   gameManager,
	}, nil
}

// GetAssetManager returns the asset manager.
func (d *DefaultDependencyManager) GetAssetManager() embeds.AssetManager {
	return d.AssetsManager
}

// GetFontManager returns the font manager.
func (d *DefaultDependencyManager) GetFontManager() resources.FontManager {
	return d.FontManager
}

// GetGameManager returns the game manager.
func (d *DefaultDependencyManager) GetGameManager() game.GameManager {
	return d.GameManager
}

// GetDefaultFont returns the default font.
func (d *DefaultDependencyManager) GetDefaultFont() font.Face {
	return d.DefaultFont
}

func (d *DefaultDependencyManager) GetTextWriter() func(dst *ebiten.Image, text string, x, y int) {

	return func(dst *ebiten.Image, text string, x, y int) {
		drawer := &font.Drawer{
			Dst:  dst,
			Src:  image.NewUniform(color.White),
			Face: d.DefaultFont,
			Dot:  fixed.P(x, y),
		}
		drawer.DrawString(text)
	}

}
