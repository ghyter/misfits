package states

type GameState int

func (s GameState) String() string {
	switch s {
	case Menu:
		return "Menu"
	case Game:
		return "Game"
	case Pause:
		return "Pause"
	case GameOver:
		return "Game Over"
	default:
		return "Unknown"
	}
}

const (
	Menu GameState = iota // Menu state of the game
	Game                  // Gameplay state
	Pause
	GameOver
)
