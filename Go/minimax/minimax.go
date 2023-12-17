package minimax

import (
	"checkers/packages/gameboard"
	"checkers/packages/possible_moves"
	"fmt"

	"github.com/samber/lo"
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

	allPiecesOfTeam := func(team string) bool {
		return lo.ContainsBy(gameBoard.GameBoard, func(row []gameboard.Piece) bool {
			return lo.ContainsBy(row, func(piece gameboard.Piece) bool {
				return piece.Team == team || piece.Team == " "
			})
		})
	}

	allLeftPiecesRed := allPiecesOfTeam("R")
	fmt.Println(allLeftPiecesRed)

	allLeftPiecesGreen := allPiecesOfTeam("G")

	return allLeftPiecesRed || allLeftPiecesGreen
}

func ValueFrom(gameBoard gameboard.GameBoard) float64 {
	positionalWeight := 0.5
	pieceCountWeight := 5.0
	kingedPieceWeight := 3.0

	valueFromPiece := func(piece gameboard.Piece) float64 {
		king := 0.0
		if piece.King {
			king = 1.0
		}
		return pieceCountWeight + kingedPieceWeight*king + float64(piece.Position.Row)*positionalWeight
	}
	filter_function := func(pieces []gameboard.Piece, currPlayer string) []gameboard.Piece {
		return lo.Filter(pieces, func(piece gameboard.Piece, _ int) bool {
			return currPlayer == piece.Team
		})
	}
	reduce_function := func(pieces []gameboard.Piece) float64 {
		return lo.Reduce(pieces, func(agg float64, item gameboard.Piece, _ int) float64 {
			return agg + valueFromPiece(item)
		}, 0)
	}

	flattenBoard := lo.Flatten(gameBoard.GameBoard)
	p1 := filter_function(flattenBoard, gameBoard.CurrPlayer)
	playerScore := reduce_function(p1)
	//TODO Smash funktion einbauen!
	opponentScore := reduce_function(filter_function(flattenBoard, opposite(gameBoard.CurrPlayer, false)))

	return opponentScore - playerScore
}
func Player(gameBoard gameboard.GameBoard) string {
	return gameBoard.CurrPlayer
}
func Actions(gameBoard gameboard.GameBoard, team string) []gameboard.Action {
	possible_moves := possible_moves.Get_all_possible_moves_for_team(gameBoard, team)
	return possible_moves
}
func Result(gameBoard gameboard.GameBoard, action gameboard.Action) gameboard.GameBoard {
	if !Terminal(gameBoard) {
		return possible_moves.Make_move(gameBoard, action)
	}
	return gameBoard
}
func Minimax(state gameboard.GameBoard, depth int, player string) BestScore {

	bestScore := BestScore{gameboard.Action{}, -1000000}
	if player == "G" {
		bestScore.Score = 1000000
	}

	if depth == 0 || Terminal(state) {
		score := ValueFrom(state)
		return BestScore{gameboard.Action{}, score}
	}

	for _, action := range Actions(state, player) {
		result1 := Result(state, action)
		value := Minimax(result1, depth-1, opposite(player, result1.CurrPlayer == player))

		if player == "R" && value.Score > bestScore.Score {
			bestScore = BestScore{action, value.Score}
		} else if player == "G" && value.Score < bestScore.Score {
			bestScore = BestScore{action, value.Score}
		}
	}
	return bestScore
}
