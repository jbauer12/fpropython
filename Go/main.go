package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"checkers/packages/gameboard"
	"checkers/packages/minimax"
	"checkers/packages/possible_moves"

	"github.com/TwiN/go-color"
)

func makeArtificialMove(gameBoard gameboard.GameBoard) (gameboard.GameBoard, error) {
	action := minimax.Minimax(gameBoard, 5, gameBoard.CurrPlayer)
	gameBoard = possible_moves.Make_move(gameBoard, action.Action)
	return gameBoard, nil
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
		return 0, fmt.Errorf("Sie müssen einen Integer aus der obigen Liste wählen")
	}
	return input, nil
}

func makeUserMove(gameBoard gameboard.GameBoard) (gameboard.GameBoard, error) {
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
	return gameBoard, nil
}

func displayWelcomeMessage() {
	fmt.Print(color.InGreen("\n\n            Willkommen bei Checkers \n"))
}

func displayBoard(gameBoard gameboard.GameBoard) {
	var message string
	if gameBoard.CurrPlayer == "R" {
		message = color.InRed("\n            Farbe Rot war am Zug \n")
	} else {
		message = color.InGreen("\n            Sie sind dran! \n \n")
	}
	fmt.Print(gameBoard)
	fmt.Print(message)
}

func playGame() {
	gameBoard, err := gameboard.GetInitialGameBoard()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	displayWelcomeMessage()
	for !minimax.Terminal(gameBoard) {
		if gameBoard.CurrPlayer == "R" {
			gameBoard, err = makeArtificialMove(gameBoard)
			showGameBoard := gameboard.GameBoard{GameBoard: gameBoard.GameBoard, CurrPlayer: "R"}
			displayBoard(showGameBoard)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		} else {
			displayBoard(gameBoard)
			gameBoard, err = makeUserMove(gameBoard)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
	}
}

func main() {
	playGame()
}
