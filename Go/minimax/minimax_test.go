package minimax

import (
	"checkers/packages/gameboard"
	"testing"

	"github.com/stretchr/testify/assert"
)

var initialGameBoard, err = gameboard.GetInitialGameBoard()

func TestTerminal(t *testing.T) {
	if err != nil {
		t.Errorf("Error while creating initial game board: %s", err)
	}
	tests := []struct {
		description string
		gameBoard   gameboard.GameBoard
		expected    bool
	}{
		{
			description: "Game not terminal",
			gameBoard:   initialGameBoard,
			expected:    false,
		},
		{
			description: "Game terminal",
			gameBoard: gameboard.GameBoard{
				GameBoard:  generateTerminalBoard("R"),
				CurrPlayer: "G",
			},
			expected: true,
		},
		{
			description: "Game terminal",
			gameBoard: gameboard.GameBoard{
				GameBoard:  generateTerminalBoard("G"),
				CurrPlayer: "G",
			},
			expected: true,
		}, {
			description: "Game terminal",
			gameBoard: gameboard.GameBoard{
				GameBoard:  generateTerminalBoardNoMoveLeft(),
				CurrPlayer: "R",
			},
			expected: true,
		},
	}

	for _, test := range tests {
		result := Terminal(test.gameBoard)
		assert.Equal(t, result, test.expected, "Test case: %s", test.description)
	}
}

func TestEvaluateHeuristicValue(t *testing.T) {
	tests := []struct {
		description string
		gameBoard   gameboard.GameBoard
		expected    float64
	}{
		{
			description: "Evaluate heuristic value for player R",
			gameBoard: gameboard.GameBoard{
				GameBoard:  initialGameBoard.GameBoard,
				CurrPlayer: "R",
			},
			expected: 30,
		},
		{
			description: "Evaluate heuristic value for player G",
			gameBoard: gameboard.GameBoard{
				GameBoard:  initialGameBoard.GameBoard,
				CurrPlayer: "G",
			},
			expected: -30,
		},
	}

	for _, test := range tests {
		result := evaluateHeuristicValue(test.gameBoard)
		assert.Equal(t, result, test.expected, "Test case: %s", test.description)
	}
}

func generateTerminalBoard(team string) [][]gameboard.Piece {

	return [][]gameboard.Piece{
		{gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 0}, Team: team, King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 3}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 4}, Team: team, King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 7}, Team: team, King: false}},
		{gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 3}, Team: team, King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 7}, Team: " ", King: false}},

		{gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 3}, Team: team, King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 7}, Team: " ", King: false}},

		{gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 3}, Team: team, King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 7}, Team: " ", King: false}},

		{gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 3}, Team: team, King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 7}, Team: " ", King: false}},

		{gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 3}, Team: team, King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 7}, Team: " ", King: false}},

		{gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 3}, Team: team, King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 7}, Team: " ", King: false}},

		{gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 3}, Team: team, King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 7}, Team: " ", King: false}},
	}
}
func generateTerminalBoardNoMoveLeft() [][]gameboard.Piece {

	return [][]gameboard.Piece{
		{gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 0}, Team: "R", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 3}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 0, Column: 7}, Team: " ", King: false}},
		{gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 1}, Team: "G", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 3}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 1, Column: 7}, Team: " ", King: false}},
		{gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 2}, Team: "G", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 3}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 2, Column: 7}, Team: " ", King: false}},
		{gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 3}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 3, Column: 7}, Team: " ", King: false}},
		{gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 3}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 4, Column: 7}, Team: " ", King: false}},
		{gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 3}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 5, Column: 7}, Team: " ", King: false}},
		{gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 3}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 6, Column: 7}, Team: " ", King: false}},
		{gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 0}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 1}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 2}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 3}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 4}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 5}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 6}, Team: " ", King: false},
			gameboard.Piece{Position: gameboard.Tuple{Row: 7, Column: 7}, Team: " ", King: false}},
	}
}
