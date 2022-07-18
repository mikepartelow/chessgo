package chessgo

type increments struct {
	incX, incY int8
}

func findKnightSrc(dstAddr Address, piece Piece, board Board) Address {
	for _, offs := range []increments{{-1, -2}, {-2, -1}, {1, -2}, {2, -1}, {1, 2}, {2, 1}, {-1, 2}, {-2, 1}} {
		addr := dstAddr.Plus(offs.incX, offs.incY)
		if board.InBounds(addr) && board.GetSquare(addr) == piece {
			return addr
		}
	}

	return ""
}

func findDiagonalSrc(dstAddr Address, piece Piece, board Board) Address {
	return findSrcOnLines(dstAddr, piece, board, []increments{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}})
}

func findHorizontalSrc(dstAddr Address, piece Piece, board Board) Address {
	return findSrcOnLines(dstAddr, piece, board, []increments{{-1, 0}, {1, 0}, {0, -1}, {0, 1}})
}

func findSrcOnLines(dstAddr Address, piece Piece, board Board, incs []increments) Address {
	for _, incs := range incs {
		if srcAddr := findSrcOnLine(dstAddr, piece, board, incs.incX, incs.incY); srcAddr != "" {
			return srcAddr
		}
	}

	return ""
}

func findSrcOnLine(dstAddr Address, piece Piece, board Board, incX, incY int8) Address {
	for addr := dstAddr.Plus(incX, incY); board.InBounds(addr); addr = addr.Plus(incX, incY) {
		switch board.GetSquare(addr) {
		case piece:
			return addr
		case NoPiece:
			continue
		default:
			// move is blocked by some piece
			return ""
		}
	}
	return ""
}

func findKingSrc(dstAddr Address, piece Piece, board Board) Address {
	for _, offs := range []increments{{-1, -1}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}} {
		addr := dstAddr.Plus(offs.incX, offs.incY)
		if board.InBounds(addr) && board.GetSquare(addr) == piece {
			return addr
		}
	}

	return ""
}
