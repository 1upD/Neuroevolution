package games

import (
	"math/rand"

	"github.com/CRRDerek/Neuroevolution/classifiers"
)

// Player is a type of function that given a game state and a list of valid moves
// will return a valid move. The types of the game state, move items, and move
// return will vary depending on the game.
type Player func(game_state interface{}, moves []interface{}) interface{}

// Each game should have a separate factory function to create an player based
// on an agent and that particular game's configuration.
type PlayerMaker func(a classifiers.Classifier) Player

// A random player that selects a valid move for the current game.
// This player will work for any game that passes it a game state and a list of
// valid moves.
// It does not even use the game state.
func RandomPlayer(game_state interface{}, moves []interface{}) interface{} {
	n := rand.Intn(len(moves))
	return moves[n]
}

type Heuristic func(game_state interface{}) float64

type Valid_moves func(game_state interface{}) []interface{}

type Game_move func(game_state interface{}, move interface{}) interface{}

// Given a heuristic function and game move function, returns a player that looks
// at the next game state and picks the one with the highest value
func DepthOneSearchPlayerMaker(h Heuristic, m Game_move) Player {
	return func(game_state interface{}, moves []interface{}) interface{} {
		n := rand.Intn(len(moves))
		move := moves[n]
		moveVal := 0.0

		for i := 0; i < len(moves); i++ {
			new_board := m(game_state, moves[i])
			board_val := h(new_board)
			if board_val > moveVal {
				moveVal = board_val
				move = moves[i]
			}
		}

		return move
	}
}
