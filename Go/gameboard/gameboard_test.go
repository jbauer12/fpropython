package gameboard

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestGetInitialGameBoard(t *testing.T) {
	gameBoard, err := GetInitialGameBoard()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	assert.Equal(t, len(gameBoard.GameBoard), 8, "Expected game board to have 8 rows")

	for _, row := range gameBoard.GameBoard {
		assert.Equal(t, len(row), 8, "Expected game board to have 8 columns")
	}
	gameboard_flat := lo.Flatten(gameBoard.GameBoard)
	green := lo.Filter(gameboard_flat, func(item Piece, i int) bool {
		return item.Team == "G"
	})
	red := lo.Filter(gameboard_flat, func(item Piece, i int) bool {
		return item.Team == "R"
	})
	assert.Equal(t, len(green), 12, "Expected game board to have 12 green pieces")
	assert.Equal(t, len(red), 12, "Expected game board to have 12 red pieces")

}

func TestMakeNewGameBoardAfterMove(t *testing.T) {
	initialGameBoard, _ := GetInitialGameBoard()
	action := Action{Start: Tuple{5, 1}, End: Tuple{4, 0}}

	smash := false
	king := false
	newGameBoard := MakeNewGameBoardAfterMove(initialGameBoard, action, smash, king)

	assert.NotEqual(t, initialGameBoard, newGameBoard, "Expected new game board to be different from the initial game board")
	assert.NotEqual(t, initialGameBoard.CurrPlayer, newGameBoard.CurrPlayer, "Expected new game board to have the same current player as the initial game board")
	assert.Equal(t, initialGameBoard.GameBoard[action.Start.Row][action.Start.Column], Piece{Position: Tuple{5, 1}, Team: "G", King: false}, "Expected initial game board to have a piece at the start position")
	assert.Equal(t, newGameBoard.GameBoard[action.Start.Row][action.Start.Column], Piece{Position: Tuple{5, 1}, Team: " ", King: false}, "Expected new game board to have no piece at the start position")
	assert.Equal(t, newGameBoard.GameBoard[action.End.Row][action.End.Column], Piece{Position: Tuple{4, 0}, Team: "G", King: false}, "Expected new game board to have a piece at the end position")
	assert.Equal(t, initialGameBoard.GameBoard[action.End.Row][action.End.Column], Piece{Position: Tuple{4, 0}, Team: " ", King: false}, "Expected initial game board to have no piece at the end position")
	assert.Equal(t, initialGameBoard.GameBoard[4][0], Piece{Position: Tuple{4, 0}, Team: " ", King: false}, "Expected initial game board to have no piece at the end position")

	smash_board := GenerateBoardWithSmash()
	smash = true
	king = false
	action = Action{Start: Tuple{4, 2}, End: Tuple{6, 4}}
	newGameBoard = MakeNewGameBoardAfterMove(smash_board, action, smash, king)
	assert.Equal(t, newGameBoard.GameBoard[4][2], Piece{Position: Tuple{4, 2}, Team: " ", King: false}, "Expected new game board to have no piece at the start position")
	assert.Equal(t, newGameBoard.GameBoard[5][3], Piece{Position: Tuple{5, 3}, Team: " ", King: false}, "Expected new game board to have no piece at the intermediate position")
	assert.Equal(t, newGameBoard.GameBoard[6][4], Piece{Position: Tuple{6, 4}, Team: "R", King: false}, "Expected new game board to have a piece at the end position")
	assert.Equal(t, smash_board.GameBoard[6][4], Piece{Position: Tuple{6, 4}, Team: " ", King: false}, "Expected initial game board to have no piece at the end position")
	assert.Equal(t, smash_board.GameBoard[5][3], Piece{Position: Tuple{5, 3}, Team: "G", King: false}, "Expected initial game board to have a piece at the intermediate position")
	assert.Equal(t, smash_board.GameBoard[4][2], Piece{Position: Tuple{4, 2}, Team: "R", King: false}, "Expected initial game board to have a piece at the start position")

	king_board := GenerateBoardWithCheckerMove("R")
	king = true
	smash = false
	action = Action{Start: Tuple{6, 0}, End: Tuple{7, 1}}
	newGameBoard = MakeNewGameBoardAfterMove(king_board, action, smash, king)
	assert.Equal(t, newGameBoard.GameBoard[6][0], Piece{Position: Tuple{6, 0}, Team: " ", King: false}, "Expected new game board to have no piece at the start position")
	assert.Equal(t, newGameBoard.GameBoard[7][1], Piece{Position: Tuple{7, 1}, Team: "R", King: true}, "Expected new game board to have a piece at the end position")
	king_board = GenerateBoardWithCheckerMove("G")
	king = true
	smash = false
	action = Action{Start: Tuple{1, 6}, End: Tuple{0, 7}}
	newGameBoard = MakeNewGameBoardAfterMove(king_board, action, smash, king)
	assert.Equal(t, newGameBoard.GameBoard[1][6], Piece{Position: Tuple{1, 6}, Team: " ", King: false}, "Expected new game board to have no piece at the start position")
	assert.Equal(t, newGameBoard.GameBoard[0][7], Piece{Position: Tuple{0, 7}, Team: "G", King: true}, "Expected new game board to have a piece at the end position")
}

