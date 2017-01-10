package games

type Game interface {
	PlayGame(func(game_state interface{}, moves interface{}) interface{},
		func(game_state interface{}, moves interface{}) interface{}) int
}
