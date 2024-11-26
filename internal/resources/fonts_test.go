package resources

import (
	"errors"
	"testing"

	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/embeds"
)

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

	// Create a mock DependencyManager and register the AssetManager
	mockDependencyManager := dependencies.NewDependencyManager()
	dependencies.Register(mockDependencyManager, func() (embeds.AssetManager, error) {
		return embeds.NewDefaultAssetManager(mockDependencyManager)
	})

	// Create the FontManager
	fontManager, err := NewDefaultFontManager(mockDependencyManager)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	// Test loading a font
	_, err = fontManager.LoadFont("DejaVuSans.ttf", 16)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
}

func TestDefaultFontManager_LoadNonExistentFont(t *testing.T) {
	// Create a mock AssetManager with no font files
	mockAssetManager := &MockAssetManager{
		files: map[string][]byte{},
	}

	// Create a mock DependencyManager and register the AssetManager
	mockDependencyManager := dependencies.NewDependencyManager()
	dependencies.Register(mockDependencyManager, func() (embeds.AssetManager, error) {
		return mockAssetManager, nil
	})

	// Create the FontManager
	fontManager, err := NewDefaultFontManager(mockDependencyManager)
	if err != nil {
		t.Fatal("Did not expect error from NewDefaultFontManager")
	}

	// Test loading a non-existent font
	_, err = fontManager.LoadFont("NonExistentFont.ttf", 16)
	if err == nil {
		t.Fatal("Expected an error for non-existent font, got nil")
	}
}
