package minimax

import (
	"checkers/packages/gameboard"
	"fmt"
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
				GameBoard:  GenerateTerminalBoardNoMoveLeft(),
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

func TestPerformAction(t *testing.T) {
	gameBoard := generateTerminalBoard("R")

	g := gameboard.GameBoard{
		GameBoard:  gameBoard,
		CurrPlayer: "G"}
	new_game_board := performAction(g, gameboard.Action{Start: gameboard.Tuple{Row: 0, Column: 0}, End: gameboard.Tuple{Row: 1, Column: 1}})
	assert.Equal(t, gameBoard, new_game_board.GameBoard, "Test case: %s", "performAction")

	gameBoard = initialGameBoard.GameBoard
	g = gameboard.GameBoard{
		GameBoard:  gameBoard,
		CurrPlayer: "R"}
	fmt.Print(new_game_board)

	new_game_board = performAction(g, gameboard.Action{Start: gameboard.Tuple{Row: 0, Column: 0}, End: gameboard.Tuple{Row: 1, Column: 1}})
	assert.NotEqual(t, gameBoard, new_game_board.GameBoard, "Test case: %s", "performAction")
}

func TestOpposite(t *testing.T) {
	tests := []struct {
		description string
		team        string
		expected    string
		smash       bool
	}{
		{
			description: "Opposite of G",
			team:        "G",
			expected:    "R",
			smash:       false,
		},
		{
			description: "Opposite of R",
			team:        "R",
			expected:    "G",
			smash:       false,
		},
		{
			description: "Turn again",
			team:        "R",
			expected:    "R",
			smash:       true,
		},
		{
			description: "Turn again",
			team:        "G",
			expected:    "G",
			smash:       true,
		},
	}
	result := ""
	for _, test := range tests {
		result = opposite(test.team, test.smash)
		assert.Equal(t, result, test.expected, "Test case: %s", test.description)
	}
}
func TestPossibleActions(t *testing.T) {
	gameBoard := initialGameBoard
	fmt.Print(gameBoard)
	actions := possibleActions(gameBoard, gameBoard.CurrPlayer)
	outcome := []gameboard.Action{
		{Start: gameboard.Tuple{Row: 5, Column: 1}, End: gameboard.Tuple{Row: 4, Column: 0}},
		{Start: gameboard.Tuple{Row: 5, Column: 1}, End: gameboard.Tuple{Row: 4, Column: 2}},
		{Start: gameboard.Tuple{Row: 5, Column: 3}, End: gameboard.Tuple{Row: 4, Column: 2}},
		{Start: gameboard.Tuple{Row: 5, Column: 3}, End: gameboard.Tuple{Row: 4, Column: 4}},
		{Start: gameboard.Tuple{Row: 5, Column: 5}, End: gameboard.Tuple{Row: 4, Column: 4}},
		{Start: gameboard.Tuple{Row: 5, Column: 5}, End: gameboard.Tuple{Row: 4, Column: 6}},
		{Start: gameboard.Tuple{Row: 5, Column: 7}, End: gameboard.Tuple{Row: 4, Column: 6}}}
	assert.Equal(t, actions, outcome, "Test case: %s", "possibleActions")
	assert.NotEqual(t, actions, []gameboard.Action{}, "Test case: %s", "possibleActions")

	fmt.Print(actions)
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
func GenerateTerminalBoardNoMoveLeft() [][]gameboard.Piece {

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
