package chessgo

type WhiteKing struct{}

func (k WhiteKing) Byte() byte     { return 'K' }
func (k WhiteKing) Color() Color   { return White }
func (p WhiteKing) String() string { return "White King" }
func (k WhiteKing) SourceForDest(dstAddr Address, board Board) Address {
	return kingSourceForDest(dstAddr, k, board)
}

type BlackKing struct{}

func (k BlackKing) Byte() byte     { return 'k' }
func (k BlackKing) Color() Color   { return Black }
func (p BlackKing) String() string { return "Black King" }
func (k BlackKing) SourceForDest(dstAddr Address, board Board) Address {
	return kingSourceForDest(dstAddr, k, board)
}

func kingSourceForDest(dstAddr Address, piece Piece, board Board) Address {
	for _, offs := range []increments{{-1, -1}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}} {
		addr := dstAddr.Plus(offs.incX, offs.incY)
		if board.InBounds(addr) && board.GetSquare(addr) == piece {
			return addr
		}
	}

	return ""
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
