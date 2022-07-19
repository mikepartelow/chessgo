package chessgo

type WhiteRook struct{}

func (r WhiteRook) Byte() byte     { return 'R' }
func (r WhiteRook) Color() Color   { return White }
func (r WhiteRook) String() string { return "White Rook" }
func (r WhiteRook) SourceForDest(dstAddr Address, board Board) Address {
	return rookSourceForDest(dstAddr, r, board)
}

type BlackRook struct{}

func (r BlackRook) Byte() byte     { return 'r' }
func (r BlackRook) Color() Color   { return Black }
func (r BlackRook) String() string { return "Black Rook" }
func (r BlackRook) SourceForDest(dstAddr Address, board Board) Address {
	return rookSourceForDest(dstAddr, r, board)
}

func rookSourceForDest(dstAddr Address, piece Piece, board Board) Address {
	return findHorizontalSrc(dstAddr, piece, board)
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
