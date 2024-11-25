package game

import "github.com/hajimehoshi/ebiten/v2"

// MockGameManager provides a mock implementation of the GameManager interface
type MockGameManager struct {
	NewGameCalled bool         // Tracks whether NewGame was called
	MockGame      *MockGame    // The mock game returned by NewGame
	Options       *GameOptions // Stores the options passed to NewGame
	Err           error        // Simulates errors for testing
}

// NewMockGameManager initializes a MockGameManager
func NewMockGameManager() *MockGameManager {
	return &MockGameManager{
		MockGame: &MockGame{},
	}
}

// NewGame tracks the call and returns the mock game
func (mgm *MockGameManager) NewGame(opts ...GameOption) (Game, error) {
	mgm.NewGameCalled = true

	// Apply options and simulate behavior
	options, err := NewGameOptions(opts...)
	if err != nil {
		mgm.Err = err
		return nil, err
	}

	mgm.Options = options
	return mgm.MockGame, mgm.Err
}

// MockGame provides a mock implementation of the Game interface
type MockGame struct {
	UpdateCalled bool
	DrawCalled   bool
	LayoutCalled bool
}

// Update mock implementation
func (mg *MockGame) Update() error {
	mg.UpdateCalled = true
	return nil
}

// Draw mock implementation
func (mg *MockGame) Draw(screen *ebiten.Image) {
	mg.DrawCalled = true
}

// Layout mock implementation
func (mg *MockGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	mg.LayoutCalled = true
	return 800, 600
}
