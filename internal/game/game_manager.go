package game

import (
	"fmt"

	"github.com/ghyter/misfits/internal/dependencies"
)

type GameManager interface {
	NewGame(opts ...GameOption) (Game, error)
}

type DefaultGameManager struct {
	dm      *dependencies.DependencyManager
	Options *GameManagerOptions
}

func NewDefaultGameManager(dm *dependencies.DependencyManager, opt ...GameManagerOption) (GameManager, error) {

	options, err := NewGameManagerOptions(dm, opt...)
	if err != nil {
		return nil, err
	}

	return &DefaultGameManager{
		dm:      dm,
		Options: options,
	}, nil
}

func (gm *DefaultGameManager) NewGame(opts ...GameOption) (Game, error) {
	// Create and validate GameOptions
	options, err := NewGameOptions(opts...)
	if err != nil {
		return nil, fmt.Errorf("invalid GameOptions: %w", err)
	}

	game := &MisfitGame{
		dm:        gm.dm,
		options:   options,
		gmOptions: gm.Options,
	}

	game.InitUI()

	return game, nil
}
