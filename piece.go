package chessgo

import "fmt"

type Piece interface {
	Byte() byte
	Color() Color
}

// White
type WhitePawn struct{}

func (p WhitePawn) Byte() byte   { return 'P' }
func (p WhitePawn) Color() Color { return White }

type WhiteRook struct{}

func (p WhiteRook) Byte() byte   { return 'R' }
func (p WhiteRook) Color() Color { return White }

type WhiteKnight struct{}

func (p WhiteKnight) Byte() byte   { return 'N' }
func (p WhiteKnight) Color() Color { return White }

type WhiteBishop struct{}

func (p WhiteBishop) Byte() byte   { return 'B' }
func (p WhiteBishop) Color() Color { return White }

type WhiteQueen struct{}

func (p WhiteQueen) Byte() byte   { return 'Q' }
func (p WhiteQueen) Color() Color { return White }

type WhiteKing struct{}

func (p WhiteKing) Byte() byte   { return 'K' }
func (p WhiteKing) Color() Color { return White }

// Black
type BlackPawn struct{}

func (p BlackPawn) Byte() byte   { return 'p' }
func (p BlackPawn) Color() Color { return Black }

type BlackRook struct{}

func (p BlackRook) Byte() byte   { return 'r' }
func (p BlackRook) Color() Color { return Black }

type BlackKnight struct{}

func (p BlackKnight) Byte() byte   { return 'n' }
func (p BlackKnight) Color() Color { return Black }

type BlackBishop struct{}

func (p BlackBishop) Byte() byte   { return 'b' }
func (p BlackBishop) Color() Color { return Black }

type BlackQueen struct{}

func (p BlackQueen) Byte() byte   { return 'q' }
func (p BlackQueen) Color() Color { return Black }

type BlackKing struct{}

func (p BlackKing) Byte() byte   { return 'k' }
func (p BlackKing) Color() Color { return Black }

func Pawn(color Color) Piece {
	switch color {
	case White:
		return WhitePawn{}
	case Black:
		return BlackPawn{}
	default:
		panic("invalid Color")
	}
}

func Rook(color Color) Piece {
	switch color {
	case White:
		return WhiteRook{}
	case Black:
		return BlackRook{}
	default:
		panic("invalid Color")
	}
}

func Knight(color Color) Piece {
	switch color {
	case White:
		return WhiteKnight{}
	case Black:
		return BlackKnight{}
	default:
		panic("invalid Color")
	}
}

func Bishop(color Color) Piece {
	switch color {
	case White:
		return WhiteBishop{}
	case Black:
		return BlackBishop{}
	default:
		panic("invalid Color")
	}
}

func Queen(color Color) Piece {
	switch color {
	case White:
		return WhiteQueen{}
	case Black:
		return BlackQueen{}
	default:
		panic("invalid Color")
	}
}

func King(color Color) Piece {
	switch color {
	case White:
		return WhiteKing{}
	case Black:
		return BlackKing{}
	default:
		panic("invalid Color")
	}
}

func PieceFromByte(b byte) Piece {
	switch b {
	case ' ':
		return nil
	case 'P':
		return WhitePawn{}
	case 'R':
		return WhiteRook{}
	case 'N':
		return WhiteKnight{}
	case 'B':
		return WhiteBishop{}
	case 'Q':
		return WhiteQueen{}
	case 'K':
		return WhiteKing{}

	case 'p':
		return BlackPawn{}
	case 'r':
		return BlackRook{}
	case 'n':
		return BlackKnight{}
	case 'b':
		return BlackBishop{}
	case 'q':
		return BlackQueen{}
	case 'k':
		return BlackKing{}
	}

	panic(fmt.Sprintf("unknown piece-byte: %c", b))
}
