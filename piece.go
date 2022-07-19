package chessgo

import "fmt"

type Piece interface {
	Byte() byte
	Color() Color
	String() string
	SourceForDest(dstAddr Address, board Board) Address
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
