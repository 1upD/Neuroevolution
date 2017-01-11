package games

// Because even Tic Tac Toe was too complicated for my earliest tests, I've come
// up with a game that's even more trivial. It's called "the XOR game". Each
// player is given two binary inputs and is expected to answer yes or no.
// If both players guess correctly, it's a draw. If one player guesses correctly
// but the other doesn't, the player guessing correctly wins and the other loses.
//
//	0	0	0
//	0	1	1
//	1	0	1
//	1	1	0
//
// This will be my most basic test case and once I have the code working I can
// move on to tic-tac-toe, then to connect four, then maybe Othello, and then if
// I am extraordinairily ambitious, Go.

func XorGame(Player, Player) int {
	// Unimplemented
	return 0
}

func XorGamePlayerMaker(a Agent) Player {
	// Unimplemented
	return nil
}
