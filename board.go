package chessgo

type Board interface {
	InBounds(addr string) bool
	GetSquare(addr string) Piece
	SetSquare(addr string, piece Piece)
	Move(srcAddr, dstAddr string) Piece
	String() string
	MaxFile() rune
	MaxRank() rune
	InCheck(Color) bool
	Mated(Color) bool
}
