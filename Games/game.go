package games

import (
	"math/rand"
)

// Player is a type of function that given a game state and a list of valid moves
// will return a valid move. The types of the game state, move items, and move
// return will vary depending on the game.
type Player func(game_state interface{}, moves []interface{}) interface{}

// Each game should have a separate factory function to create an player based
// on an agent and that particular game's configuration.
type PlayerMaker func(a Agent) Player

// Every two player game must have a function to play the game given two appropriate
// player objects.
type Game func(Player, Player) int

// A random player that selects a valid move for the current game.
// This player will work for any game that passes it a game state and a list of
// valid moves.
// It does not even use the game state.
func RandomPlayer(game_state interface{}, moves []interface{}) interface{} {
	n := rand.Intn(len(moves))
	return moves[n]
}

// Plays a game pitting the given agent against a random player.
// Useful for fitness functions.
func PlayerTrial(g Game, p Player) int {
	return g(p, RandomPlayer)
}
