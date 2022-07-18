package chessgo

type Color uint8

var White = Color(0)
var Black = Color(1)

func (c Color) String() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	}
	panic("Invalid Color")
}

func (c Color) Opponent() Color {
	switch c {
	case White:
		return Black
	case Black:
		return White
	}
	panic("Invalid Color")
}
