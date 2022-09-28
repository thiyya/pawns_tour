package main

import (
	"log"
	brd "pawns-tour/board"
	"pawns-tour/global"
)

func main() {
	board := brd.New()
	firstMove := 1
	board.Visit(global.STARTING_POSITION_X, global.STARTING_POSITION_Y, firstMove)
	if board.Solve(firstMove+1, global.STARTING_POSITION_X, global.STARTING_POSITION_Y) {
		board.PrintBoard()
	} else {
		err := global.NewError(global.ErrNoSolutionExist, "There is no solution for this board!")
		log.Fatal(err)
	}
}
