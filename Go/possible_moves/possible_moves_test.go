package possible_moves

import (
	"checkers/packages/gameboard"
	. "checkers/packages/gameboard"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfPositionOutOfBounds(t *testing.T) {
	assert.True(t, checkIfPositionOutOfBounds(Tuple{Row: 3, Column: 5}))
	assert.False(t, checkIfPositionOutOfBounds(Tuple{Row: -1, Column: 5}))
	assert.False(t, checkIfPositionOutOfBounds(Tuple{Row: 3, Column: 10}))
}

func TestIsPieceThereOnRowColumn(t *testing.T) {
	piece := Piece{Position: Tuple{Row: 2, Column: 3}, Team: "R"}
	assert.True(t, isPieceThereOnRowColumn(piece))
	piece = Piece{Position: Tuple{Row: -1, Column: 3}, Team: "R"}
	assert.False(t, isPieceThereOnRowColumn(piece))
	emptyPiece := Piece{Position: Tuple{Row: 2, Column: 3}, Team: " "}
	assert.False(t, isPieceThereOnRowColumn(emptyPiece))
}

func TestGetPossibleMoveVectors(t *testing.T) {

	emptyPiece := Piece{Position: Tuple{Row: 2, Column: 3}, Team: " "}
	vectors := getPossibleMoveVectors(emptyPiece, false)
	assert.Empty(t, vectors, "Empty space should have no possible move vectors")

	regularPiece := Piece{Position: Tuple{Row: 2, Column: 3}, Team: "R"}
	vectors = getPossibleMoveVectors(regularPiece, true)
	assert.ElementsMatch(t, []Tuple{{Row: 1, Column: -1}, {Row: 1, Column: 1}}, vectors,
		"Regular piece of team 'R' should have correct possible move vectors")

	greenPiece := Piece{Position: Tuple{Row: 2, Column: 3}, Team: "G"}
	vectors = getPossibleMoveVectors(greenPiece, true)
	assert.ElementsMatch(t, []Tuple{{Row: -1, Column: -1}, {Row: -1, Column: 1}}, vectors,
		"Regular piece of team 'G' should have correct possible move vectors")

	kingPiece := Piece{Position: Tuple{Row: 2, Column: 3}, Team: "R", King: true}
	vectors = getPossibleMoveVectors(kingPiece, true)
	assert.ElementsMatch(t, []Tuple{{Row: -1, Column: -1}, {Row: -1, Column: 1}, {Row: 1, Column: -1}, {Row: 1, Column: 1}},
		vectors, "King piece should have correct possible move vectors")
}
func TestFilterWithoutSmash(t *testing.T) {
	gameBoard, err := GetInitialGameBoard()
	assert.Nil(t, err, "GetInitialGameBoard should not return an error")

	piece := gameBoard.GameBoard[5][1]
	vector := Tuple{Row: 4, Column: 0}
	result := filterWithoutSmash(gameBoard, piece, vector)
	assert.True(t, result, "Valid move without smash should return true")

	piece = gameBoard.GameBoard[2][0]
	vector = Tuple{Row: 1, Column: -1}
	result = filterWithoutSmash(gameBoard, piece, vector)
	assert.False(t, result, "Invalid move with out-of-bounds position should return false")

	// Test for an invalid move with a piece in the destination position
	piece = gameBoard.GameBoard[2][2]
	gameBoard.GameBoard[3][3] = Piece{Position: Tuple{Row: 3, Column: 3}, Team: "R", King: false}
	vector = Tuple{Row: 1, Column: 1}
	result = filterWithoutSmash(gameBoard, piece, vector)
	assert.False(t, result, "Invalid move with a piece in the destination position should return false")
	gameBoard = GenerateBoardWithSmash()
	vector = Tuple{Row: 1, Column: 1}
	filterWithoutSmash(gameBoard, piece, vector)
	assert.False(t, result, "There is a smash move")
}
func TestFilterWithSmash(t *testing.T) {
	gameBoard := GenerateBoardWithSmash()
	fmt.Print(gameBoard)
	piece := gameBoard.GameBoard[4][2]
	vector := Tuple{Row: 3, Column: 1}
	result := filterWithSmash(gameBoard, piece, vector)
	assert.False(t, result, "Invalid move with smash should return true")
	vector = Tuple{Row: 6, Column: 4}
	result = filterWithSmash(gameBoard, piece, vector)
	assert.True(t, result, "There is a smash move")

	piece = gameBoard.GameBoard[2][2]
	vector = Tuple{Row: 1, Column: 1}
	result = filterWithSmash(gameBoard, piece, vector)
	assert.False(t, result, "No smash move")

	piece = gameBoard.GameBoard[2][2]
	vector = Tuple{Row: 1, Column: -1}
	result = filterWithSmash(gameBoard, piece, vector)
	assert.False(t, result, "No smash move")

}
func TestGetAllPossibleMovesForTeam(t *testing.T) {
	gameBoard := GenerateBoardWithSmash()
	fmt.Print(gameBoard)
	team := "R"

	expectedMoves := []Piece{
		Piece{Position: Tuple{Row: 4, Column: 2}, Team: "R", King: false, PossiblePositions: []Tuple{{Row: 6, Column: 4}}},
	}

	moves := GetAllPossibleMovesForTeam(gameBoard, team)
	fmt.Print(moves)
	fmt.Print(expectedMoves)
	assert.ElementsMatch(t, expectedMoves, moves, "Incorrect possible moves for team 'R'")
}

func GenerateBoardWithSmash() GameBoard {
	gameBoard := [][]Piece{
		// Row 0
		{Piece{Position: Tuple{Row: 0, Column: 0}, Team: " ", King: false}, Piece{Position: Tuple{Row: 0, Column: 1}, Team: " ", King: false}, GetPieceWithRightTeam(0, 2), Piece{Position: Tuple{Row: 0, Column: 3}, Team: " ", King: false},
			GetPieceWithRightTeam(0, 4), Piece{Position: Tuple{Row: 0, Column: 5}, Team: " ", King: false}, Piece{Position: Tuple{Row: 0, Column: 6}, Team: " ", King: false}, Piece{Position: Tuple{Row: 0, Column: 7}, Team: " ", King: false}},
		// Row 1
		{Piece{Position: Tuple{Row: 1, Column: 0}, Team: " ", King: false}, GetPieceWithRightTeam(1, 1), Piece{Position: Tuple{Row: 1, Column: 2}, Team: " ", King: false}, Piece{Position: Tuple{Row: 1, Column: 3}, Team: " ", King: false},
			Piece{Position: Tuple{Row: 1, Column: 4}, Team: " ", King: false}, Piece{Position: Tuple{Row: 1, Column: 5}, Team: " ", King: false}, Piece{Position: Tuple{Row: 1, Column: 6}, Team: " ", King: false}, Piece{Position: Tuple{Row: 1, Column: 7}, Team: " ", King: false}},
		// Row 2
		{GetPieceWithRightTeam(2, 0), Piece{Position: Tuple{Row: 2, Column: 1}, Team: " ", King: false}, GetPieceWithRightTeam(2, 2), Piece{Position: Tuple{Row: 2, Column: 3}, Team: " ", King: false},
			Piece{Position: Tuple{Row: 2, Column: 4}, Team: " ", King: false}, Piece{Position: Tuple{Row: 2, Column: 5}, Team: " ", King: false}, Piece{Position: Tuple{Row: 2, Column: 6}, Team: " ", King: false}, Piece{Position: Tuple{Row: 2, Column: 7}, Team: " ", King: false}},
		// Row 3
		{Piece{Position: Tuple{Row: 3, Column: 0}, Team: " ", King: false}, Piece{Position: Tuple{Row: 3, Column: 1}, Team: " ", King: false}, GetPieceWithRightTeam(3, 2), Piece{Position: Tuple{Row: 3, Column: 3}, Team: " ", King: false},
			Piece{Position: Tuple{Row: 3, Column: 4}, Team: " ", King: false}, Piece{Position: Tuple{Row: 3, Column: 5}, Team: " ", King: false}, Piece{Position: Tuple{Row: 3, Column: 6}, Team: " ", King: false}, Piece{Position: Tuple{Row: 3, Column: 7}, Team: " ", King: false}},
		// Row 4
		{Piece{Position: Tuple{Row: 4, Column: 0}, Team: " ", King: false}, GetPieceWithRightTeam(4, 1), Piece{Position: Tuple{Row: 4, Column: 2}, Team: "R", King: false}, Piece{Position: Tuple{Row: 4, Column: 3}, Team: " ", King: false},
			Piece{Position: Tuple{Row: 4, Column: 4}, Team: " ", King: false}, Piece{Position: Tuple{Row: 4, Column: 5}, Team: " ", King: false}, Piece{Position: Tuple{Row: 4, Column: 6}, Team: " ", King: false}, Piece{Position: Tuple{Row: 4, Column: 7}, Team: " ", King: false}},
		// Row 5
		{GetPieceWithRightTeam(5, 0), Piece{Position: Tuple{Row: 5, Column: 1}, Team: " ", King: false}, Piece{Position: Tuple{Row: 5, Column: 2}, Team: " ", King: false}, GetPieceWithRightTeam(5, 3),
			Piece{Position: Tuple{Row: 5, Column: 4}, Team: " ", King: false}, GetPieceWithRightTeam(5, 5), Piece{Position: Tuple{Row: 5, Column: 6}, Team: " ", King: false}, GetPieceWithRightTeam(5, 7)},
		// Row 6
		{Piece{Position: Tuple{Row: 6, Column: 0}, Team: " ", King: false}, GetPieceWithRightTeam(6, 1), Piece{Position: Tuple{Row: 6, Column: 2}, Team: " ", King: false}, Piece{Position: Tuple{Row: 6, Column: 3}, Team: " ", King: false},
			Piece{Position: Tuple{Row: 6, Column: 4}, Team: " ", King: false}, Piece{Position: Tuple{Row: 6, Column: 5}, Team: " ", King: false}, Piece{Position: Tuple{Row: 6, Column: 6}, Team: " ", King: false}, Piece{Position: Tuple{Row: 6, Column: 7}, Team: " ", King: false}},
		// Row 7
		{GetPieceWithRightTeam(7, 0), Piece{Position: Tuple{Row: 7, Column: 1}, Team: " ", King: false}, GetPieceWithRightTeam(7, 2), Piece{Position: Tuple{Row: 7, Column: 3}, Team: " ", King: false},
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
		{Piece{Position: Tuple{0, 0}, Team: " ", King: false}, Piece{Position: Tuple{0, 1}, Team: " ", King: false}, gameboard.GetPieceWithRightTeam(0, 2), Piece{Position: Tuple{0, 3}, Team: " ", King: false},
			gameboard.GetPieceWithRightTeam(0, 4), Piece{Position: Tuple{0, 5}, Team: " ", King: false}, Piece{Position: Tuple{0, 6}, Team: " ", King: false}, Piece{Position: Tuple{0, 7}, Team: " ", King: false}},
		// Row 1
		{Piece{Position: Tuple{1, 0}, Team: " ", King: false}, gameboard.GetPieceWithRightTeam(1, 1), Piece{Position: Tuple{1, 2}, Team: " ", King: false}, Piece{Position: Tuple{1, 3}, Team: " ", King: false},
			Piece{Position: Tuple{1, 4}, Team: " ", King: false}, Piece{Position: Tuple{1, 5}, Team: " ", King: false}, Piece{Position: Tuple{1, 6}, Team: "G", King: false}, Piece{Position: Tuple{1, 7}, Team: " ", King: false}},
		// Row 2
		{gameboard.GetPieceWithRightTeam(2, 0), Piece{Position: Tuple{2, 1}, Team: " ", King: false}, gameboard.GetPieceWithRightTeam(2, 2), Piece{Position: Tuple{2, 3}, Team: " ", King: false},
			Piece{Position: Tuple{2, 4}, Team: " ", King: false}, Piece{Position: Tuple{2, 5}, Team: " ", King: false}, Piece{Position: Tuple{2, 6}, Team: " ", King: false}, Piece{Position: Tuple{2, 7}, Team: " ", King: false}},
		// Row 3
		{Piece{Position: Tuple{3, 0}, Team: " ", King: false}, Piece{Position: Tuple{3, 1}, Team: " ", King: false}, gameboard.GetPieceWithRightTeam(3, 2), Piece{Position: Tuple{3, 3}, Team: " ", King: false},
			Piece{Position: Tuple{3, 4}, Team: " ", King: false}, Piece{Position: Tuple{3, 5}, Team: " ", King: false}, Piece{Position: Tuple{3, 6}, Team: " ", King: false}, Piece{Position: Tuple{3, 7}, Team: " ", King: false}},
		// Row 4
		{Piece{Position: Tuple{4, 0}, Team: " ", King: false}, gameboard.GetPieceWithRightTeam(4, 1), Piece{Position: Tuple{4, 2}, Team: " ", King: false}, Piece{Position: Tuple{4, 3}, Team: " ", King: false},
			Piece{Position: Tuple{4, 4}, Team: " ", King: false}, Piece{Position: Tuple{4, 5}, Team: " ", King: false}, Piece{Position: Tuple{4, 6}, Team: " ", King: false}, Piece{Position: Tuple{4, 7}, Team: " ", King: false}},
		// Row 5
		{gameboard.GetPieceWithRightTeam(5, 0), Piece{Position: Tuple{5, 1}, Team: " ", King: false}, Piece{Position: Tuple{5, 2}, Team: " ", King: false}, gameboard.GetPieceWithRightTeam(5, 3),
			Piece{Position: Tuple{5, 4}, Team: " ", King: false}, gameboard.GetPieceWithRightTeam(5, 5), Piece{Position: Tuple{5, 6}, Team: " ", King: false}, gameboard.GetPieceWithRightTeam(5, 7)},
		// Row 6
		{Piece{Position: Tuple{6, 0}, Team: "R", King: false}, gameboard.GetPieceWithRightTeam(6, 1), Piece{Position: Tuple{6, 2}, Team: " ", King: false}, Piece{Position: Tuple{6, 3}, Team: " ", King: false},
			Piece{Position: Tuple{6, 4}, Team: " ", King: false}, Piece{Position: Tuple{6, 5}, Team: " ", King: false}, Piece{Position: Tuple{6, 6}, Team: " ", King: false}, Piece{Position: Tuple{6, 7}, Team: " ", King: false}},
		// Row 7
		{gameboard.GetPieceWithRightTeam(7, 0), Piece{Position: Tuple{7, 1}, Team: " ", King: false}, gameboard.GetPieceWithRightTeam(7, 2), Piece{Position: Tuple{7, 3}, Team: " ", King: false},
			gameboard.GetPieceWithRightTeam(7, 4), gameboard.GetPieceWithRightTeam(7, 5), gameboard.GetPieceWithRightTeam(7, 6), gameboard.GetPieceWithRightTeam(7, 7)},
	}

	return GameBoard{
		GameBoard:  gameBoard,
		CurrPlayer: CurrPlayer,
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
