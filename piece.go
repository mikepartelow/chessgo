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
