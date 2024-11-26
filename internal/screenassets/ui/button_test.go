package ui

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

// MockInputHandler simulates input for testing.
type MockInputHandler struct {
	CursorX, CursorY   int
	MouseButtonPressed map[ebiten.MouseButton]bool
}

func NewMockInputHandler() *MockInputHandler {
	return &MockInputHandler{
		MouseButtonPressed: make(map[ebiten.MouseButton]bool),
	}
}

func (m *MockInputHandler) CursorPosition() (int, int) {
	return m.CursorX, m.CursorY
}

func (m *MockInputHandler) IsMouseButtonPressed(button ebiten.MouseButton) bool {
	return m.MouseButtonPressed[button]
}

func TestButton_UpdateAndClick(t *testing.T) {
	mockInput := NewMockInputHandler()
	button := NewButton(10, 10, 100, 50, "Test", nil)
	button.Input = mockInput

	// Simulate hover
	mockInput.CursorX, mockInput.CursorY = 20, 20
	button.Update()
	assert.True(t, button.isHovered, "Expected button to be hovered")

	// Simulate click
	mockInput.MouseButtonPressed[ebiten.MouseButtonLeft] = true
	button.Update()
	assert.True(t, button.isClicked, "Expected button to be clicked")

	// Simulate release
	mockInput.MouseButtonPressed[ebiten.MouseButtonLeft] = false
	clicked := false
	button.OnClick = func() { clicked = true }
	button.Update()
	assert.False(t, button.isClicked, "Expected button not to be clicked")
	assert.True(t, clicked, "Expected OnClick to be called")
}
