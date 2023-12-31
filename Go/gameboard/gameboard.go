package gameboard

import (
	"fmt"
	"strconv"

	"github.com/TwiN/go-color"
	"github.com/samber/lo"
)

type PieceFunction func(row, column int) Piece

type Tuple struct {
	Row, Column int
}

func (t Tuple) String() string {
	return fmt.Sprintf("(%d, %d)", t.Row, t.Column)
}

type Action struct {
	Start, End Tuple
}

func (a Action) String() string {
	var result string

	result += fmt.Sprintf("Startposition Spielstein: "+color.InGreen("%s   "), a.Start)
	result += fmt.Sprintf("Endposition: "+color.InGreen("%s\n"), a.End)
	return result

}

type Piece struct {
	Position          Tuple
	Team              string
	King              bool
	PossiblePositions []Tuple
}

func (p Piece) String() string {
	teamString := p.Team + " "

	if p.King {
		teamString = p.Team + "K"
	}
	if p.Team == "R" {
		return fmt.Sprintf("%s|", color.InRed(teamString))
	} else if p.Team == "G" {
		return fmt.Sprintf("%s|", color.InGreen(teamString))
	} else {
		return fmt.Sprintf("%s|", teamString)
	}

}

type GameBoard struct {
	GameBoard  [][]Piece
	CurrPlayer string
}

func (g GameBoard) String() string {
	var result string
	green := g.CurrPlayer == "G"
	for i, row := range g.GameBoard {
		if green {

			if i == 0 || i == 7 {
				if i == 0 {

					result += color.InGreen("\n\n      0   1   2   3   4   5   6   7\n")
					result += color.InGreen("   ---------------------------------\n")
					result += color.InGreen(strconv.Itoa(i)+"|   ") + fmt.Sprint(row) + color.InGreen("   |"+strconv.Itoa(i)+"\n")

				} else {
					result += color.InGreen(strconv.Itoa(i)+"|   ") + fmt.Sprint(row) + color.InGreen("   |"+strconv.Itoa(i)) + "\n"
					result += color.InGreen("   ---------------------------------\n")
					result += color.InGreen("      0   1   2   3   4   5   6   7\n")
				}
			} else {
				result += color.InGreen(strconv.Itoa(i)+"|   ") + fmt.Sprint(row) + color.InGreen("   |"+strconv.Itoa(i)+"\n")
			}

		} else {
			if i == 0 || i == 7 {
				if i == 0 {

					result += color.InRed("\n\n      0   1   2   3   4   5   6   7\n")
					result += color.InRed("   ---------------------------------\n")
					result += color.InRed(strconv.Itoa(i)+"|   ") + fmt.Sprint(row) + color.InRed("   |"+strconv.Itoa(i)+"\n")

				} else {
					result += color.InRed(strconv.Itoa(i)+"|   ") + fmt.Sprint(row) + color.InRed("   |"+strconv.Itoa(i)) + "\n"
					result += color.InRed("   ---------------------------------\n")
					result += color.InRed("      0   1   2   3   4   5   6   7\n")
				}
			} else {
				result += color.InRed(strconv.Itoa(i)+"|   ") + fmt.Sprint(row) + color.InRed("   |"+strconv.Itoa(i)+"\n")
			}

		}
	}
	return result

}

func GetPieceWithRightTeam(row, column int) Piece {
	team := " "
	if (row+column)%2 == 0 && row < 3 {
		team = "R"
	} else if (row+column)%2 == 0 && row > 4 {
		team = "G"
	}
	return Piece{
		Position: Tuple{row, column},
		Team:     team,
		King:     false,
	}
}

func generateRow(get_piece_with_team_function PieceFunction, row int) []Piece {
	pieces := mapRange(8, func(col int) Piece {
		return get_piece_with_team_function(row, col)
	})
	return pieces
}

func (gameboard GameBoard) generateBoard(get_piece_with_team_function PieceFunction) [][]Piece {

	board := lo.Map(gameboard.GameBoard, func(item []Piece, row int) []Piece {
		return generateRow(get_piece_with_team_function, row)
	})

	return board
}

func mapRange(count int, f func(int) Piece) []Piece {
	result := make([]Piece, count)
	for i := range result {
		result[i] = f(i)
	}
	return result
}

func GetInitialGameBoard() (GameBoard, error) {
	gameBoard := GameBoard{GameBoard: make([][]Piece, 8)}.generateBoard(GetPieceWithRightTeam)

	return GameBoard{
		GameBoard:  gameBoard,
		CurrPlayer: "G",
	}, nil

}
func (gameboard GameBoard) MakeNewGameBoardAfterMove(action Action, smash bool, king bool) GameBoard {
	piece := gameboard.GameBoard[action.Start.Row][action.Start.Column]

	piece_function := func(piece Piece, smash bool, king bool) func(row int, column int) Piece {
		oldrow := piece.Position.Row
		oldcolumn := piece.Position.Column
		return func(row int, column int) Piece {
			team := " "
			king_type := false
			if row == oldrow && column == oldcolumn {
				team = " "
				king_type = false
			} else if row == action.End.Row && column == action.End.Column {
				team = piece.Team
				king_type = king
			} else if smash && (row == int((oldrow+action.End.Row)/2) && column == int((oldcolumn+action.End.Column)/2)) {
				team = " "
				king_type = false
			} else {
				return gameboard.GameBoard[row][column]
			}

			return Piece{
				Position: Tuple{row, column},
				Team:     team,
				King:     king_type,
			}
		}
	}
	function_with_piece := piece_function(piece, smash, king)
	newGameBoard := gameboard.generateBoard(function_with_piece)
	return GameBoard{GameBoard: newGameBoard, CurrPlayer: opposite(piece.Team, smash)}
}

func opposite(team string, smash bool) string {
	if smash {
		return team
	} else if team == "G" {
		return "R"
	} else {
		return "G"
	}
}
