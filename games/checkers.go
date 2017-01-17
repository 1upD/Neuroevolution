package games

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Checkers
//
// X	R	X	R	X	R	X	R
// R	X	R	X	R	X	R	X
// X	R	X	R	X	R	X	R
// 0	X	0	X	0	X	0	X
// X	0	X	0	X	0	X	0
// B	X	B	X	B	X	B	X
// X	B	X	B	X	B	X	B
// B	X	B	X	B	X	B	X
//
//0	|	 0	-1	 0	-1 	 0	-1	 0	-1
//1	|	-1	 0	-1	 0	-1	 0	-1	 0
//2	|	 0	-1	 0	-1	 0	-1	 0	-1
//3	|	 0	 0	 0	 0	 0	 0	 0	 0
//4	|	 0	 0	 0	 0	 0	 0	 0	 0
//5	|	 1	 0	 1	 0	 1	 0	 1	 0
//6	|	 0	 1	 0	 1	 0	 1	 0	 1
//7	|	 1	 0	 1	 0	 1	 0	 1	 0
//		 -	 -	 - 	 -
//		 0	 1	 2	 3	4	5	6	7
//
//
// Possible moves
// y % 2 != 0 && x==0
// y + 1
// else y== 7
// y -1, y-1 x-1
//
// y % 2 == 0 && x==3
// y - 1
// else y == 6
// y -1, y-1 x+1
func Checkers(black_player Player, red_player Player) int {
	// There are four functional columns with 8 rows.
	// I used columns first for easier indexing.
	game_state := [8][8]int{
		[8]int{0, -1, 0, 0, 0, 1, 0, 1},
		[8]int{-1, 0, -1, 0, 0, 0, 1, 0},
		[8]int{0, -1, 0, 0, 0, 1, 0, 1},
		[8]int{-1, 0, -1, 0, 0, 0, 1, 0},
		[8]int{0, -1, 0, 0, 0, 1, 0, 1},
		[8]int{-1, 0, -1, 0, 0, 0, 1, 0},
		[8]int{0, -1, 0, 0, 0, 1, 0, 1},
		[8]int{-1, 0, -1, 0, 0, 0, 1, 0}}

	// Number of pieces taken from the opponent. When this number reaches the
	// number of checkers on that side (12) the game is over.
	red_score := 0
	black_score := 0

	// Moves should store arrays of four integers where the first two represent
	// a coordinate pair of the piece to be moved and the second two represent the
	// space to move to.
	var moves []interface{}
	var move_score int
	var player_move [4]int

	// Main game loop
	for red_score < 12 && black_score < 12 {
		// Black player move
		moves = calculate_checkers_moves(game_state)
		for len(moves) > 0 {
			//			fmt.Println("Black player's turn")
			player_move = black_player(game_state, moves).([4]int)
			game_state, move_score = checkers_make_move(game_state, player_move)
			moves = calculate_checkers_captures_per_piece(game_state, [2]int{player_move[2], player_move[3]}, game_state[player_move[2]][player_move[3]] == 2)
			black_score += move_score
		}

		// Flip the board in preparation for red player
		game_state = checkers_board_flip(game_state)
		moves = calculate_checkers_moves(game_state)

		// Red player move
		for len(moves) > 0 {
			//			fmt.Println("Red player's turn")

			player_move = red_player(game_state, moves).([4]int)
			game_state, move_score = checkers_make_move(game_state, player_move)
			moves = calculate_checkers_captures_per_piece(game_state, [2]int{player_move[2], player_move[3]}, game_state[player_move[2]][player_move[3]] == 2)
			red_score += move_score
		}

		// Flip board back
		game_state = checkers_board_flip(game_state)

	}

	if black_score > red_score {
		return 1
	} else {
		return -1
	}
}

// Calculates all possible normal moves for the black player on a checkers board
func calculate_checkers_moves(game_state [8][8]int) []interface{} {
	moves := []interface{}{}
	// TODO Use goroutines and channels to speed this up!

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if game_state[i][j] > 0 {
				moves = append(moves, calculate_checkers_moves_per_piece(game_state, [2]int{i, j}, game_state[i][j] == 2)...)
				moves = append(moves, calculate_checkers_captures_per_piece(game_state, [2]int{i, j}, game_state[i][j] == 2)...)
			}
		}
	}

	return moves
}

// Calculate capture moves for a given black checker on a checkers board
func calculate_checkers_moves_per_piece(game_state [8][8]int, checker [2]int, isKing bool) []interface{} {
	moves := []interface{}{}
	x := checker[0]
	y := checker[1]

	// Upper left diagonal
	if x > 0 && y > 0 && game_state[x-1][y-1] == 0 {
		moves = append(moves, [4]int{x, y, x - 1, y - 1})
	}

	// Lower left diagonal
	if isKing && x > 0 && y < 8 && game_state[x-1][y+1] == 0 {
		moves = append(moves, [4]int{x, y, x - 1, y + 1})
	}

	// Upper right diagonal
	if x < 7 && y > 0 && game_state[x+1][y-1] == 0 {
		moves = append(moves, [4]int{x, y, x + 1, y - 1})

	}

	// Lower right diagonal
	if isKing && x < 7 && y < 7 && game_state[x+1][y+1] == 0 {
		moves = append(moves, [4]int{x, y, x + 1, y + 1})
	}

	return moves
}