func TestOpposite(t *testing.T) {
	tests := []struct {
		inputTeam        string
		inputSmash       bool
		expectedOpposite string
	}{
		{"G", true, "G"},
		{"R", false, "G"},
		{"G", false, "R"},
	}

	for _, test := range tests {
		result := opposite(test.inputTeam, test.inputSmash)
		assert.Equal(t, result, test.expectedOpposite, "Expected opposite to be %s no %s", test.expectedOpposite, result)
	}
}

func GenerateBoardWithSmash() GameBoard {
	gameBoard := [][]Piece{
		// Row 0
		{Piece{Position: Tuple{0, 0}, Team: " ", King: false}, Piece{Position: Tuple{0, 1}, Team: " ", King: false}, GetPieceWithRightTeam(0, 2), Piece{Position: Tuple{0, 3}, Team: " ", King: false},
			GetPieceWithRightTeam(0, 4), Piece{Position: Tuple{0, 5}, Team: " ", King: false}, Piece{Position: Tuple{0, 6}, Team: " ", King: false}, Piece{Position: Tuple{0, 7}, Team: " ", King: false}},
		// Row 1
		{Piece{Position: Tuple{1, 0}, Team: " ", King: false}, GetPieceWithRightTeam(1, 1), Piece{Position: Tuple{1, 2}, Team: " ", King: false}, Piece{Position: Tuple{1, 3}, Team: " ", King: false},
			Piece{Position: Tuple{1, 4}, Team: " ", King: false}, Piece{Position: Tuple{1, 5}, Team: " ", King: false}, Piece{Position: Tuple{1, 6}, Team: " ", King: false}, Piece{Position: Tuple{1, 7}, Team: " ", King: false}},
		// Row 2
		{GetPieceWithRightTeam(2, 0), Piece{Position: Tuple{2, 1}, Team: " ", King: false}, GetPieceWithRightTeam(2, 2), Piece{Position: Tuple{2, 3}, Team: " ", King: false},
			Piece{Position: Tuple{2, 4}, Team: " ", King: false}, Piece{Position: Tuple{2, 5}, Team: " ", King: false}, Piece{Position: Tuple{2, 6}, Team: " ", King: false}, Piece{Position: Tuple{2, 7}, Team: " ", King: false}},
		// Row 3
		{Piece{Position: Tuple{3, 0}, Team: " ", King: false}, Piece{Position: Tuple{3, 1}, Team: " ", King: false}, GetPieceWithRightTeam(3, 2), Piece{Position: Tuple{3, 3}, Team: " ", King: false},
			Piece{Position: Tuple{3, 4}, Team: " ", King: false}, Piece{Position: Tuple{3, 5}, Team: " ", King: false}, Piece{Position: Tuple{3, 6}, Team: " ", King: false}, Piece{Position: Tuple{3, 7}, Team: " ", King: false}},
		// Row 4
		{Piece{Position: Tuple{4, 0}, Team: " ", King: false}, GetPieceWithRightTeam(4, 1), Piece{Position: Tuple{4, 2}, Team: "R", King: false}, Piece{Position: Tuple{4, 3}, Team: " ", King: false},
			Piece{Position: Tuple{4, 4}, Team: " ", King: false}, Piece{Position: Tuple{4, 5}, Team: " ", King: false}, Piece{Position: Tuple{4, 6}, Team: " ", King: false}, Piece{Position: Tuple{4, 7}, Team: " ", King: false}},
		// Row 5
		{GetPieceWithRightTeam(5, 0), Piece{Position: Tuple{5, 1}, Team: " ", King: false}, Piece{Position: Tuple{5, 2}, Team: " ", King: false}, GetPieceWithRightTeam(5, 3),
			Piece{Position: Tuple{5, 4}, Team: " ", King: false}, GetPieceWithRightTeam(5, 5), Piece{Position: Tuple{5, 6}, Team: " ", King: false}, GetPieceWithRightTeam(5, 7)},
		// Row 6
		{Piece{Position: Tuple{6, 0}, Team: " ", King: false}, GetPieceWithRightTeam(6, 1), Piece{Position: Tuple{6, 2}, Team: " ", King: false}, Piece{Position: Tuple{6, 3}, Team: " ", King: false},
			Piece{Position: Tuple{6, 4}, Team: " ", King: false}, Piece{Position: Tuple{6, 5}, Team: " ", King: false}, Piece{Position: Tuple{6, 6}, Team: " ", King: false}, Piece{Position: Tuple{6, 7}, Team: " ", King: false}},
		// Row 7
		{GetPieceWithRightTeam(7, 0), Piece{Position: Tuple{7, 1}, Team: " ", King: false}, GetPieceWithRightTeam(7, 2), Piece{Position: Tuple{7, 3}, Team: " ", King: false},
			GetPieceWithRightTeam(7, 4), GetPieceWithRightTeam(7, 5), GetPieceWithRightTeam(7, 6), GetPieceWithRightTeam(7, 7)},
	}

	return GameBoard{
		GameBoard:  gameBoard,
		CurrPlayer: "R",
	}
}

