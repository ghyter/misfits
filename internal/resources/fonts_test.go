package resources

import (
	"errors"
	"testing"

	_ "embed"
)

// Embed a real `.ttf` font for testing
//
//go:embed testdata/DejaVuSans.ttf
var testFont []byte

// MockAssetManager is a mock implementation of embeds.AssetManager for testing.
type MockAssetManager struct {
	files map[string][]byte
}

func (m *MockAssetManager) Get(path string) ([]byte, error) {
	if data, exists := m.files[path]; exists {
		return data, nil
	}
	return nil, errors.New("file not found")
}

func TestDefaultFontManager_LoadFont(t *testing.T) {
	// Create a mock AssetManager with the embedded test font
	mockAssetManager := &MockAssetManager{
		files: map[string][]byte{
			"fonts/DejaVuSans.ttf": testFont,
		},
	}

	// Create the FontManager
	fontManager := NewDefaultFontManager(mockAssetManager)

	// Test loading a font
	fontFace, err := fontManager.LoadFont("DejaVuSans.ttf", 16)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if fontFace == nil {
		t.Fatal("Expected a valid font.Face, got nil")
	}

	// Test caching: Load the same font again and verify it comes from the cache
	cachedFontFace, err := fontManager.LoadFont("DejaVuSans.ttf", 16)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if cachedFontFace != fontFace {
		t.Fatal("Expected cached font.Face, got a new instance")
	}
}

func TestDefaultFontManager_LoadNonExistentFont(t *testing.T) {
	// Create a mock AssetManager with no font files
	mockAssetManager := &MockAssetManager{
		files: map[string][]byte{},
	}

	// Create the FontManager
	fontManager := NewDefaultFontManager(mockAssetManager)

	// Test loading a non-existent font
	_, err := fontManager.LoadFont("NonExistentFont.ttf", 16)
	if err == nil {
		t.Fatal("Expected an error for non-existent font, got nil")
	}
}

func TestDefaultFontManager_ConcurrentAccess(t *testing.T) {
	// Create a mock AssetManager with the embedded test font
	mockAssetManager := &MockAssetManager{
		files: map[string][]byte{
			"fonts/DejaVuSans.ttf": testFont,
		},
	}

	// Create the FontManager
	fontManager := NewDefaultFontManager(mockAssetManager)

	// Load fonts concurrently
	errs := make(chan error, 10)
	for i := 0; i < 10; i++ {
		go func() {
			_, err := fontManager.LoadFont("DejaVuSans.ttf", 16)
			errs <- err
		}()
	}

	// Check for errors
	for i := 0; i < 10; i++ {
		if err := <-errs; err != nil {
			t.Fatalf("Unexpected error during concurrent font loading: %v", err)
		}
	}
}
