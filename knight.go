package chessgo

type WhiteKnight struct{}

func (n WhiteKnight) Byte() byte     { return 'N' }
func (n WhiteKnight) Color() Color   { return White }
func (n WhiteKnight) String() string { return "White Knight" }
func (n WhiteKnight) SourceForDest(dstAddr Address, board Board) Address {
	return knightSourceForDest(dstAddr, n, board)
}

type BlackKnight struct{}

func (n BlackKnight) Byte() byte     { return 'n' }
func (n BlackKnight) Color() Color   { return Black }
func (n BlackKnight) String() string { return "Black Knight" }
func (n BlackKnight) SourceForDest(dstAddr Address, board Board) Address {
	return knightSourceForDest(dstAddr, n, board)
}

func knightSourceForDest(dstAddr Address, piece Piece, board Board) Address {
	for _, offs := range []increments{{-1, -2}, {-2, -1}, {1, -2}, {2, -1}, {1, 2}, {2, 1}, {-1, 2}, {-2, 1}} {
		addr := dstAddr.Plus(offs.incX, offs.incY)
		if board.InBounds(addr) && board.GetSquare(addr) == piece {
			return addr
		}
	}

	return ""
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
