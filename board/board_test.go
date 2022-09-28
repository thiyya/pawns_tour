package brd

import (
	. "github.com/smartystreets/goconvey/convey"
	"pawns-tour/global"
	"testing"
)

func TestNewBoard(t *testing.T) {
	Convey("New Board", t, func(c C) {
		board := New()
		c.So(board.tryCounter, ShouldEqual, 0)
		c.So(board.status[0][0].Visited, ShouldEqual, false)
		c.So(board.status[0][0].MoveCounter, ShouldEqual, 0)
	})

	Convey("Solve", t, func(c C) {
		board := New()
		res := board.Solve(1, 0, 0)
		c.So(res, ShouldEqual, true)
		c.So(board.tryCounter, ShouldBeGreaterThan, global.BOARD_SIZE*global.BOARD_SIZE)
		c.So(board.status[3][5].Visited, ShouldEqual, true)
		c.So(board.status[3][5].MoveCounter, ShouldBeGreaterThan, 0)
	})

	Convey("Visit", t, func(c C) {
		board := New()
		board.Visit(3, 5, 11)
		c.So(board.status[3][5].Visited, ShouldEqual, true)
		c.So(board.status[3][5].MoveCounter, ShouldEqual, 11)
	})

	Convey("IsValidMove", t, func(c C) {
		board := New()
		board.Visit(3, 5, 11)
		res := board.IsValidMove(3, 5)
		c.So(res, ShouldEqual, false)
		res = board.IsValidMove(-1, 5)
		c.So(res, ShouldEqual, false)
		res = board.IsValidMove(3, -5)
		c.So(res, ShouldEqual, false)
		res = board.IsValidMove(3, 6)
		c.So(res, ShouldEqual, true)
	})

	Convey("IsValidMove", t, func(c C) {
		board := New()
		board.Visit(3, 5, 11)
		res := board.IsValidMove(3, 5)
		c.So(res, ShouldEqual, false)
		res = board.IsValidMove(-1, 5)
		c.So(res, ShouldEqual, false)
		res = board.IsValidMove(3, -5)
		c.So(res, ShouldEqual, false)
		res = board.IsValidMove(3, 6)
		c.So(res, ShouldEqual, true)
	})

	Convey("Clear", t, func(c C) {
		board := New()
		board.Visit(3, 5, 11)
		res := board.IsValidMove(3, 5)
		c.So(res, ShouldEqual, false)

		board.Clear(3, 5)
		res = board.IsValidMove(3, 5)
		c.So(res, ShouldEqual, true)
	})

	Convey("Clear", t, func(c C) {
		board := New()
		board.SetDirection(3, 5, "ok")
		c.So(board.status[3][5].NextDirection, ShouldEqual, "ok")
	})
}
