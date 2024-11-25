package main

import (
	"github.com/ghyter/misfits/internal/dependencies"
	"github.com/hajimehoshi/ebiten/v2"
)

var dm dependencies.DependencyManager

func main() {

	dm, err := dependencies.NewDefaultDependencies()
	if err != nil {
		panic(err)
	}

	game := dm.GetGameManager().NewGame()
	ebiten.RunGame(game)
}
