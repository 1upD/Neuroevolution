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
		pop[i] = neuralnetwork.RandomNetwork(2, 8, 1)
	}

	evolution.EvolveAgents(games.XorGame, games.XorGamePlayerMaker,
		1000, 256, pop)

}
