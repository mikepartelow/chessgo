package chessgo

type Piece rune

var NoPiece Piece = Piece(' ')

var BlackRook Piece = Piece('r')
var BlackKnight Piece = Piece('n')
var BlackBishop Piece = Piece('b')
var BlackQueen Piece = Piece('q')
var BlackKing Piece = Piece('k')
var BlackPawn Piece = Piece('p')

var WhiteRook Piece = Piece('R')
var WhiteKnight Piece = Piece('N')
var WhiteBishop Piece = Piece('B')
var WhiteQueen Piece = Piece('Q')
var WhiteKing Piece = Piece('K')
var WhitePawn Piece = Piece('P')

func Pawn(color Color) Piece {
	switch color {
	case White:
		return WhitePawn
	case Black:
		return BlackPawn
	default:
		panic("Invalid color")
	}
}

func Bishop(color Color) Piece {
	switch color {
	case White:
		return WhiteBishop
	case Black:
		return BlackBishop
	default:
		panic("Invalid color")
	}
}

func Queen(color Color) Piece {
	switch color {
	case White:
		return WhiteQueen
	case Black:
		return BlackQueen
	default:
		panic("Invalid color")
	}
}

func King(color Color) Piece {
	switch color {
	case White:
		return WhiteKing
	case Black:
		return BlackKing
	default:
		panic("Invalid color")
	}
}

func Knight(color Color) Piece {
	switch color {
	case White:
		return WhiteKnight
	case Black:
		return BlackKnight
	default:
		panic("Invalid color")
	}
}

func Rook(color Color) Piece {
	switch color {
	case White:
		return WhiteRook
	case Black:
		return BlackRook
	default:
		panic("Invalid color")
	}
}

func validPiece(p Piece) bool {
	switch p {
	case NoPiece,
		BlackRook,
		BlackKnight,
		BlackBishop,
		BlackQueen,
		BlackKing,
		BlackPawn,

		WhiteRook,
		WhiteKnight,
		WhiteBishop,
		WhiteQueen,
		WhiteKing,
		WhitePawn:
		return true
	}
	return false
}
