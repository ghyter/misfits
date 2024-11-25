package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestGameOptions_Defaults(t *testing.T) {
	opts := DefaultGameOptions()

	if opts.NumPlayers != 1 {
		t.Errorf("Expected default NumPlayers to be 1, got %d", opts.NumPlayers)
	}
	if opts.ScreenWidth != 800 || opts.ScreenHeight != 600 {
		t.Errorf("Expected default screen size to be 800x600, got %dx%d", opts.ScreenWidth, opts.ScreenHeight)
	}
}

func TestGameOptions_FunctionalOptions(t *testing.T) {
	opts, err := NewGameOptions(
		WithPlayers(2),
		WithScreenSize(1024, 768),
	)
	if err != nil {
		t.Fatalf("Failed to create GameOptions: %v", err)
	}

	if opts.NumPlayers != 2 {
		t.Errorf("Expected NumPlayers to be 2, got %d", opts.NumPlayers)
	}
	if opts.ScreenWidth != 1024 || opts.ScreenHeight != 768 {
		t.Errorf("Expected screen size to be 1024x768, got %dx%d", opts.ScreenWidth, opts.ScreenHeight)
	}
}

func TestGameOptions_Validation(t *testing.T) {
	_, err := NewGameOptions(WithPlayers(0))
	if err == nil || err.Error() != "NumPlayers must be greater than 0" {
		t.Fatalf("Expected validation error for NumPlayers <= 0, got: %v", err)
	}

	_, err = NewGameOptions(WithScreenSize(0, 600))
	if err == nil || err.Error() != "ScreenWidth and ScreenHeight must be greater than 0" {
		t.Fatalf("Expected validation error for invalid screen size, got: %v", err)
	}
}

func TestGameManager_NewGame(t *testing.T) {
	textWriter := func(dst *ebiten.Image, text string, x, y int) {}
	gm, err := NewDefaultGameManager(textWriter)
	if err != nil {
		t.Fatalf("Failed to create DefaultGameManager: %v", err)
	}

	game, err := gm.NewGame(
		WithPlayers(3),
		WithScreenSize(1280, 720),
	)
	if err != nil {
		t.Fatalf("Failed to create new game: %v", err)
	}

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
