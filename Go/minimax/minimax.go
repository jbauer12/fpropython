package minimax

import (
	"checkers/packages/gameboard"
	"checkers/packages/possible_moves"

	"github.com/samber/lo"
)

const (
	positionalWeight  = 0.5
	pieceCountWeight  = 5.0
	kingedPieceWeight = 3.0
)

type BestScore struct {
	Action gameboard.Action
	Score  float64
}

func opposite(team string, smash bool) string {
	if smash {
		return team
	}
	if team == "G" {
		return "R"
	}
	return "G"
}

func Terminal(gameBoard gameboard.GameBoard) bool {
	pieces := lo.Flatten(gameBoard.GameBoard)
	allLeftPiecesRed := lo.CountValuesBy(pieces, func(piece gameboard.Piece) bool {
		return piece.Team == "R" || piece.Team == " "
	})
	allLeftPiecesGreen := lo.CountValuesBy(pieces, func(piece gameboard.Piece) bool {
		return piece.Team == "G" || piece.Team == " "
	})
	moves := possible_moves.GetAllPossibleMovesForTeam(gameBoard, gameBoard.CurrPlayer)
	return allLeftPiecesRed[false] == 0 || allLeftPiecesGreen[false] == 0 || len(moves) == 0
}

func evaluateHeuristicValue(gameBoard gameboard.GameBoard) float64 {
	evalPlayer := func(player string) float64 {
		valueFromPiece := func(piece gameboard.Piece) float64 {
			king := 0.0
			if piece.King {
				king = 1.0
			}
			return pieceCountWeight + kingedPieceWeight*king + float64(piece.Position.Row)*positionalWeight
		}

		filteredPieces := lo.Filter(lo.Flatten(gameBoard.GameBoard), func(piece gameboard.Piece, _ int) bool {
			return player == piece.Team
		})

		return lo.Reduce(filteredPieces, func(agg float64, item gameboard.Piece, _ int) float64 {
			return agg + valueFromPiece(item)
		}, 0)
	}

	playerScore := evalPlayer(gameBoard.CurrPlayer)
	opponentScore := evalPlayer(opposite(gameBoard.CurrPlayer, false))

	return opponentScore - playerScore
}

func possibleActions(gameBoard gameboard.GameBoard, team string) []gameboard.Action {
	possibleMovePieces := possible_moves.GetAllPossibleMovesForTeam(gameBoard, team)
	return possible_moves.GetActionsFromPossibleMoves(gameBoard, possibleMovePieces)
}

func performAction(gameBoard gameboard.GameBoard, action gameboard.Action) gameboard.GameBoard {
	if !Terminal(gameBoard) {
		return possible_moves.MakeMove(gameBoard, action)
	}
	return gameBoard
}

func Minimax(state gameboard.GameBoard, depth int, player string) BestScore {
	bestScore := BestScore{Action: gameboard.Action{}, Score: -1000000}
	if player == "G" {
		bestScore.Score = 1000000
	}

	if depth == 0 || Terminal(state) {
		score := evaluateHeuristicValue(state)
		return BestScore{Action: gameboard.Action{}, Score: score}
	}

	for _, action := range possibleActions(state, player) {
		result := performAction(state, action)
		value := Minimax(result, depth-1, opposite(player, result.CurrPlayer == player))

		if player == "R" && value.Score > bestScore.Score {
			bestScore = BestScore{Action: action, Score: value.Score}
		} else if player == "G" && value.Score < bestScore.Score {
			bestScore = BestScore{Action: action, Score: value.Score}
		}
	}
	return bestScore
}
