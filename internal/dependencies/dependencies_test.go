package dependencies

import (
	"testing"
)

func TestNewDependencies_Defaults(t *testing.T) {
	deps, err := NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if deps.GetAssetManager() == nil {
		t.Fatal("AssetManager is nil")
	}

	if deps.GetFontManager() == nil {
		t.Fatal("FontManager is nil")
	}

	if deps.GetGameManager() == nil {
		t.Fatal("GameManager is nil")
	}

	if deps.GetDefaultFont() == nil {
		t.Fatal("DefaultFont is nil")
	}
}

/*func TestNewDependencies_WithMocks(t *testing.T) {
	mockAssetManager := embeds.NewMockAssetManager()                  // hypothetical mock
	mockFontManager := resources.NewMockFontManager(mockAssetManager) // hypothetical mock
	mockGameManager := game.NewMockGameManager()                      // hypothetical mock

	deps, err := NewDependencies(mockAssetManager, mockFontManager, mockGameManager)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if deps.GetAssetManager() != mockAssetManager {
		t.Fatal("Expected the mock asset manager, but got something else")
	}

	if deps.GetFontManager() != mockFontManager {
		t.Fatal("Expected the mock font manager, but got something else")
	}

	if deps.GetGameManager() != mockGameManager {
		t.Fatal("Expected the mock game manager, but got something else")
	}
}
*/
