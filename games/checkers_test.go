package games

import (
	"fmt"
	"testing"
)

// Scenarios
//
// Standard board
//
//	X	R	X	R	X	R	X	R
//	R	X	R	X	R	X	R	X
//	X	R	X	R	X	R	X	R
//	0	X	0	X	0	X	0	X
//	X	0	X	0	X	0	X	0
//	B	X	B	X	B	X	B	X
//	X	B	X	B	X	B	X	B
//	B	X	B	X	B	X	B	X
//
// game_state := [4][8]int{[8]int{-1, -1, -1, 0, 0, 1, 1, 1},
//		[8]int{-1, -1, -1, 0, 0, 1, 1, 1},
//		[8]int{-1, -1, -1, 0, 0, 1, 1, 1},
//		[8]int{-1, -1, -1, 0, 0, 1, 1, 1}}
//
//
// Blank board
//
//	X	0	X	0	X	0	X	0
//	0	X	0	X	0	X	0	X
//	X	0	X	0	X	0	X	0
//	0	X	0	X	0	X	0	X
//	X	0	X	0	X	0	X	0
//	0	X	0	X	0	X	0	X
//	X	0	X	0	X	0	X	0
//	0	X	0	X	0	X	0	X
//
// game_state := [4][8]int{
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0}}
//
// Odd capture
//
//	0	X	0	X	0	X	0	X	0
//	1	0	X	0	X	0	X	0	X
//	2	X	0	X	0	X	0	X	0
//	3	0	X	-1	X	-1	X	0	X
//	4	X	0	X	1	X	0	X	0
//	5	0	X	-1	X	-1	X	0	X
//	6	X	0	X	0	X	0	X	0
//	7	0	X	0	X	0	X	0	X
//
//		0	0	1	1	2	2	3	3
//
// game_state := [4][8]int{
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
//		[8]int{0, 0, -1, 1, -1, 0, 0, 0},
//		[8]int{0, 0, -1, 0, -1, 0, 0, 0},
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0}}
//
// Possible moves: [[1, 4, 0, 2], [1, 4, 2, 2]]
//
// Odd king capture
//
//	0	X	0	X	0	X	0	X	0
//	1	0	X	0	X	0	X	0	X
//	2	X	0	X	0	X	0	X	0
//	3	0	X	-1	X	-1	X	0	X
//	4	X	0	X	2	X	0	X	0
//	5	0	X	-1	X	-1	X	0	X
//	6	X	0	X	0	X	0	X	0
//	7	0	X	0	X	0	X	0	X
//
//		0	0	1	1	2	2	3	3

// game_state := [4][8]int{
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
//		[8]int{0, 0, -1, 2, -1, 0, 0, 0},
//		[8]int{0, 0, -1, 0, -1, 0, 0, 0},
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0}}
//
// Possible moves: [[1, 4, 0, 2], [1, 4, 2, 2], [1, 4, 0, 6], [1, 4, 2, 6]]
//
//
// Odd capture
//
//	0	X	0	X	0	X	0	X	0
//	1	0	X	0	X	0	X	0	X
//	2	X	-1	X	-2	X	0	X	0
//	3	0	X	1	X	0	X	0	X
//	4	X	-1	X	-2	X	0	X	0
//	5	0	X	0	X	0	X	0	X
//	6	X	0	X	0	X	0	X	0
//	7	0	X	0	X	0	X	0	X
//
//		0	0	1	1	2	2	3	3
//
// game_state := [4][8]int{
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
//		[8]int{0, 0, 0, -1, 1, -2, 0, 0},
//		[8]int{0, 0, 0, -1, 0, -2, 0, 0},
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0}}
//
// Possible moves: [[1, 3, 0, 1], [1, 3, 2, 1]]
//
// Odd king capture
//
//	0	X	0	X	0	X	0	X	0
//	1	0	X	0	X	0	X	0	X
//	2	X	-1	X	-2	X	0	X	0
//	3	0	X	1	X	0	X	0	X
//	4	X	-1	X	-2	X	0	X	0
//	5	0	X	0	X	0	X	0	X
//	6	X	0	X	0	X	0	X	0
//	7	0	X	0	X	0	X	0	X
//
//		0	0	1	1	2	2	3	3
//
// game_state := [4][8]int{
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
//		[8]int{0, 0, 0, -1, 2, -2, 0, 0},
//		[8]int{0, 0, 0, -1, 0, -2, 0, 0},
//		[8]int{0, 0, 0, 0, 0, 0, 0, 0}}
//
// Possible moves: [[1, 3, 0, 1], [1, 3, 2, 1], [1, 3, 0, 5], [1, 3, 2, 5]]

func TestCalculateCheckersMovesPerPiece(t *testing.T) {
	// TODO unimplemented
	t.Error("Test is unimplemented.")
}

func TestCalculateCheckersCapturesPerPiece(t *testing.T) {
	fmt.Println("Testing captures on standard board...")
	game_state := [4][8]int{[8]int{-1, -1, -1, 0, 0, 1, 1, 1},
		[8]int{-1, -1, -1, 0, 0, 1, 1, 1},
		[8]int{-1, -1, -1, 0, 0, 1, 1, 1},
		[8]int{-1, -1, -1, 0, 0, 1, 1, 1}}
	moves := []interface{}{}
	result := testCalculateCheckersCapturesPerPiece(game_state, [2]int{1, 5}, moves)
	if !result {
		t.Error("Failed test.")
	}

	fmt.Println("Testing captures on blank board...")
	game_state = [4][8]int{
		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
		[8]int{0, 0, 0, 0, 0, 0, 0, 0}}
	moves = []interface{}{}
	result = testCalculateCheckersCapturesPerPiece(game_state, [2]int{1, 5}, moves)
	if !result {
		t.Error("Failed test.")
	}

	fmt.Println("Testing captures on an even space...")
	game_state = [4][8]int{
		[8]int{0, 0, 0, 0, 0, 0, 0, 0},
		[8]int{0, 0, -1, 1, -1, 0, 0, 0},
		[8]int{0, 0, -1, 0, -1, 0, 0, 0},
		[8]int{0, 0, 0, 0, 0, 0, 0, 0}}
	moves = []interface{}{[4]int{1, 4, 0, 2}, [4]int{1, 4, 2, 2}}
	result = testCalculateCheckersCapturesPerPiece(game_state, [2]int{1, 5}, moves)
	if !result {
		t.Error("Failed test.")
	}

}

func testCalculateCheckersCapturesPerPiece(game_state [4][8]int, piece_to_move [2]int, expected_moves []interface{}) bool {
	result := calculate_checkers_captures_per_piece(game_state, piece_to_move, game_state[piece_to_move[0]][piece_to_move[1]] == 2)
	if len(expected_moves) != len(result) {
		fmt.Println("Expected move list size did not match result move list")
		fmt.Printf("\nElements did not match: \nExpected:\t%\nResult:\t%v\n", expected_moves, result)
		return false
	}
	for i := 0; i < len(expected_moves); i++ {
		for j := 0; j < len(expected_moves[i].([4]int)); j++ {
			if result[i].([4]int)[j] != expected_moves[i].([4]int)[j] {
				fmt.Printf("\nElements did not match: \nExpected:\t%\nResult:\t%v\n", expected_moves, result)
				return false
			}
		}
	}
	return true
}

func TestCheckersMakeMove(t *testing.T) {
	// TODO unimplemented
	t.Error("Test is unimplemented.")
}
