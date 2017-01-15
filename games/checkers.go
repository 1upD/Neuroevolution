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
//X	R	X	R	X	R	X	R
//R	X	R	X	R	X	R	X
//X	R	X	R	X	R	X	R
//0	X	0	X	0	X	0	X
//X	0	X	0	X	0	X	0
//B	X	B	X	B	X	B	X
//X	B	X	B	X	B	X	B
//B	X	B	X	B	X	B	X
//
//0	|	-1	-1	-1	-1
//1	|	-1	-1	-1	-1
//2	|	-1	-1	-1	-1
//3	|	 0	 0	 0	 0
//4	|	 0	 0	 0	 0
//5	|	 1	 1	 1	 1
//6	|	 1	 1	 1	 1
//7	|	 1	 1	 1	 1
//		 -	 -	 - 	 -
//		 0	 1	 2	 3
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
	game_state := [4][8]int{[8]int{-1, -1, -1, 0, 0, 1, 1, 1},
		[8]int{-1, -1, -1, 0, 0, 1, 1, 1},
		[8]int{-1, -1, -1, 0, 0, 1, 1, 1},
		[8]int{-1, -1, -1, 0, 0, 1, 1, 1}}

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
func calculate_checkers_moves(game_state [4][8]int) []interface{} {
	moves := []interface{}{}
	// TODO Use goroutines and channels to speed this up!

	for i := 0; i < 4; i++ {
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
func calculate_checkers_moves_per_piece(game_state [4][8]int, checker [2]int, isKing bool) []interface{} {
	moves := []interface{}{}
	x := checker[0]
	y := checker[1]

	if y > 0 && game_state[x][y-1] == 0 {
		moves = append(moves, [4]int{x, y, x, y - 1})
	}

	if isKing && y < 7 && game_state[x][y+1] == 0 {
		moves = append(moves, [4]int{x, y, x, y + 1})
	}

	if checker[1]%2 == 0 {
		// Even rows
		if y > 0 && x < 3 {
			if game_state[x+1][y-1] == 0 {
				moves = append(moves, [4]int{x, y, x + 1, y - 1})
			}
		}
		// Kings can move backwards
		if isKing && y < 7 && x > 0 {
			if game_state[x-1][y+1] == 0 {
				moves = append(moves, [4]int{x, y, x - 1, y + 1})
			}
		}

	} else {
		// Odd rows
		if y > 0 && x > 0 {
			if game_state[x-1][y-1] == 0 {
				moves = append(moves, [4]int{x, y, x - 1, y - 1})
			}
		}
		// Kings can move backwards
		if isKing && y < 7 && x < 3 {
			if game_state[x+1][y+1] == 0 {
				moves = append(moves, [4]int{x, y, x + 1, y + 1})
			}
		}
	}

	return moves
}

// Calculate capture moves for a given black checker on a checkers board
func calculate_checkers_captures_per_piece(game_state [4][8]int, checker [2]int, isKing bool) []interface{} {
	moves := []interface{}{}
	//	x := checker[0]
	//	y := checker[1]
	return moves
}

// Given a checkers board configuration, flip the board so that red player is now
// black player and vice versa. This way all players will see themselves as black
// player.
func checkers_board_flip(game_state [4][8]int) [4][8]int {
	// TODO unimplemented
	return game_state
}

// Given a black player move, flip the board so that it is played as red player
func flip_move(move [4]int) [4]int {
	// TODO unimplemented
	return move
}

// Given a board state and a valid move, make the move
// Returns a board state and the number of captures
func checkers_make_move(game_state [4][8]int, move [4]int) ([4][8]int, int) {
	// Pick up a piece
	game_state[move[0]][move[1]] = 0

	// TODO Complete this section
	// Check if this move is a capture and remove opposing pieces
	captured := 0

	// Place the piece
	game_state[move[2]][move[3]] = 1

	// Return the game state
	return game_state, captured
}

// This function prints a game state to the console and prompts the user to select a move.
func HumanCheckersPlayer(game_state interface{}, moves []interface{}) interface{} {
	state := game_state.([4][8]int)
	fmt.Printf("\n0:\tX\t%v\tX\t%v\tX\t%v\tX\t%v", state[0][0], state[1][0], state[2][0], state[3][0])
	fmt.Printf("\n1:\t%v\tX\t%v\tX\t%v\tX\t%v\tX", state[0][1], state[1][1], state[2][1], state[3][1])
	fmt.Printf("\n2:\tX\t%v\tX\t%v\tX\t%v\tX\t%v", state[0][2], state[1][2], state[2][2], state[3][2])
	fmt.Printf("\n3:\t%v\tX\t%v\tX\t%v\tX\t%v\tX", state[0][3], state[1][3], state[2][3], state[3][3])
	fmt.Printf("\n4:\tX\t%v\tX\t%v\tX\t%v\tX\t%v", state[0][4], state[1][4], state[2][4], state[3][4])
	fmt.Printf("\n5:\t%v\tX\t%v\tX\t%v\tX\t%v\tX", state[0][5], state[1][5], state[2][5], state[3][5])
	fmt.Printf("\n6:\tX\t%v\tX\t%v\tX\t%v\tX\t%v", state[0][6], state[1][6], state[2][6], state[3][6])
	fmt.Printf("\n7:\t%v\tX\t%v\tX\t%v\tX\t%v\tX", state[0][7], state[1][7], state[2][7], state[3][7])
	fmt.Println("\n\n-:\t0\t0\t1\t1\t2\t2\t3\t3\t")

	fmt.Println("\nWhat is your move? ")
	fmt.Printf("\nPossible moves: %v", moves)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nChoose an X coordinate: ")
		fmt.Println("\nEnter an integer 0-3: ")
		text, _ := reader.ReadString('\n')
		x1, err := strconv.Atoi(strings.Trim(text, "\n\r"))

		if err != nil {
			fmt.Printf("\nError: %v", err)
		}

		fmt.Println("\nChoose a Y coordinate: ")
		fmt.Println("\nEnter an integer 0-7: ")
		text, _ = reader.ReadString('\n')
		y1, err := strconv.Atoi(strings.Trim(text, "\n\r"))

		if err != nil {
			fmt.Printf("\nError: %v", err)
		}

		fmt.Println("\nChoose an X coordinate: ")
		fmt.Println("\nEnter an integer 0-3: ")
		text, _ = reader.ReadString('\n')
		x2, err := strconv.Atoi(strings.Trim(text, "\n\r"))

		if err != nil {
			fmt.Printf("\nError: %v", err)
		}

		fmt.Println("\nChoose a Y coordinate: ")
		fmt.Println("\nEnter an integer 0-7: ")
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
