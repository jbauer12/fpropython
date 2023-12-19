package main

import (
	"bufio"
	"checkers/packages/gameboard"
	"checkers/packages/minimax"
	"checkers/packages/possible_moves"
	"fmt"
	"os"
	"strconv"

	"github.com/TwiN/go-color"
)

func makeArtificialMove(gameBoard gameboard.GameBoard) gameboard.GameBoard {
	action := minimax.Minimax(gameBoard, 7, gameBoard.CurrPlayer)
	gameBoard = possible_moves.Make_move(gameBoard, action.Action)
	return gameBoard
}

func printActions(actions []gameboard.Action) string {
	var result string
	for i, action := range actions {
		result += fmt.Sprintf("Drücke "+color.InGreen("[%d]")+" für: %s", i, action)
	}
	return result
}

func getInputFromUser() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Wählen sie ihren Spielzug: ")
	scanner.Scan()
	inputText := scanner.Text()
	input, err := strconv.Atoi(inputText)
	if err != nil {
		fmt.Println("Sie müssen einen Integer wählen aus der obigen Liste.")
	}

	return input, err
}
func makeUserMove(gameBoard gameboard.GameBoard) gameboard.GameBoard {

	actions := possible_moves.GetActionsFromPossibleMoves(gameBoard, possible_moves.GetAllPossibleMovesForTeam(gameBoard, gameBoard.CurrPlayer))
	fmt.Print(printActions(actions))
	input, err := getInputFromUser()

	if err != nil || input >= len(actions) {
		for err != nil || input >= len(actions) {
			fmt.Printf("Wählen sie als Spielzug eine Zahl zwischen 0 und %d \n", len(actions)-1)
			input, err = getInputFromUser()
		}
	}
	action := actions[input]
	gameBoard = possible_moves.Make_move(gameBoard, action)
	return gameBoard

}

func main() {
	gameBoard, err := gameboard.GetInitialGameBoard()
	fmt.Print(color.InGreen("\n\n            Willkommen bei Checkers \n"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for !minimax.Terminal(gameBoard) {
		if gameBoard.CurrPlayer == "R" {
			gameBoard = makeArtificialMove(gameBoard)
			fmt.Print(color.InRed("\n            Farbe Rot war am Zug \n"))
			showGameBoard := gameboard.GameBoard{GameBoard: gameBoard.GameBoard, CurrPlayer: "R"}
			fmt.Print(showGameBoard)
			fmt.Print("\n\n")

		} else {
			fmt.Print(gameBoard)
			fmt.Print(color.InGreen("\n            Sie sind dran! \n \n"))
			gameBoard = makeUserMove(gameBoard)
		}
	}

}
func main1() {
	gameBoard, err := gameboard.GetInitialGameBoard()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	actions := possible_moves.GetActionsFromPossibleMoves(gameBoard, possible_moves.GetAllPossibleMovesForTeam(gameBoard, gameBoard.CurrPlayer))
	fmt.Print(gameBoard)
	fmt.Print(printActions(actions))
	makeUserMove(gameBoard)
	//gameBoard = possible_moves.Make_move(gameBoard, gameboard.Action{Start: gameboard.Tuple{Row: 5, Column: 1}, End: gameboard.Tuple{Row: 4, Column: 2}})
	gameBoard = possible_moves.Make_move(gameBoard, gameboard.Action{Start: gameboard.Tuple{Row: 2, Column: 0}, End: gameboard.Tuple{Row: 3, Column: 1}})
	fmt.Print(gameBoard)

	actions = possible_moves.GetActionsFromPossibleMoves(gameBoard, possible_moves.GetAllPossibleMovesForTeam(gameBoard, gameBoard.CurrPlayer))

	fmt.Print(gameBoard)
	fmt.Print(printActions(actions))
}
