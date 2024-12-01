package resources

import (
	"fmt"
	"sync"

	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/embeds"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const DefaultFontName string = "DejaVuSans.ttf"

type FontManager interface {
	LoadFont(name string, size float64) (font.Face, error)
}

type DefaultFontManager struct {
	dm *dependencies.DependencyManager
	//assetManager embeds.AssetManager
	cache map[string]font.Face
	mu    sync.Mutex
}

// NewDefaultFontManager creates a new DefaultFontManager instance.
func NewDefaultFontManager(dm *dependencies.DependencyManager) (FontManager, error) {

	return &DefaultFontManager{
		dm: dm,

		cache: make(map[string]font.Face),
	}, nil
}

// LoadFont loads a font from the embeds and caches it.
// name: The file name of the font (e.g., "DejaVuSans.ttf").
// size: The size of the font in points.
func (fm *DefaultFontManager) LoadFont(name string, size float64) (font.Face, error) {
	cacheKey := fmt.Sprintf("%s-%f", name, size)

	// Check if the font is already cached
	fm.mu.Lock()
	defer fm.mu.Unlock()
	if face, exists := fm.cache[cacheKey]; exists {
		return face, nil
	}

	assetManager, err := dependencies.Get[embeds.AssetManager](fm.dm)
	// Load the font data from the AssetManager
	fontData, err := assetManager.Get(fmt.Sprintf("fonts/%s", name))
	if err != nil {
		return nil, fmt.Errorf("failed to load font %s: %w", name, err)
	}

	// Parse the font
	parsedFont, err := opentype.Parse(fontData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font %s: %w", name, err)
	}

	// Create a font.Face for the given size
	face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create font face for %s: %w", name, err)
	}

	// Cache the font and return it
	fm.cache[cacheKey] = face
	return face, nil
}
