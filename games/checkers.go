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

	// Moves should store coordinate pairs as integers
	var moves []interface{}
	var player_move int
	var victor int

	return 0
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
