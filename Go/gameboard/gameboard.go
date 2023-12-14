package gameboard

import "fmt"

type Tuple struct {
	X, Y int
}

type Piece struct {
	Position          Tuple
	Team              string
	King              bool
	PossiblePositions []Tuple
}

func (p Piece) String() string {
	return fmt.Sprintf("%s|", p.Team)
}

type PieceWithPositions struct {
	Piece
	PossiblePositions []Tuple
}

type GameBoard struct {
	GameBoard  [][]Piece
	CurrPlayer string
}

func (g GameBoard) String() string {
	var result string
	for _, row := range g.GameBoard {
		result += " " + fmt.Sprint(row) + "\n"
	}
	return result
}

func getPieceWithRightTeam(row, column int) Piece {
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

func generateRow(row int) ([]Piece, error) {
	pieces := mapRange(8, func(col int) Piece {
		return getPieceWithRightTeam(row, col)
	})
	return pieces, nil
}

func generateBoard() ([][]Piece, error) {
	board := make([][]Piece, 8)
	for i := range board {
		row, err := generateRow(i)
		if err != nil {
			return nil, err
		}
		board[i] = row
	}
	return board, nil
}

func mapRange(count int, f func(int) Piece) []Piece {
	result := make([]Piece, count)
	for i := range result {
		result[i] = f(i)
	}
	return result
}

func GetInitialGameBoard() (GameBoard, error) {
	gameBoard, err := generateBoard()
	if err != nil {
		return GameBoard{}, err
	}

	return GameBoard{
		GameBoard:  gameBoard,
		CurrPlayer: "R",
	}, nil
}
