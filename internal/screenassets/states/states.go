package states

type GameState int

func (s GameState) String() string {
	switch s {
	case Menu:
		return "Menu"
	case Playing:
		return "Playing"
	case Pause:
		return "Pause"
	case GameOver:
		return "Game Over"
	default:
		return "Unknown"
	}
}

const (
	Menu    GameState = iota // Menu state of the game
	Playing                  // Gameplay state
	Pause
	GameOver
)
