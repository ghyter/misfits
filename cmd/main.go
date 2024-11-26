package main

import (
	"github.com/ghyter/misfits/internal/config"
	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/ghyter/misfits/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	dm, err := config.NewDefaultDependencies()
	if err != nil {
		panic(err)
	}
	gameManager, err := dependencies.Get[game.GameManager](dm)
	if err != nil {
		panic(err)
	}
	game, err := gameManager.NewGame()
	if err != nil {
		panic(err)
	}
	// Apply Ebiten settings
	cfg := config.GetEbitenConfig()
	cfg.Apply()

	ebiten.RunGame(game)
}
