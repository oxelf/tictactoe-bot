package main

import (
	"fmt"
	"math/rand"
)

func (board *Board) CalculateNextMove() int {
	possiblePositions := []int{}
	for i := 0; i < len(board.Board); i++ {
		if board.Board[i] == 0 {
			possiblePositions = append(possiblePositions, i)
		}
	}
	fmt.Println("possible positions: ", possiblePositions)
	if len(possiblePositions) == 0 {
		return 0
	}
	return possiblePositions[rand.Intn((len(possiblePositions)))]
}
