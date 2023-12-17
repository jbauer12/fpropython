package main

import (
	"checkers/packages/gameboard"
	"checkers/packages/minimax"
	"fmt"
)

func main() {
	gameBoard, err := gameboard.GetInitialGameBoard()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	minimax.Terminal(gameBoard)
	flatSlice := minimax.ValueFrom(gameBoard)
	// Print the flattened slice
	fmt.Println(flatSlice)
}
