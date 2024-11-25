package main

import (
	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/hajimehoshi/ebiten/v2"
)

var dm dependencies.DependencyManager

func main() {

	dm, err := dependencies.NewDependencies()
	if err != nil {
		panic(err)
	}

	game, err := dm.GetGameManager().NewGame()
	if err != nil {
		panic(err)
	}

	ebiten.RunGame(game)
}
