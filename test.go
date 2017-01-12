package main

import (
	"github.com/CRRDerek/Neuroevolution/evolution"
	"github.com/CRRDerek/Neuroevolution/games"
	"github.com/CRRDerek/Neuroevolution/neuralnetwork"
)

func main() {
	// Seed the initial population
	pop := make([]games.Agent, 100)
	for i := 0; i < 100; i++ {
		pop[i] = neuralnetwork.RandomNetwork(28, 56, 9)
	}

	evolution.EvolveAgents(games.TicTacToe, games.TicTacToePlayerMaker,
		10000, 256, pop)

}
