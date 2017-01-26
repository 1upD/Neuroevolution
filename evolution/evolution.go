package evolution

import (
	"fmt"
	"math/rand"

	"github.com/CRRDerek/Neuroevolution/classifiers"
	"github.com/CRRDerek/Neuroevolution/games"
)

// Function type for fitness functions
type fitnessFunc func(g games.Game, player games.Player, max_games int, fitnessChan *chan float64)

// Given a game, a player factory function, (specific to that game) the number of
// generations, maximum number of games, and an initial population of agents
//(which must all be the same type!) run an evolutionary algorithm and return
// the best agent.
func EvolveAgents(g games.Game, playerMaker games.PlayerMaker, generations int,
	max_games int, max_streak int, pop []classifiers.Classifier, f fitnessFunc) classifiers.Classifier {

	// Initialize an array of channels for each member of the population
	fitness_channels := make([]chan float64, len(pop))
	fitness_values := make([]float64, len(pop))

	// Initialize an array for the new population
	var new_pop []classifiers.Classifier

	// Initialize variables for max fitness and max agent
	max_fitness := -9999999999.0
	var max_agent classifiers.Classifier

	//Initialize each channel
	for i := 0; i < len(pop); i++ {
		fitness_channels[i] = make(chan float64)
	}

	streak := 0

	// Loop the algorithm for as many iterations are specified in the number of
	// generations.
	i := 0
	for {
		// Start a goroutine to test each member of the population.
		for j := 0; j < len(pop); j++ {
			index := j
			go f(g, playerMaker(pop[index]), max_games, &(fitness_channels[index]))
		}

		// Receive fitness values from channels and find the maximum fitness
		max_fitness = -9999999.0
		for j := 0; j < len(pop); j++ { // TODO Why was this -1?
			//			fmt.Println("Preparing to receive fitness ", j)
			fitness_values[j] = <-fitness_channels[j]
			//			fmt.Println("Received fitness ", j)
			if fitness_values[j] > max_fitness {
				max_fitness = fitness_values[j]
				max_agent = pop[j]
			}
		}

		if fitness_values[0] == float64(max_games) { // This used to check the maximum value
			// but I changed it to instead look at the inherited value. A better measure
			// of convergence of the population is whether the inherited network reached
			// the maximum number of wins because the games are random.
			// TODO This does not work with percentage based fitness.
			streak += 1
		} else {
			streak = 0
		}

		// Print generation info
		fmt.Println("Generation: ", i)
		fmt.Println("Max fitness: ", max_fitness)
		fmt.Println(fitness_values)

		// Iterate the generation number and return if the algorithm is complete.
		i++
		if i >= generations || streak >= max_streak {
			return max_agent
		}

		// Create a new array for the new population
		new_pop = make([]classifiers.Classifier, len(pop))
		new_pop[0] = max_agent

		// Create the next generation by mating based on fitness
		for j := 1; j < len(pop); j++ {
			p1 := weighted_selection(pop, fitness_values)
			p2 := weighted_selection(pop, fitness_values)
			new_pop[j] = p1.(classifiers.Classifier).Mate(p2.(classifiers.Classifier))
		}

		pop = new_pop
	}

}

// A fitness function that runs trials until either the maximum number of games is reached
// or the player being tested loses.
func Elimination_fitness(g games.Game, player games.Player, max_games int, fitnessChan *chan float64) {
	score := 0
	// Keep testing this player until the maximum number of games is
	// reached.
	for k := 0; k < max_games; k++ {
		// Play the game against a random opponent
		switch games.PlayerTrial(g, player) {
		// If the agent player wins, reward it
		case 1:
			score += 1
		// In case of a draw, don't stop
		case 0:
			score += 0 // After the code review, I changed the reward
			// for draws to 0. In Tic Tac Toe this produces better
			// results because a perfect player should win against
			// a random player. This does make it nearly impossible
			// to reach the maximum number of wins because some games
			// will always be draws
		// If they lose, break out of the loop.
		case -1:
			k = max_games
		}
	}

	// Send the score over the appropriate channel
	*fitnessChan <- float64(score)
}

// Fitness function that runs a fixed number of trials and computes fitness as a
// percentage of those trials. Players have an expected fitness of 50% against
// random players, so anything below 50 is treated as 0.
func Percentage_fitness(g games.Game, player games.Player, max_games int, fitnessChan *chan float64) {
	score := 0.0
	wins := 0
	draws := 0
	// Keep testing this player until the maximum number of games is
	// reached.
	for k := 0; k < max_games; k++ {
		// Play the game against a random opponent
		switch games.PlayerTrial(g, player) {
		// If the agent player wins, reward it
		case 1:
			wins += 1.0
		// Reward draws too.
		case 0:
			draws += 0.0 // After the code review, I changed the reward
			// for draws to 0. In Tic Tac Toe this produces better
			// results because a perfect player should win against
			// a random player. This does make it nearly impossible
			// to reach the maximum number of wins because some games
			// will always be draws
		}
	}

	// If every game is a win or a draw, treat it as 100%
	if wins+draws == max_games {
		score = 100
	} else {
		// Score fitness as a percentage - 100 is max, 50 is expected, 0 is min
		score = ((float64(wins) / float64(max_games)) * 100.0)
		if score < 0 {
			score = 0
		}
	}

	// Send the score over the appropriate channel
	*fitnessChan <- score
}

// Used to select pairs to mate for the next generation
// Based on a Stack Overflow answer:
// http://stackoverflow.com/a/3679747
func weighted_selection(items []classifiers.Classifier, weights []float64) classifiers.Classifier {
	total := 0.0
	for i := 0; i < len(weights); i++ {
		total += weights[i]
	}

	r := rand.Float64() * total
	upto := 0.0

	for i := 0; i < len(items); i++ {
		w := weights[i]
		if upto+w >= r {
			return items[i]
		}
		upto += w
	}
	fmt.Println("ERROR: Weighted selection failed.")
	return items[0]
}
