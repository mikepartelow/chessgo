package chessgo

import "fmt"

type Color uint8

var White = Color(0)
var Black = Color(1)

func (c *Color) String() string {
	switch *c {
	case White:
		return "White"
	case Black:
		return "Black"
	}
	panic("Invalid Color")
}

func (c *Color) Opponent() Color {
	switch *c {
	case White:
		return Black
	case Black:
		return White
	}
	panic("Invalid Color")
}

func ColorOf(p Piece) Color {
	switch p {
	case BlackRook,
		BlackKnight,
		BlackBishop,
		BlackQueen,
		BlackKing,
		BlackPawn:
		return Black

	case WhiteRook,
		WhiteKnight,
		WhiteBishop,
		WhiteQueen,
		WhiteKing,
		WhitePawn:
		return White
	}

	panic(fmt.Sprintf("No color for %c", p))
}
