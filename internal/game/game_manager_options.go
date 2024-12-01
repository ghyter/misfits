package game

import (
	"github.com/ghyter/misfits/internal/dependencies"
)

type GameManagerOption func(*GameManagerOptions)

type GameManagerOptions struct {
	dm *dependencies.DependencyManager
}

func NewGameManagerOptions(dm *dependencies.DependencyManager, opts ...GameManagerOption) (*GameManagerOptions, error) {
	options := &GameManagerOptions{
		dm: dm,
	}
	for _, opt := range opts {
		opt(options)
	}
	return options, nil
}
