package config_test

import (
	"testing"

	"github.com/ghyter/misfits/internal/config"
	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/embeds"
	"github.com/ghyter/misfits/internal/game"
	"github.com/ghyter/misfits/internal/resources"
	"golang.org/x/image/font"
)

func TestNewDefaultDependencies_Success(t *testing.T) {
	// Initialize the DependencyManager with default dependencies
	dm, err := config.NewDefaultDependencies()
	if err != nil {
		t.Fatalf("Failed to initialize dependencies: %v", err)
	}

	// Test each dependency is registered and retrievable
	t.Run("AssetManager", func(t *testing.T) {
		if _, err := dependencies.Get[embeds.AssetManager](dm); err != nil {
			t.Errorf("AssetManager retrieval failed: %v", err)
		}
	})

	t.Run("FontManager", func(t *testing.T) {
		if _, err := dependencies.Get[resources.FontManager](dm); err != nil {
			t.Errorf("FontManager retrieval failed: %v", err)
		}
	})

	t.Run("GameManager", func(t *testing.T) {
		if _, err := dependencies.Get[game.GameManager](dm); err != nil {
			t.Errorf("GameManager retrieval failed: %v", err)
		}
	})

	t.Run("DefaultFont", func(t *testing.T) {
		if _, err := dependencies.Get[font.Face](dm); err != nil {
			t.Errorf("DefaultFont retrieval failed: %v", err)
		}
	})
}
