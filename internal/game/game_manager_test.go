package game

import (
	"errors"
	"testing"

	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/embeds"
	"github.com/ghyter/misfits/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

// Mock implementations for testing
type MockFontManager struct{}

func (m *MockFontManager) LoadFont(name string, size float64) (font.Face, error) {
	if name == resources.DefaultFontName {
		return nil, nil // Simulate a valid font.Face
	}
	return nil, errors.New("font not found")
}

type MockAssetManager struct{}

func (m *MockAssetManager) Get(path string) ([]byte, error) {
	if path == "fonts/DejaVuSans.ttf" {
		return []byte("mock font data"), nil
	}
	return nil, errors.New("asset not found")
}

func (m *MockAssetManager) GetEbitenImage(path string) (*ebiten.Image, error) {
	if path == "images/test_image.png" {
		// Create a dummy Ebiten image for testing
		img := ebiten.NewImage(100, 100) // Mock an image with 100x100 pixels
		return img, nil
	}
	return nil, errors.New("image not found")
}

func TestGameManager_NewGame(t *testing.T) {
	// Mock DependencyManager
	mockDependencyManager := dependencies.NewDependencyManager()

	// Register mock AssetManager
	dependencies.Register(mockDependencyManager, func(dm *dependencies.DependencyManager) (embeds.AssetManager, error) {
		return &MockAssetManager{}, nil
	})

	// Register mock FontManager
	dependencies.Register(mockDependencyManager, func(dm *dependencies.DependencyManager) (resources.FontManager, error) {
		return &MockFontManager{}, nil
	})

	// Create GameManager
	gm, err := NewDefaultGameManager(mockDependencyManager)
	if err != nil {
		t.Fatalf("Failed to create DefaultGameManager: %v", err)
	}

	// Create a new game using GameManager
	game, err := gm.NewGame(
		WithPlayers(3),
		WithScreenSize(1280, 720),
	)
	if err != nil {
		t.Fatalf("Failed to create new game: %v", err)
	}

	// Validate the created game
	misfitGame, ok := game.(*MisfitGame)
	if !ok {
		t.Fatal("Expected game to be of type MisfitGame")
	}

	if misfitGame.options.NumPlayers != 3 {
		t.Errorf("Expected NumPlayers to be 3, got %d", misfitGame.options.NumPlayers)
	}
	if misfitGame.options.ScreenWidth != 1280 || misfitGame.options.ScreenHeight != 720 {
		t.Errorf("Expected screen size to be 1280x720, got %dx%d", misfitGame.options.ScreenWidth, misfitGame.options.ScreenHeight)
	}
}

func TestGameManager_InvalidGameOptions(t *testing.T) {
	// Mock DependencyManager
	mockDependencyManager := dependencies.NewDependencyManager()

	// Register mock AssetManager and FontManager
	dependencies.Register(mockDependencyManager, func(dm *dependencies.DependencyManager) (embeds.AssetManager, error) {
		return &MockAssetManager{}, nil
	})
	dependencies.Register(mockDependencyManager, func(dm *dependencies.DependencyManager) (resources.FontManager, error) {
		return &MockFontManager{}, nil
	})

	// Create GameManager
	gm, err := NewDefaultGameManager(mockDependencyManager)
	if err != nil {
		t.Fatalf("Failed to create DefaultGameManager: %v", err)
	}

	// Attempt to create a game with invalid options
	_, err = gm.NewGame(WithPlayers(0)) // Invalid number of players
	if err == nil {
		t.Fatal("Expected error for invalid GameOptions, got nil")
	}
}
