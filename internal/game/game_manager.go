package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameManager interface {
	NewGame(opts ...GameOption) (Game, error)
}

type DefaultGameManager struct {
	drawText func(dst *ebiten.Image, text string, x, y int)
}

func NewDefaultGameManager(textWriter func(dst *ebiten.Image, text string, x, y int)) (GameManager, error) {
	return &DefaultGameManager{
		drawText: textWriter,
	}, nil
}

func (gm *DefaultGameManager) NewGame(opts ...GameOption) (Game, error) {
	// Create and validate GameOptions
	options, err := NewGameOptions(opts...)
	if err != nil {
		return nil, fmt.Errorf("invalid GameOptions: %w", err)
	}

	return &MisfitGame{
		drawText: gm.drawText,
		options:  options,
	}, nil
}
