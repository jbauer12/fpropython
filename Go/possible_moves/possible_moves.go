package possible_moves

import (
	. "checkers/packages/gameboard"

	"github.com/samber/lo"
)

type PossibleMoveFunction func(gameBoard GameBoard, piece Piece, vectorsWhichPieceCanMove []Tuple) []Tuple

type FilterFunction func(gameBoard GameBoard, piece Piece, vector Tuple) bool

func checkIfPositionOutOfBounds(position Tuple) bool {
	return 0 <= position.Row && position.Row < 8 && 0 <= position.Column && position.Column < 8
}

func isPieceThereOnRowColumn(piece Piece) bool {
	return checkIfPositionOutOfBounds(piece.Position) && piece.Team != " "
}

func GetPossibleMoveVectors(piece Piece, isPieceThere bool) []Tuple {
	vectorsWhichPieceCanMove := []Tuple{}
	if isPieceThere && piece.Team != " " {
		if piece.Team == "R" {
			vectorsWhichPieceCanMove = []Tuple{{Row: 1, Column: -1}, {Row: 1, Column: 1}}
		} else if piece.Team == "G" {
			vectorsWhichPieceCanMove = []Tuple{{Row: -1, Column: -1}, {Row: -1, Column: 1}}
		}
		if piece.King {
			vectorsWhichPieceCanMove = []Tuple{{Row: -1, Column: -1}, {Row: -1, Column: 1}, {Row: 1, Column: -1}, {Row: 1, Column: 1}}
		}
	}
	return vectorsWhichPieceCanMove
}

func FilterWithoutSmash(gameBoard GameBoard, piece Piece, vector Tuple) bool {
	row, column := piece.Position.Row, piece.Position.Column
	rowVector, columnVector := vector.Row-row, vector.Column-column
	position := Tuple{Row: row + rowVector, Column: column + columnVector}
	inBound := checkIfPositionOutOfBounds(position)
	return inBound && !isPieceThereOnRowColumn(gameBoard.GameBoard[position.Row][position.Column])
}

func FilterWithSmash(gameBoard GameBoard, piece Piece, vector Tuple) bool {
	oppositePiecesOnPosition := func(ownPiece, newPositionPiece Piece) bool {
		return ownPiece.Team != " " && newPositionPiece.Team != " " && ownPiece.Team != newPositionPiece.Team
	}

	row, column := piece.Position.Row, piece.Position.Column
	rowVector, columnVector := vector.Row-row, vector.Column-column
	intermediate_row, intermediate_column := row+int(float64(rowVector)*0.5), column+int(float64(columnVector)*0.5)

	inBoundIntermediate := checkIfPositionOutOfBounds(Tuple{Row: intermediate_row, Column: intermediate_column})
	inBoundDestination := checkIfPositionOutOfBounds(Tuple{Row: row + rowVector, Column: column + columnVector})

	if inBoundIntermediate && inBoundDestination {
		opposite := oppositePiecesOnPosition(piece, gameBoard.GameBoard[intermediate_row][intermediate_column])
		withoutPieceOnPosition := !isPieceThereOnRowColumn(gameBoard.GameBoard[row+rowVector][column+columnVector])
		return opposite && withoutPieceOnPosition
	} else {
		return false
	}
}

func GetPossibleMoves(smash bool, filterFunction func(GameBoard, Piece, Tuple) bool) func(gameBoard GameBoard, piece Piece, vectorsWhichPieceCanMove []Tuple) []Tuple {
	return func(gameBoard GameBoard, piece Piece, vectorsWhichPieceCanMove []Tuple) []Tuple {
		row, column := piece.Position.Row, piece.Position.Column
		var possiblePositions []Tuple
		map_function := func(smash bool, row int, column int) []Tuple {
			factor := 1
			if smash {
				factor = 2
			}
			return lo.Map(vectorsWhichPieceCanMove, func(vector Tuple, _ int) Tuple {
				rowVector, columnVector := vector.Row, vector.Column
				return Tuple{Row: row + factor*rowVector, Column: column + factor*columnVector}
			})
		}
		possiblePositions = lo.Filter(map_function(smash, row, column), func(vector Tuple, _ int) bool {
			return filterFunction(gameBoard, piece, vector)
		})

		return possiblePositions

	}
}