// Calculate capture moves for a given black checker on a checkers board
func calculate_checkers_captures_per_piece(game_state [8][8]int, checker [2]int, isKing bool) []interface{} {
	moves := []interface{}{}
	x := checker[0]
	y := checker[1]

	// Upper left diagonal
	if x > 1 && y > 1 && game_state[x-1][y-1] < 0 && game_state[x-2][y-2] == 0 {
		moves = append(moves, [4]int{x, y, x - 2, y - 2})
	}

	// Lower left diagonal
	if isKing && x > 1 && y < 6 && game_state[x-1][y+1] < 0 && game_state[x-2][y+2] == 0 {
		moves = append(moves, [4]int{x, y, x - 2, y + 2})
	}

	// Upper right diagonal
	if x < 6 && y > 1 && game_state[x+1][y-1] < 0 && game_state[x+2][y-2] == 0 {
		moves = append(moves, [4]int{x, y, x + 2, y - 2})

	}

	// Lower right diagonal
	if isKing && x < 6 && y < 6 && game_state[x+1][y+1] < 0 && game_state[x+2][y+2] == 0 {
		moves = append(moves, [4]int{x, y, x + 2, y + 2})
	}

	return moves
}

// Given a checkers board configuration, flip the board so that red player is now
// black player and vice versa. This way all players will see themselves as black
// player.
func checkers_board_flip(game_state [8][8]int) [8][8]int {
	flip_state := [8][8]int{
		[8]int{0, -1, 0, 0, 0, 1, 0, 1},
		[8]int{-1, 0, -1, 0, 0, 0, 1, 0},
		[8]int{0, -1, 0, 0, 0, 1, 0, 1},
		[8]int{-1, 0, -1, 0, 0, 0, 1, 0},
		[8]int{0, -1, 0, 0, 0, 1, 0, 1},
		[8]int{-1, 0, -1, 0, 0, 0, 1, 0},
		[8]int{0, -1, 0, 0, 0, 1, 0, 1},
		[8]int{-1, 0, -1, 0, 0, 0, 1, 0}}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			flip_state[i][j] = -1 * game_state[7-i][7-j]
		}
	}

	return flip_state
}

// Given a board state and a valid move, make the move
// Returns a board state and the number of captures
func checkers_make_move(game_state [8][8]int, move [4]int) ([8][8]int, int) {
	//	fmt.Println("Move made: ", move)

	isKing := false
	captured := 0

	// Pick up a piece
	// Is it already a king?
	if game_state[move[0]][move[1]] == 2 {
		isKing = true
	}

	game_state[move[0]][move[1]] = 0

	// TODO Complete this section
	// Check if this move is a capture and remove opposing pieces
	// If the difference in Y is 2, this move is a capture
	//	fmt.Println("\nIs this a capture? ", (move[3]-move[1])*(move[3]-move[1]))
	if (move[3]-move[1])*(move[3]-move[1]) == 4 {
		captured_x := (move[0] + move[2]) / 2
		captured_y := (move[1] + move[3]) / 2

		game_state[captured_x][captured_y] = 0
		captured = 1
	}

	// Should the piece be kinged?
	if move[3] == 0 {
		isKing = true
	}

	// Place the piece
	if isKing {
		game_state[move[2]][move[3]] = 2

	} else {
		game_state[move[2]][move[3]] = 1

	}

	// Return the game state
	return game_state, captured
}

