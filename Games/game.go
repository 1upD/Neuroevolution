package games

// Every two player game must have a function to play the game given two appropriate
// player objects.
type Game func(Player, Player) int

// Plays a game pitting the given agent against a random player.
// Useful for fitness functions.
func PlayerTrial(g Game, p Player) int {
	return g(p, RandomPlayer)
}