func GetAllPossibleMovesForPiece(gameBoard GameBoard, piece Piece, possibleMoveFunction PossibleMoveFunction, filterFunction FilterFunction) Piece {
	vectorsWhichPieceCanMove := GetPossibleMoveVectors(piece, isPieceThereOnRowColumn(piece))
	if len(vectorsWhichPieceCanMove) > 0 {
		possiblePositions := possibleMoveFunction(gameBoard, piece, vectorsWhichPieceCanMove)
		piece.PossiblePositions = possiblePositions
	}
	return piece
}

func GetAllPossibleMovesForTeam(gameBoard GameBoard, team string) []Piece {
	PossibleMoveFunctionWithSmash := GetPossibleMoves(true, FilterWithSmash)
	PossibleMoveFunctionWithoutSmash := GetPossibleMoves(false, FilterWithoutSmash)
	map_function := func(possibleMoveFunction PossibleMoveFunction, piece Piece) Piece {
		smashes := possibleMoveFunction(gameBoard, piece, GetPossibleMoveVectors(piece, isPieceThereOnRowColumn(piece)))
		piece.PossiblePositions = smashes
		return piece
	}
	filter_function := func(piece Piece, _ int) bool {
		return len(piece.PossiblePositions) > 0
	}

	flatten_board := lo.Flatten(gameBoard.GameBoard)
	team_pieces := lo.Filter(flatten_board, func(piece Piece, _ int) bool {
		return piece.Team == team
	})
	team_pieces_with_smash_positions := lo.Map(team_pieces, func(piece Piece, _ int) Piece {
		return map_function(PossibleMoveFunctionWithSmash, piece)
	})
	true_smashes := lo.Filter(team_pieces_with_smash_positions, filter_function)
	if len(true_smashes) == 0 {
		team_pieces_with_possible_positions := lo.Map(team_pieces, func(piece Piece, _ int) Piece {
			return map_function(PossibleMoveFunctionWithoutSmash, piece)
		})
		team_pieces_with_possible_positions = lo.Filter(team_pieces_with_possible_positions, filter_function)
		return team_pieces_with_possible_positions
	} else {
		return true_smashes
	}
}

func Make_move(gameBoard GameBoard, action Action) GameBoard {
	piece := gameBoard.GameBoard[action.Start.Row][action.Start.Column]
	king := isPieceCheckerAfterMove(piece, action)
	smash := isOppositePieceSmashed(action)
	new_game_board := MakeNewGameBoardAfterMove(gameBoard, action, smash, king)
	return new_game_board
}
func GetActionsFromPossibleMoves(gameBoard GameBoard, possibleMoves []Piece) []Action {

	convertPieceToAction := func(piece Piece, possiblePosition Tuple) Action {
		return Action{Start: piece.Position, End: possiblePosition}
	}
	getActionFromOnePiece := func(piece Piece, _ int) []Action {
		return lo.Map(piece.PossiblePositions, func(possiblePosition Tuple, _ int) Action {
			return convertPieceToAction(piece, possiblePosition)
		})
	}
	actions := lo.Map(possibleMoves, getActionFromOnePiece)
	actions_flattened := lo.Flatten(actions)
	return actions_flattened
}

func isOppositePieceSmashed(action Action) bool {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	oldRow, oldColumn := action.Start.Row, action.Start.Column
	row, column := action.End.Row, action.End.Column
	if abs(row-oldRow) == 2 || abs(column-oldColumn) == 2 {
		return true
	}
	return false
}

func isPieceCheckerAfterMove(piece Piece, action Action) bool {
	row := action.End.Row
	if (piece.Team == "R" && row == 7) || (piece.Team == "G" && row == 0) || piece.King {
		return true
	}
	return false
}
