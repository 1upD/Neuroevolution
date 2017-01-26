package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/CRRDerek/Neuroevolution/classifiers"
	"github.com/CRRDerek/Neuroevolution/evolution"
	"github.com/CRRDerek/Neuroevolution/games"
)

func main() {
	// Parse flags
	game_name := flag.String("game", "Tic Tac Toe", "Name of game to be played. Currently Tic Tac Toe and Checkers are supported.")
	filename := flag.String("filename", "None", "JSON file to load containing a neural network")
	generations := flag.Int("generations", 128, "Number of generations to evolve before returning the best network")
	max_games := flag.Int("maxgames", 1024, "Maximum number of games to play within a population")
	max_streak := flag.Int("streak", 4, "Number of generations that achieve the maximum score before ending the algorithm early")
	output := flag.String("output", "data\\results.json", "Name of a JSON file to write the results to")
	population := flag.Int("population", 256, "Number of individuals in the population")
	flag.Parse()

	// Declare variables
	var game games.Game
	var playerMaker games.PlayerMaker
	var humanPlayer games.Player
	var inputs int
	var hiddens int
	var outputs int
	var network classifiers.Classifier

	// Configure the correct game functions and network size based on which game the user selected
	if strings.ToLower(*game_name) == "tic tac toe" {
		game = games.TicTacToe
		playerMaker = games.TicTacToePlayerMaker
		humanPlayer = games.HumanTicTacToePlayer
		// Configure a network size for a Tic Tac Toe policy network
		inputs = 28
		hiddens = 56
		outputs = 9
	} else if strings.ToLower(*game_name) == "checkers" {
		// Impose a 1024 turn limit so the games don't get stuck in a loop
		game = games.MakeCheckers(1024)
		playerMaker = games.CheckersPlayerMaker
		humanPlayer = games.HumanCheckersPlayer
		// Configure a network size for a Checkers policy network
		inputs = 65
		hiddens = 256
		outputs = 24
	} else {
		fmt.Println("Please choose Tic Tac Toe or Checkers.")
		return
	}

	// If there is a file specificed, load the file instead of evolving
	if *filename != "None" {
		var err error
		network, err = classifiers.LoadJSON(*filename)
		if err != nil {
			fmt.Println("Error loading classifier: ", err)
			return
		}
	} else {
		// Seed the initial population
		pop := make([]classifiers.Classifier, *population)
		for i := 0; i < *population; i++ {
			pop[i] = classifiers.RandomNetwork(inputs, hiddens, outputs)
		}

		network = evolution.EvolveAgents(game, playerMaker, *generations, *max_games, *max_streak, pop, evolution.Elimination_fitness)

		// Save the results to a file
		save(network, *output)
	}

	// Make a player out of the network
	player := playerMaker(network)

	// Play the game indefinitely
	for {
		victor := game(player, humanPlayer)
		if victor == -1 {
			fmt.Println("\n\nYou win!")
		} else if victor == 0 {
			fmt.Println("\n\nDraw!")
		} else if victor == 1 {
			fmt.Println("\n\nYou lose!")
		}
	}
}

func DemoTicTacToe() {
	network, _ := classifiers.LoadJSON("data\\Final_TicTacToe.json")
	ai_player := games.TicTacToePlayerMaker(network)

	for {
		victor := games.TicTacToe(ai_player, games.HumanTicTacToePlayer)
		if victor == -1 {
			fmt.Println("\n\nYou win!")
		} else if victor == 0 {
			fmt.Println("\n\nDraw!")
		} else if victor == 1 {
			fmt.Println("\n\nYou lose!")
		}
	}

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
		64, 1024, 8, pop, evolution.Elimination_fitness)

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

func testEvolveCheckersPolicyNetwork() {
	testEvolveCheckers(games.CheckersPlayerMaker, "data/CheckersPolicyNetwork.json")
}

func testEvolveCheckersValueNetwork() {
	testEvolveCheckers(games.ClassifierHeuristicPlayerMakerMaker(games.Checkers_make_move, games.CheckersTranslateInputs), "data/CheckersValueNetwork.json")
}

// Seed a population of networks capable of learning Checkers (input size 65,
// output size 24) and run neuroevolution to produce an agent that has evolved to
// play Checkers.
//
// Run Checkers games against the user indefinitely once the evolved agent is ready.
func testEvolveCheckers(playerMaker games.PlayerMaker, filename string) {
	// Seed the initial population
	pop_size := 256
	pop := make([]classifiers.Classifier, pop_size)

	for i := 0; i < pop_size; i++ {
		pop[i] = classifiers.RandomNetwork(65, 256, 24)
	}

	// Run neuroevolution to produce an agent. The checkers games used by the
	// evolutionary algorithm will be cut off after 100 moves to prevent
	// random players from prolonging the game indefinitely.
	evolved_agent := evolution.EvolveAgents(games.MakeCheckers(1024), playerMaker,
		1024, 1024, 4, pop, evolution.Elimination_fitness) // Each member of the population will be tested at maximum 128 times.
	// After 256 generations the algorithm concludes if it hasn't already spawned
	// an agent that can win 128 times for 4 generations.
	fmt.Println("Training complete!")

	save(evolved_agent, filename)

	// Play checkers against the user indefinitely
	for {
		victor := games.Checkers(playerMaker(evolved_agent), games.HumanCheckersPlayer)
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

func testEvolvedNetworks() {
	// The first networks are all policy networks
	network_a, _ := classifiers.LoadJSON("data\\Checkers_01232017_1.json")
	player_a := games.CheckersPlayerMaker(network_a)

	network_b, _ := classifiers.LoadJSON("data\\Checkers_01232017_2.json")
	player_b := games.CheckersPlayerMaker(network_b)

	network_c, _ := classifiers.LoadJSON("data\\Checkers_01242017_1.json")
	player_c := games.CheckersPlayerMaker(network_c)

	// These are value networks
	network_d, _ := classifiers.LoadJSON("data\\Checkers_01242017_2.json")
	player_d := games.ClassifierHeuristicPlayerMakerMaker(games.Checkers_make_move, games.CheckersTranslateInputs)(network_d)

	network_e, _ := classifiers.LoadJSON("data\\Checkers_01242017_3.json")
	player_e := games.ClassifierHeuristicPlayerMakerMaker(games.Checkers_make_move, games.CheckersTranslateInputs)(network_e)

	players := []games.Player{player_a, player_b, player_c, player_d, player_e, games.RandomPlayer, games.DepthOneSearchPlayerMaker(games.Checkers_heuristic, games.Checkers_make_move)}
	player_names := []string{"Policy Network A", "Policy Network B", "Policy Network C", "Value Network D", "Value Network E", "Random Player", "Depth One Heuristic Player"}

	for i := 0; i < len(players); i++ {
		score := 0
		for j := 0; j < len(players); j++ {
			victor := games.MakeCheckers(1024)(players[j], players[i])
			if victor == -1 {
				score += 1
				fmt.Printf("%v defeated %v\n", player_names[i], player_names[j])
			} else if victor == 1 {
				fmt.Printf("%v was defeated by %v\n", player_names[i], player_names[j])
			} else {
				fmt.Printf("%v and %v tied!\n", player_names[i], player_names[j])
			}

		}

		fmt.Printf("\n%v won %v games\n***\n\n", player_names[i], score)
	}

}
