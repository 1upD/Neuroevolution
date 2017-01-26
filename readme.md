# Neuroevolution in Go #

## Overview ##

For my CS344 Artificial Intelligence course final project, I wrote a neuroevolution algorithm in Golang. By treating the weights of neurons as dimensions of a search space, it attempts to evolve networks that are more capable of surviving in game playing environments. It has support for Tic Tac Toe and Checkers so far. 

## Usage ##

Usage of C:\GoWorkspace\src\github.com\CRRDerek\Neuroevolution\Neuroevolution.exe:

	-filename string
        JSON file to load containing a neural network (default "None")
 
	-game string
        Name of game to be played. Currently Tic Tac Toe and Checkers are supported. (default "Tic Tac Toe")

	-generations int
        Number of generations to evolve before returning the best network (default 128)
  
	-maxgames int
        Maximum number of games to play within a population (default 1024)
	
	-output string
        Name of a JSON file to write the results to (default "data\\results.json")
	
	-population int
        Number of individuals in the population (default 256)
	
	-streak int
        Number of generations that achieve the maximum score before ending the algorithm early (default 4)

		
## Repository ##

This project may be found at:
https://github.com/CRRDerek/Neuroevolution.git