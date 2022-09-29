package brd

import (
	"fmt"
	"pawns-tour/global"
)

type Board struct {
	status     [][]Point
	tryCounter int
}

type Point struct {
	Visited       bool
	MoveCounter   int
	NextDirection string
}

type Coordinates struct {
	X         int
	Y         int
	Direction string
}

var movesToTry = [8]Coordinates{
	{3, 0, "↓"},
	{-3, 0, "↑"},
	{0, 3, "→"},
	{0, -3, "←"},

	{2, 2, "↘"},
	{2, -2, "↙"},
	{-2, 2, "↗"},
	{-2, -2, "↖"},
}

func New() *Board {
	board := make([][]Point, global.BOARD_SIZE)
	for i := range board {
		board[i] = make([]Point, global.BOARD_SIZE)
	}
	return &Board{board, 1}
}

func (b *Board) Solve(move int, x int, y int) bool {
	// to see all moves :
	// b.PrintBoard()
	if move == global.BOARD_SIZE*global.BOARD_SIZE+1 {
		b.SetDirection(x, y, "*")
		return true
	}
	b.tryCounter += 1
	for _, coordinate := range movesToTry {
		nextX := x + coordinate.X
		nextY := y + coordinate.Y
		direction := coordinate.Direction
		if b.IsValidMove(nextX, nextY) {
			b.Visit(nextX, nextY, move)
			b.SetDirection(x, y, direction)
			if b.Solve(move+1, nextX, nextY) {
				return true
			}
			b.Clear(nextX, nextY)
		}
	}
	return false
}

func (b *Board) IsValidMove(x int, y int) bool {
	if x < 0 || x >= global.BOARD_SIZE || y < 0 || y >= global.BOARD_SIZE || b.status[x][y].Visited {

		return false
	}

	return true
}

func (b *Board) Visit(x int, y int, move int) {
	b.status[x][y] = Point{
		Visited:     true,
		MoveCounter: move,
	}
}

func (b *Board) Clear(x int, y int) {
	b.status[x][y] = Point{
		Visited:       false,
		MoveCounter:   0,
		NextDirection: "",
	}
}

func (b *Board) PrintBoard() {
	fmt.Printf("----------%dth try----------\n", b.tryCounter)
	for i := 0; i < global.BOARD_SIZE; i++ {
		for j := 0; j < global.BOARD_SIZE; j++ {
			fmt.Printf("%02d%s ", b.status[i][j].MoveCounter, b.status[i][j].NextDirection)
		}
		fmt.Println()
	}
}

func (b *Board) SetDirection(x int, y int, direction string) {
	b.status[x][y].NextDirection = direction
}
