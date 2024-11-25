package dependencies

import (
	"fmt"
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
func NewDependencies(
	assetManager embeds.AssetManager,
	fontManager resources.FontManager,
	defaultFont *string,
	gameManager game.GameManager,
) (*DefaultDependencyManager, error) {
	// Initialize AssetManager if nil
	if assetManager == nil {
		assetManager = embeds.NewDefaultAssetManager()
	}

	// Initialize FontManager if nil
	if fontManager == nil {
		fontManager = resources.NewDefaultFontManager(assetManager)
	}
	var loadedFont font.Face
	var err error

	if defaultFont == nil {
		// Load the default font
		loadedFont, err = fontManager.LoadFont(resources.DefaultFontName, 16)
		if err != nil {
			return nil, fmt.Errorf("failed to load default font: %w", err)
		}
	}

	dm := &DefaultDependencyManager{
		AssetsManager: assetManager,
		FontManager:   fontManager,
		DefaultFont:   loadedFont,
	}

	// Initialize GameManager if nil
	if gameManager == nil {
		dm.GameManager, err = game.NewDefaultGameManager(dm.GetTextWriter())
		if err != nil {
			return nil, fmt.Errorf("failed to initialize game manager: %w", err)
		}
	}

	return dm, nil
}

// NewDefaultDependencies with error handling
func NewDefaultDependencies() (*DefaultDependencyManager, error) {
	return NewDependencies(nil, nil, nil, nil)
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