// Given an agent that accepts 65 inputs and classifies 24 outputs, this function
// creates a player for that agent that may be used in Checkers.
//
// Breakdown of inputs and outputs:
// I tried to minimize the number of inputs and outputs to speed up the agent.
// The Checkerboard is fed into the network as an 8x4 grid, ommitting the light
// spaces on the board
//
// The first input is an activation neuron that is always 1.
//
// Each space is given one input with a value of 0 for no piece, -1 for red piece
// or 1 for black piece. These three cases should be separate binary inputs but
// for maximum speed I've opted to consider them opposites.
//
// The final 32 inputs represent whether the piece at the given space is a king
// or not.
//
// As for the outputs, the 24 outputs represent two sets of twelve coordinates on
// the 8x4 grid represented in binary.
//
// The player's move is chosen by examining each valid move, summing up the
// predictions on that moves coordinates, and choosing the maximum summed prediction.
func CheckersPlayerMaker(a Agent) Player {
	return func(game_state interface{}, moves []interface{}) interface{} {
		// Activation input - always on
		inputs := []float64{1.0}

		// The first 32 inputs represent which side the space belongs to
		for i := 0; i < 8; i++ {
			j := 0

			if i%2 == 0 {
				i = 1
			}

			for j < 8 {
				val := game_state.([8][8]int)[i][j]
				if val < 0 {
					inputs = append(inputs, -1.0)
				} else if val > 0 {
					inputs = append(inputs, 1.0)
				} else {
					inputs = append(inputs, 0.0)
				}

				j += 2
			}

		}

		// The last 32 inputs represent whether or not the piece on a space is a king
		for i := 0; i < 8; i++ {
			j := 0

			if i%2 == 0 {
				i = 1
			}

			for j < 8 {
				val := game_state.([8][8]int)[i][j]
				if val*val == 4 {
					inputs = append(inputs, 1.0)
				} else {
					inputs = append(inputs, 0.0)
				}

				j += 2
			}

		}

		prediction := a.Predict(inputs)

		max_choice := moves[0]
		max_val := -999.0
		for i := 0; i < len(moves); i++ {
			move := moves[i].([4]int)
			move_val := 0.0

			// First x-coord
			x1 := move[0] / 2 // Do integer division on the X-coords because
			// we can ignore half of the spaces

			move_val += prediction[x1] // Range: 0 - 3

			y1 := move[1]
			move_val += prediction[4+y1] // Range 4 - 11

			// First x-coord
			x2 := move[2] / 2 // Do integer division on the X-coords because
			// we can ignore half of the spaces

			move_val += prediction[12+x2] // Range: 12 - 15

			y2 := move[3]
			move_val += prediction[16+y2] // Range 16 - 23

			if move_val > max_val {
				max_val = move_val
				max_choice = move
			}
		}

		return max_choice

	}
}

// This function prints a game state to the console and prompts the user to select a move.
func HumanCheckersPlayer(game_state interface{}, moves []interface{}) interface{} {
	state := game_state.([8][8]int)
	fmt.Printf("\n0:\t \t%v\t \t%v\t \t%v\t \t%v", state[1][0], state[3][0], state[5][0], state[7][0])
	fmt.Printf("\n1:\t%v\t \t%v\t \t%v\t \t%v\t ", state[0][1], state[2][1], state[4][1], state[6][1])
	fmt.Printf("\n2:\t \t%v\t \t%v\t \t%v\t \t%v", state[1][2], state[3][2], state[5][2], state[7][2])
	fmt.Printf("\n3:\t%v\t \t%v\t \t%v\t \t%v\t ", state[0][3], state[2][3], state[4][3], state[6][3])
	fmt.Printf("\n4:\t \t%v\t \t%v\t \t%v\t \t%v", state[1][4], state[3][4], state[5][4], state[7][4])
	fmt.Printf("\n5:\t%v\t \t%v\t \t%v\t \t%v\t ", state[0][5], state[2][5], state[4][5], state[6][5])
	fmt.Printf("\n6:\t \t%v\t \t%v\t \t%v\t \t%v", state[1][6], state[3][6], state[5][6], state[7][6])
	fmt.Printf("\n7:\t%v\t \t%v\t \t%v\t \t%v\t ", state[0][7], state[2][7], state[4][7], state[6][7])
	fmt.Println("\n\n-:\t0\t1\t2\t3\t4\t5\t6\t7\t")

	fmt.Println("\nWhat is your move? ")
	fmt.Printf("\nPossible moves: %v", moves)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nChoose an X coordinate: ")
		text, _ := reader.ReadString('\n')
		x1, err := strconv.Atoi(strings.Trim(text, "\n\r"))

		if err != nil {
			fmt.Printf("\nError: %v", err)
		}

		fmt.Println("\nChoose a Y coordinate: ")
		text, _ = reader.ReadString('\n')
		y1, err := strconv.Atoi(strings.Trim(text, "\n\r"))

		if err != nil {
			fmt.Printf("\nError: %v", err)
		}

		fmt.Println("\nChoose an X coordinate: ")
		text, _ = reader.ReadString('\n')
		x2, err := strconv.Atoi(strings.Trim(text, "\n\r"))

		if err != nil {
			fmt.Printf("\nError: %v", err)
		}

		fmt.Println("\nChoose a Y coordinate: ")
		text, _ = reader.ReadString('\n')
		y2, err := strconv.Atoi(strings.Trim(text, "\n\r"))

		if err != nil {
			fmt.Printf("\nError: %v", err)
		}

		move := [4]int{x1, y1, x2, y2}

		fmt.Printf("\nYou entered: %v", move)
		for j := 0; j < len(moves); j++ {
			if moves[j].([4]int)[0] == move[0] && moves[j].([4]int)[1] == move[1] && moves[j].([4]int)[2] == move[2] && moves[j].([4]int)[3] == move[3] {
				return moves[j]
			}
		}

		fmt.Println("\nInvalid move. Please try again.")
	}
}
