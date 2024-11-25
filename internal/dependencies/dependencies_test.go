package dependencies

import (
	"testing"

	"github.com/ghyter/misfits/internal/embeds"
	"github.com/ghyter/misfits/internal/game"
)

func TestNewDependencies_CustomGameManager(t *testing.T) {
	// Use the real AssetManager with embedded assets
	assetManager := embeds.NewDefaultAssetManager()

	// Create a mock GameManager
	mockGameManager := game.NewMockGameManager()

	// Initialize dependencies with the mock GameManager
	dm, err := NewDependencies(
		WithAssetManager(assetManager),
		WithGameManager(mockGameManager),
	)
	if err != nil {
		t.Fatalf("Failed to create dependencies with custom GameManager: %v", err)
	}

	// Validate the custom GameManager is used
	if dm.GetGameManager() != mockGameManager {
		t.Fatal("Expected custom GameManager to be used")
	}

	// Call NewGame on the mock GameManager
	gameInstance, err := dm.GetGameManager().NewGame()
	if err != nil {
		t.Fatal("Expected GetGameManager.NewGame to return not nil")
	}

	// Validate that NewGame was called
	if !mockGameManager.NewGameCalled {
		t.Fatal("Expected NewGame to be called on the GameManager")
	}

	// Validate the game instance is a MockGame
	mockGame, ok := gameInstance.(*game.MockGame)
	if !ok {
		t.Fatal("Expected game instance to be of type MockGame")
	}

	// Simulate game loop calls
	err = mockGame.Update()
	if err != nil {
		t.Fatalf("Expected no error from Update, got: %v", err)
	}
	mockGame.Draw(nil)
	mockGame.Layout(800, 600)

	// Validate the mock game methods were called
	if !mockGame.UpdateCalled {
		t.Fatal("Expected Update to be called on the game")
	}
	if !mockGame.DrawCalled {
		t.Fatal("Expected Draw to be called on the game")
	}
	if !mockGame.LayoutCalled {
		t.Fatal("Expected Layout to be called on the game")
	}
}
