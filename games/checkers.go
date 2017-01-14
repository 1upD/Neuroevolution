package games

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
			moves = calculate_checkers_captures_per_piece(game_state, [2]int{player_move[2], player_move[3]})
			black_score += move_score
		}

		// Flip the board in preparation for red player
		game_state = checkers_board_flip(game_state)
		moves = calculate_checkers_moves(game_state)

		// Red player move
		for len(moves) > 0 {
			player_move = red_player(game_state, moves).([4]int)
			game_state, move_score = checkers_make_move(game_state, player_move)
			moves = calculate_checkers_moves_per_piece(game_state, [2]int{player_move[2], player_move[3]})
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
			if game_state[i][j] == 1 {
				moves = append(moves, calculate_checkers_moves_per_piece(game_state, [2]int{i, j})...)
				moves = append(moves, calculate_checkers_captures_per_piece(game_state, [2]int{i, j})...)
			}
		}
	}

	return moves
}

// Calculate capture moves for a given black checker on a checkers board
func calculate_checkers_moves_per_piece(game_state [4][8]int, checker [2]int) []interface{} {
	// TODO unimplemented
	return nil
}

// Calculate capture moves for a given black checker on a checkers board
func calculate_checkers_captures_per_piece(game_state [4][8]int, checker [2]int) []interface{} {
	// TODO unimplemented
	return nil
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
	// TODO unimplemented
	return game_state, 0
}
