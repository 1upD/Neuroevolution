package games

import (
	"math/rand"

	"github.com/CRRDerek/Neuroevolution/classifiers"
)

// Because even Tic Tac Toe was too complicated for my earliest tests, I've come
// up with a game that's even more trivial. It's called "the XOR game". Each
// player is given two binary inputs and is expected to answer yes or no.
// If both players guess correctly, it's a draw. If one player guesses correctly
// but the other doesn't, the player guessing correctly wins and the other loses.
//
//	0	0	0
//	0	1	1
//	1	0	1
//	1	1	0
//
// This will be my most basic test case and once I have the code working I can
// move on to tic-tac-toe, then to connect four, then maybe Othello, and then if
// I am extraordinairily ambitious, Go.

func XorGame(p1 Player, p2 Player) int {
	var a float64
	var b float64
	var c int

	switch rand.Intn(4) {
	case 0:
		a = 0
		b = 0
		c = 0
	case 1:
		a = 0
		b = 1
		c = 1
	case 2:
		a = 1
		b = 0
		c = 1
	case 3:
		a = 1
		b = 1
		c = 0
	}

	p1move := p1([]float64{a, b}, []interface{}{0, 1})
	//	p2move := p2([]float64{a, b}, []interface{}{0, 1})

	if p1move == c {
		return 1
	} else {
		return -1
	}
	// Ignore player 2 entirely

	return 0
}

func XorGamePlayerMaker(a classifiers.Classifier) Player {
	return func(game_state interface{}, moves []interface{}) interface{} {
		inputs := []float64{1.0}

		inputs = append(inputs, game_state.([]float64)[0])
		inputs = append(inputs, game_state.([]float64)[1])

		prediction := a.Classify(inputs)
		if prediction[0] > 0.5 {
			return 1
		} else {
			return 0
		}

	}
}
