package main

import (
	"fmt"

	"github.com/CRRDerek/Neuroevolution/classifiers"
	"github.com/CRRDerek/Neuroevolution/evolution"
	"github.com/CRRDerek/Neuroevolution/games"
)

func main() {
	//testSaveJSON()
	//testXOR()
	//testTicTacToe()
	testCheckers()
	//testDepthOneCheckers()
}

// Seed a population of networks capable of learning XOR and then run neuroevolution
// on the XOR game.
func testXOR() {
	// Seed the initial population
	pop := make([]classifiers.Classifier, 100)
	for i := 0; i < 100; i++ {
		pop[i] = classifiers.RandomNetwork(3, 4, 1)
	}

	evolution.EvolveAgents(games.XorGame, games.XorGamePlayerMaker,
		2000, 256, 10, pop, evolution.Elimination_fitness)

}

func testSaveJSON() {
	evolved_agent := classifiers.RandomNetwork(28, 56, 9)

	fmt.Println("Training complete!")
	fmt.Println("Saving to file...")
	err := evolved_agent.SaveJSON("data/testSave.json")
	if err != nil {
		fmt.Println("Error saving agent: ", err)
	}

	fmt.Println("Loading from file...")
	// Load agent
	loaded_agent, err := classifiers.LoadJSON("data/testSave.json")
	if err != nil {
		fmt.Println("Error saving agent: ", err)
	}

	// Play tic tac toe against the user
	for {
		victor := games.TicTacToe(games.TicTacToePlayerMaker(loaded_agent), games.HumanTicTacToePlayer)
		if victor == -1 {
			fmt.Println("\n\nYou win!")
		} else if victor == 0 {
			fmt.Println("\n\nDraw!")
		} else if victor == 1 {
			fmt.Println("\n\nYou lose!")
		}
	}

}

// Seed a population of networks capable of learning Tic Tac Toe (input size 28,
// output size 9) and run neuroevolution to produce an agent that has evolved to
// play tic tac toe.
//
// Run tic tac toe games against the user indefinitely once the evolved agent is ready.
func testTicTacToe() {
	// Seed the initial population
	pop_size := 256
	pop := make([]classifiers.Classifier, pop_size)
	for i := 0; i < pop_size; i++ {
		pop[i] = classifiers.RandomNetwork(28, 56, 9)
	}

	// Evolve an agent capable of playing
	evolved_agent := evolution.EvolveAgents(games.TicTacToe, games.TicTacToePlayerMaker,
		512, 256, 4, pop, evolution.Elimination_fitness)

	fmt.Println("Training complete!")

	save(evolved_agent, "data/TicTacToe.json")

	// Play tic tac toe against the user
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

// Seed a population of networks capable of learning Checkers (input size 65,
// output size 24) and run neuroevolution to produce an agent that has evolved to
// play Checkers.
//
// Run Checkers games against the user indefinitely once the evolved agent is ready.
func testCheckers() {
	// Seed the initial population
	pop_size := 256
	pop := make([]classifiers.Classifier, pop_size)
	for i := 0; i < pop_size; i++ {
		pop[i] = classifiers.RandomNetwork(65, 130, 1)
	}

	// Run neuroevolution to produce an agent. The checkers games used by the
	// evolutionary algorithm will be cut off after 100 moves to prevent
	// random players from prolonging the game indefinitely.
	evolved_agent := evolution.EvolveAgents(games.MakeCheckers(32), games.ClassifierHeuristicPlayerMakerMaker(games.Checkers_make_move, games.CheckersTranslateInputs),
		100, 64, 4, pop, evolution.Elimination_fitness) // Each member of the population will be tested at maximum 128 times.
	// After 256 generations the algorithm concludes if it hasn't already spawned
	// an agent that can win 128 times for 4 generations.
	fmt.Println("Training complete!")

	save(evolved_agent, "data/Checkers.json")

	// Play checkers against the user indefinitely
	for {
		victor := games.Checkers(games.CheckersPlayerMaker(evolved_agent), games.HumanCheckersPlayer)
		if victor == -1 {
			fmt.Println("\n\nYou win!")
		} else if victor == 0 {
			fmt.Println("\n\nDraw!")
		} else if victor == 1 {
			fmt.Println("\n\nYou lose!")
		}
	}

}

// Test the user interface of checkers against a random player
func testRandomCheckers() {
	for {
		victor := games.Checkers(games.HumanCheckersPlayer, games.RandomPlayer)
		if victor == 1 {
			fmt.Println("\n\nYou win!")
		} else if victor == 0 {
			fmt.Println("\n\nDraw!")
		} else if victor == -1 {
			fmt.Println("\n\nYou lose!")
		}
	}
}

// Test the user interface of checkers against a random player
func testDepthOneCheckers() {
	for {
		victor := games.Checkers(games.HumanCheckersPlayer, games.DepthOneSearchPlayerMaker(games.Checkers_heuristic, games.Checkers_make_move))
		if victor == 1 {
			fmt.Println("\n\nYou win!")
		} else if victor == 0 {
			fmt.Println("\n\nDraw!")
		} else if victor == -1 {
			fmt.Println("\n\nYou lose!")
		}
	}
}

func save(c classifiers.Classifier, filename string) {
	err := c.SaveJSON(filename)
	if err != nil {
		fmt.Println("Error saving agent: ", err)
	}
	fmt.Println("Saved to file: ", filename)
}