func GenerateBoardWithCheckerMove(CurrPlayer string) GameBoard {
	gameBoard := [][]Piece{
		// Row 0
		{Piece{Position: Tuple{0, 0}, Team: " ", King: false}, Piece{Position: Tuple{0, 1}, Team: " ", King: false}, GetPieceWithRightTeam(0, 2), Piece{Position: Tuple{0, 3}, Team: " ", King: false},
			GetPieceWithRightTeam(0, 4), Piece{Position: Tuple{0, 5}, Team: " ", King: false}, Piece{Position: Tuple{0, 6}, Team: " ", King: false}, Piece{Position: Tuple{0, 7}, Team: " ", King: false}},
		// Row 1
		{Piece{Position: Tuple{1, 0}, Team: " ", King: false}, GetPieceWithRightTeam(1, 1), Piece{Position: Tuple{1, 2}, Team: " ", King: false}, Piece{Position: Tuple{1, 3}, Team: " ", King: false},
			Piece{Position: Tuple{1, 4}, Team: " ", King: false}, Piece{Position: Tuple{1, 5}, Team: " ", King: false}, Piece{Position: Tuple{1, 6}, Team: "G", King: false}, Piece{Position: Tuple{1, 7}, Team: " ", King: false}},
		// Row 2
		{GetPieceWithRightTeam(2, 0), Piece{Position: Tuple{2, 1}, Team: " ", King: false}, GetPieceWithRightTeam(2, 2), Piece{Position: Tuple{2, 3}, Team: " ", King: false},
			Piece{Position: Tuple{2, 4}, Team: " ", King: false}, Piece{Position: Tuple{2, 5}, Team: " ", King: false}, Piece{Position: Tuple{2, 6}, Team: " ", King: false}, Piece{Position: Tuple{2, 7}, Team: " ", King: false}},
		// Row 3
		{Piece{Position: Tuple{3, 0}, Team: " ", King: false}, Piece{Position: Tuple{3, 1}, Team: " ", King: false}, GetPieceWithRightTeam(3, 2), Piece{Position: Tuple{3, 3}, Team: " ", King: false},
			Piece{Position: Tuple{3, 4}, Team: " ", King: false}, Piece{Position: Tuple{3, 5}, Team: " ", King: false}, Piece{Position: Tuple{3, 6}, Team: " ", King: false}, Piece{Position: Tuple{3, 7}, Team: " ", King: false}},
		// Row 4
		{Piece{Position: Tuple{4, 0}, Team: " ", King: false}, GetPieceWithRightTeam(4, 1), Piece{Position: Tuple{4, 2}, Team: " ", King: false}, Piece{Position: Tuple{4, 3}, Team: " ", King: false},
			Piece{Position: Tuple{4, 4}, Team: " ", King: false}, Piece{Position: Tuple{4, 5}, Team: " ", King: false}, Piece{Position: Tuple{4, 6}, Team: " ", King: false}, Piece{Position: Tuple{4, 7}, Team: " ", King: false}},
		// Row 5
		{GetPieceWithRightTeam(5, 0), Piece{Position: Tuple{5, 1}, Team: " ", King: false}, Piece{Position: Tuple{5, 2}, Team: " ", King: false}, GetPieceWithRightTeam(5, 3),
			Piece{Position: Tuple{5, 4}, Team: " ", King: false}, GetPieceWithRightTeam(5, 5), Piece{Position: Tuple{5, 6}, Team: " ", King: false}, GetPieceWithRightTeam(5, 7)},
		// Row 6
		{Piece{Position: Tuple{6, 0}, Team: "R", King: false}, GetPieceWithRightTeam(6, 1), Piece{Position: Tuple{6, 2}, Team: " ", King: false}, Piece{Position: Tuple{6, 3}, Team: " ", King: false},
			Piece{Position: Tuple{6, 4}, Team: " ", King: false}, Piece{Position: Tuple{6, 5}, Team: " ", King: false}, Piece{Position: Tuple{6, 6}, Team: " ", King: false}, Piece{Position: Tuple{6, 7}, Team: " ", King: false}},
		// Row 7
		{GetPieceWithRightTeam(7, 0), Piece{Position: Tuple{7, 1}, Team: " ", King: false}, GetPieceWithRightTeam(7, 2), Piece{Position: Tuple{7, 3}, Team: " ", King: false},
			GetPieceWithRightTeam(7, 4), GetPieceWithRightTeam(7, 5), GetPieceWithRightTeam(7, 6), GetPieceWithRightTeam(7, 7)},
	}

	return GameBoard{
		GameBoard:  gameBoard,
		CurrPlayer: CurrPlayer,
	}
}
