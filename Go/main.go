package main

import (
	"checkers/packages/gameboard"
	"fmt"
)

func main() {
	// Example usage
	gameBoard, err := gameboard.GetInitialGameBoard()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(gameBoard)
}
