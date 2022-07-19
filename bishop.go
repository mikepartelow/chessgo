package chessgo

type WhiteBishop struct{}

func (b WhiteBishop) Byte() byte     { return 'B' }
func (b WhiteBishop) Color() Color   { return White }
func (b WhiteBishop) String() string { return "White Bishop" }
func (b WhiteBishop) SourceForDest(dstAddr Address, board Board) Address {
	return bishopSourceForDest(dstAddr, b, board)
}

type BlackBishop struct{}

func (b BlackBishop) Byte() byte     { return 'b' }
func (b BlackBishop) Color() Color   { return Black }
func (b BlackBishop) String() string { return "Black Bishop" }
func (b BlackBishop) SourceForDest(dstAddr Address, board Board) Address {
	return bishopSourceForDest(dstAddr, b, board)
}

func bishopSourceForDest(dstAddr Address, piece Piece, board Board) Address {
	return findDiagonalSrc(dstAddr, piece, board)
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
