package game

import (
	"errors"
)

type GameOptions struct {
	NumPlayers   int
	ScreenWidth  int
	ScreenHeight int
}

type GameOption func(*GameOptions)

func WithPlayers(playerCount int) GameOption {
	return func(o *GameOptions) {
		o.NumPlayers = playerCount
	}
}

func WithScreenSize(width, height int) GameOption {
	return func(o *GameOptions) {
		o.ScreenWidth = width
		o.ScreenHeight = height
	}
}

// DefaultGameOptions provides default settings for the game
func DefaultGameOptions() *GameOptions {
	return &GameOptions{
		NumPlayers:   1,
		ScreenWidth:  800,
		ScreenHeight: 600,
	}
}

// NewGameOptions applies functional options and ensures defaults are set
func NewGameOptions(opts ...GameOption) (*GameOptions, error) {
	options := DefaultGameOptions()

	for _, opt := range opts {
		opt(options)
	}

	// Validate options
	if err := options.Validate(); err != nil {
		return nil, err
	}

	return options, nil
}

// Validate ensures the GameOptions are valid
func (opts *GameOptions) Validate() error {
	if opts.NumPlayers <= 0 {
		return errors.New("NumPlayers must be greater than 0")
	}
	if opts.ScreenWidth <= 0 || opts.ScreenHeight <= 0 {
		return errors.New("ScreenWidth and ScreenHeight must be greater than 0")
	}
	return nil
}
