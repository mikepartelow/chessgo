package chessgo

type Board interface {
	InBounds(addr Address) bool
	GetSquare(addr Address) Piece
	SetSquare(addr Address, piece Piece)
	Move(srcAddr, dstAddr Address) Piece
	String() string
	MaxFile() byte
	MaxRank() byte
	InCheck(Color) bool
	Mated(Color) bool
}
