package games

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//   1 | 2 | 3
//   4 | 5 | 6
//   7 | 8 | 9
//
// A game state is stored as an array of 9 integers. 0 means no one has placed,
// -1 means X has placed, 1 means O has placed.
func TicTacToe(x_player Player, o_player Player) int {
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

		player_move = x_player(invert_game_state(game_state), moves).(int)
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

		player_move = o_player(game_state, moves).(int)
		// For now let's assume all players always return valid moves.
		// TODO Check that this move is a valid move

		game_state = move(game_state, player_move, 1)
		score = score_move(score, player_move, 1)
		victor = checkScore(score)
		if victor != 0 {
			return victor
		}
	}

	// What happened? Return a draw.
	return 0
}

// Given a game state for tic tac toe, return a list of valid moves
func calculate_moves(game_state [9]int) []interface{} {
	var moves []interface{}

	for i := 0; i < 9; i++ {
		if game_state[i] == 0 {
			moves = append(moves, i)
		}
	}

	return moves
}

// Given a game state, inverts the player names of each space so that either player
// can see the state as theirs
func invert_game_state(game_state [9]int) [9]int {
	inverse_game_state := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 9; i++ {
		inverse_game_state[i] = -game_state[i]
	}
	return inverse_game_state
}

// Given a game state, player move, and player number,
// plays the chosen move on the state and returns the
// new state.
func move(game_state [9]int, player_move int, player_number int) [9]int {
	game_state[player_move] = player_number
	return game_state
}

func score_move(score [8][2]int, player_move int, player_index int) [8][2]int {
	switch player_move {
	case 0:
		score[0][player_index] += 1
		score[3][player_index] += 1
		score[7][player_index] += 1
	case 1:
		score[0][player_index] += 1
		score[4][player_index] += 1
	case 2:
		score[0][player_index] += 1
		score[5][player_index] += 1
		score[6][player_index] += 1
	case 3:
		score[1][player_index] += 1
		score[3][player_index] += 1
	case 4:
		score[1][player_index] += 1
		score[4][player_index] += 1
		score[6][player_index] += 1
		score[7][player_index] += 1
	case 5:
		score[1][player_index] += 1
		score[5][player_index] += 1
	case 6:
		score[2][player_index] += 1
		score[3][player_index] += 1
		score[6][player_index] += 1
	case 7:
		score[2][player_index] += 1
		score[4][player_index] += 1
	case 8:
		score[2][player_index] += 1
		score[5][player_index] += 1
		score[7][player_index] += 1
	}
	return score
}

func checkScore(score [8][2]int) int {
	for i := 0; i < 8; i++ {
		if score[i][0] == 3 && score[i][1] == 0 {
			return 1
		}
		if score[i][0] == 0 && score[i][1] == 3 {
			return -1
		}
	}

	return 0

}

func TicTacToePlayerMaker(a Agent) Player {
	return func(game_state interface{}, moves []interface{}) interface{} {

		inputs := []float64{1.0}

		for i := 0; i < 9; i++ {
			val := game_state.([9]int)[i]
			if val == -1 {
				inputs = append(inputs, 1.0)
			} else {
				inputs = append(inputs, 0.0)
			}

			if val == 0 {
				inputs = append(inputs, 1.0)
			} else {
				inputs = append(inputs, 0.0)
			}

			if val == 1 {
				inputs = append(inputs, 1.0)
			} else {
				inputs = append(inputs, 0.0)
			}

		}

		predictions := a.Predict(inputs)

		max_choice := -1
		max_val := -999.0
		for i := 0; i < len(moves); i++ {
			move := moves[i].(int)
			val := predictions[move]
			if val > max_val {
				max_val = val
				max_choice = move
			}
		}

		return max_choice

	}
}

func HumanTicTacToePlayer(game_state interface{}, moves []interface{}) interface{} {
	state := game_state.([9]int)
	fmt.Printf("\n%v | %v | %v", state[0], state[1], state[2])
	fmt.Printf("\n%v | %v | %v", state[3], state[4], state[5])
	fmt.Printf("\n%v | %v | %v", state[6], state[7], state[8])
	fmt.Println("\nWhat is your move? ")
	fmt.Printf("\nPossible moves: %v", moves)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nEnter an integer 0-8: ")
		text, _ := reader.ReadString('\n')
		i, err := strconv.Atoi(strings.Trim(text, "\n\r"))

		if err != nil {
			fmt.Printf("\nError: %v", err)
		}

		fmt.Printf("\nYou entered: %v -> %v", text, i)
		for j := 0; j < len(moves); j++ {
			if moves[j] == i {
				return moves[j]
			}
		}

		fmt.Println("\nInvalid move. Please try again.")
	}
}
