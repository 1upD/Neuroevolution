package main

import (
	"fmt"

	"github.com/CRRDerek/Neuroevolution/evolution"
	"github.com/CRRDerek/Neuroevolution/games"
	"github.com/CRRDerek/Neuroevolution/neuralnetwork"
)

func main() {
	//testXOR()
	testTicTacToe()
}

func testXOR() {
	// Seed the initial population
	pop := make([]games.Agent, 100)
	for i := 0; i < 100; i++ {
		pop[i] = neuralnetwork.RandomNetwork(3, 4, 1)
	}

	evolution.EvolveAgents(games.XorGame, games.XorGamePlayerMaker,
		2000, 256, pop)

}

func testTicTacToe() {
	// Seed the initial population
	pop_size := 256
	pop := make([]games.Agent, pop_size)
	for i := 0; i < pop_size; i++ {
		pop[i] = neuralnetwork.RandomNetwork(28, 56, 9)
	}

	evolved_agent := evolution.EvolveAgents(games.TicTacToe, games.TicTacToePlayerMaker,
		10000, 1024, pop)

	fmt.Println("Training complete!")

	for {
		victor := games.TicTacToe(games.TicTacToePlayerMaker(evolved_agent), games.HumanTicTacToePlayer)
		if victor == -1 {
			fmt.Println("\n\nYou win!")
		} else if victor == 0 {
			fmt.Println("\n\nDraw!")
		} else if victor == 1 {
			fmt.Println("\n\nYou lose!")
		}
	}

}
