package chessgo

type increments struct {
	incX, incY int8
}

func findKnightSrc(dstAddr string, piece Piece, board Board) string {
	for _, offs := range []increments{{-1, -2}, {-2, -1}, {1, -2}, {2, -1}, {1, 2}, {2, 1}, {-1, 2}, {-2, 1}} {
		addr := addressPlus(dstAddr, offs.incX, offs.incY)
		if board.InBounds(addr) && board.GetSquare(addr) == piece {
			return addr
		}
	}

	return ""
}

func findDiagonalSrc(dstAddr string, piece Piece, board Board) string {
	return findSrcOnLines(dstAddr, piece, board, []increments{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}})
}

func findHorizontalSrc(dstAddr string, piece Piece, board Board) string {
	return findSrcOnLines(dstAddr, piece, board, []increments{{-1, 0}, {1, 0}, {0, -1}, {0, 1}})
}

func findSrcOnLines(dstAddr string, piece Piece, board Board, incs []increments) string {
	for _, incs := range incs {
		if srcAddr := findSrcOnLine(dstAddr, piece, board, incs.incX, incs.incY); srcAddr != "" {
			return srcAddr
		}
	}

	return ""
}

func findSrcOnLine(dstAddr string, piece Piece, board Board, incX, incY int8) string {
	for addr := addressPlus(dstAddr, incX, incY); board.InBounds(addr); addr = addressPlus(addr, incX, incY) {
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

func findKingSrc(dstAddr string, piece Piece, board Board) string {
	for _, offs := range []increments{{-1, -1}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}} {
		addr := addressPlus(dstAddr, offs.incX, offs.incY)
		if board.InBounds(addr) && board.GetSquare(addr) == piece {
			return addr
		}
	}

	return ""
}
