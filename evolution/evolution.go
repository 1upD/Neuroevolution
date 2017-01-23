package evolution

import (
	"fmt"
	"math/rand"

	"github.com/CRRDerek/Neuroevolution/classifiers"
	"github.com/CRRDerek/Neuroevolution/games"
)

// Given a game, a player factory function, (specific to that game) the number of
// generations, maximum number of games, and an initial population of agents
//(which must all be the same type!) run an evolutionary algorithm and return
// the best agent.
func EvolveAgents(g games.Game, playerMaker games.PlayerMaker, generations int,
	max_games int, max_streak int, pop []classifiers.Classifier) classifiers.Classifier {

	// Initialize an array of channels for each member of the population
	fitness_channels := make([]chan int, len(pop))
	fitness_values := make([]int, len(pop))

	// Initialize an array for the new population
	var new_pop []classifiers.Classifier

	// Initialize variables for max fitness and max agent
	max_fitness := -9999999999
	var max_agent classifiers.Classifier

	//Initialize each channel
	for i := 0; i < len(pop); i++ {
		fitness_channels[i] = make(chan int)
	}

	streak := 0

	// Loop the algorithm for as many iterations are specified in the number of
	// generations.
	i := 0
	for {
		// Start a goroutine to test each member of the population.
		for j := 0; j < len(pop); j++ {
			index := j
			go func() {
				// TODO Running the trials until a loss is unreliable because
				// of the randomness in the game. Consider instead running a fixed
				// number of trials and computing fitness as a percentage of those
				// trials. Agents have an expected fitness of 50%, treat them accordingly.
				// The only problem with running the maximum number of trials is that
				// it slows the algorithm down substantially; stopping after the first
				// loss quickly prunes the population. Perhaps I should instead
				// run a small number of trials, say 10, and if it reaches the maximum
				// continue running it until there is a loss? Then use the percentage
				// as the fitness score.

				score := 0
				player := playerMaker(pop[index])
				// Keep testing this player until the maximum number of games is
				// reached.
				for k := 0; k < max_games; k++ {
					// Play the game against a random opponent
					switch games.PlayerTrial(g, player) {
					// If the agent player wins, reward it
					case 1:
						score += 1
					// Reward draws too.
					case 0:
						score += 0 // After the code review, I changed the reward
						// for draws to 0. In Tic Tac Toe this produces better
						// results because a perfect player should win against
						// a random player. This does make it nearly impossible
						// to reach the maximum number of wins because some games
						// will always be draws
					// If they lose, break out of the loop.
					case -1:
						// k = max_games
						// I commented out the code that stops the trials after
						// a failure. This will take a lot longer but should counteract
						// some of the randomness of my earlier method.
					}
				}

				// Send the score over the appropriate channel
				//				fmt.Println("Preparing to send fitness ", index)
				fitness_channels[index] <- score
				//				fmt.Println("Sent fitness ", index)
			}()
		}

		// Receive fitness values from channels and find the maximum fitness
		max_fitness = -9999999
		for j := 0; j < len(pop); j++ { // TODO Why was this -1?
			//			fmt.Println("Preparing to receive fitness ", j)
			fitness_values[j] = <-fitness_channels[j]
			//			fmt.Println("Received fitness ", j)
			if fitness_values[j] > max_fitness {
				max_fitness = fitness_values[j]
				max_agent = pop[j]
			}
		}

		if fitness_values[0] == max_games { // This used to check the maximum value
			// but I changed it to instead look at the inherited value. A better measure
			// of convergence of the population is whether the inherited network reached
			// the maximum number of wins because the games are random
			// TODO If the fitness function becomes less random, consider replacing this
			// check with the maximum value again
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

func weighted_selection(items []classifiers.Classifier, weights []int) classifiers.Classifier {
	total := 0
	for i := 0; i < len(weights); i++ {
		total += weights[i]
	}

	r := rand.Intn(total)
	upto := 0

	for i := 0; i < len(items); i++ {
		w := weights[i]
		if upto+w >= r {
			//			fmt.Println("Selected item ", i, " with weight ", weights[i])
			return items[i]
		}
		upto += w
	}
	fmt.Println("ERROR: Weighted selection failed.")
	return items[0]
}
