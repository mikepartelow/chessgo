package chessgo

type WhiteQueen struct{}

func (q WhiteQueen) Byte() byte     { return 'Q' }
func (q WhiteQueen) Color() Color   { return White }
func (q WhiteQueen) String() string { return "White Queen" }
func (q WhiteQueen) SourceForDest(dstAddr Address, board Board) Address {
	return queenSourceForDest(dstAddr, q, board)
}

type BlackQueen struct{}

func (q BlackQueen) Byte() byte     { return 'q' }
func (q BlackQueen) Color() Color   { return Black }
func (q BlackQueen) String() string { return "Black Queen" }
func (q BlackQueen) SourceForDest(dstAddr Address, board Board) Address {
	return queenSourceForDest(dstAddr, q, board)
}

func queenSourceForDest(dstAddr Address, piece Piece, board Board) Address {
	srcAddr := findDiagonalSrc(dstAddr, piece, board)
	if srcAddr == "" {
		srcAddr = findHorizontalSrc(dstAddr, piece, board)
	}

	return srcAddr
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
