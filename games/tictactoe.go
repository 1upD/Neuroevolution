package games

//   1 | 2 | 3
//   4 | 5 | 6
//   7 | 8 | 9
//
// A game state is stored as an array of 9 integers. 0 means no one has placed,
// -1 means X has placed, 1 means O has placed.
func tic_tac_toe(o_player Player, x_player Player) int {
	game_state := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	score := [8][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	var moves []interface{}
	var player_move int
	var victor int

	for {
		moves = calculate_moves(game_state)

		// If there are no moves remaining, call a draw
		if len(moves) == 0 {
			return 0
		}

		player_move = x_player(game_state, moves).(int)
		// For now let's assume all players always return valid moves.
		// TODO Check that this move is a valid move

		game_state = move(game_state, player_move, -1)
		score = score_move(score, player_move, 0)
		victor = checkScore(score)
		if victor != 0 {
			return victor
		}

		// O player
		moves = calculate_moves(game_state)

		// If there are no moves remaining, call a draw
		if len(moves) == 0 {
			return 0
		}

		player_move = o_player(invert_game_state(game_state), moves).(int)
		// For now let's assume all players always return valid moves.
		// TODO Check that this move is a valid move

		game_state = move(game_state, player_move, 1)
		score = score_move(score, player_move, 1)
		victor = checkScore(score)
		if victor != 0 {
			return victor
		}
	}

	return 0
}

// Given a game state for tic tac toe, return a list of valid moves
func calculate_moves(game_state [9]int) []interface{} {
	// TODO unimplemented
	return []interface{}{0}
}

func invert_game_state(game_state [9]int) [9]int {
	// TODO unimplemented
	return [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func move(game_state [9]int, player_move int, player_number int) [9]int {
	// TODO unimplemented
	return [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func score_move(score [8][2]int, player_move int, player_index int) [8][2]int {
	// TODO unimplemented
	return [8][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
}

func checkScore(score [8][2]int) int {
	// TODO unimplemented
	return 0
}
