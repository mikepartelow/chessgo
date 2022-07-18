package chessgo

import (
	"bytes"
	"fmt"
)

type MoveInfo struct {
	piece         Piece
	dstAddr       string
	srcAddr       string
	captured      Piece
	expectCapture bool
	check         bool
}

type increments struct {
	incX, incY int8
}

func parseMove(move string, g Game) (*MoveInfo, error) {
	mi, err := parseDst(move, g)
	if err != nil {
		return nil, fmt.Errorf("could not parse dst from %q: %v", move, err)
	}

	// todo: *yuck
	src, err := findSrc(*mi, g)
	if err != nil {
		return nil, fmt.Errorf("couldn't find src for %q: %v", move, err)
	}

	mi.srcAddr = src
	return mi, nil
}

func parseDst(move string, g Game) (*MoveInfo, error) {
	mi := MoveInfo{piece: Pawn(g.Turn), captured: NoPiece}

	dstAddrBuf := bytes.Buffer{}

	if move[len(move)-1:] == "+" {
		mi.check = true
		move = move[:len(move)-1]
	}

	for i := len(move) - 1; i >= 0; i-- {
		if dstAddrBuf.Len() < 2 {
			dstAddrBuf.WriteByte(move[i])
		} else {
			switch move[i] {
			case 'x':
				mi.expectCapture = true
			case 'B':
				mi.piece = Bishop(g.Turn)
			case 'Q':
				mi.piece = Queen(g.Turn)
			case 'K':
				mi.piece = King(g.Turn)
			case 'N':
				mi.piece = Knight(g.Turn)
			}
		}
	}

	mi.dstAddr = fmt.Sprintf("%c%c", dstAddrBuf.Bytes()[1], dstAddrBuf.Bytes()[0])
	mi.captured = g.Board.GetSquare(mi.dstAddr)

	if mi.captured != NoPiece && ColorOf(mi.captured) == g.Turn {
		return nil, fmt.Errorf("attempt to capture own piece at %q with %c", mi.dstAddr, mi.piece)
	}

	return &mi, nil
}

func findSrc(mi MoveInfo, g Game) (string, error) {
	switch mi.piece {
	case WhitePawn, BlackPawn:
		return findPawnSrc(mi, g)
	case WhiteBishop, BlackBishop:
		return findDiagonalSrc(mi, g)
	case WhiteQueen, BlackQueen:
		src, err := findDiagonalSrc(mi, g)
		if err != nil {
			return findHorizontalSrc(mi, g)
		}
		return src, nil
	case WhiteKing, BlackKing:
		return findKingSrc(mi, g)
	case WhiteKnight, BlackKnight:
		return findKnightSrc(mi, g)
	default:
		return "", fmt.Errorf("unhandled Piece: %c", mi.piece)
	}
}

func isSrc(mi MoveInfo, srcAddr string, g Game) bool {
	if g.Board.InBounds(srcAddr) && g.Board.GetSquare(srcAddr) == mi.piece {
		return true
	}
	return false
}

func findPawnSrc(mi MoveInfo, g Game) (string, error) {
	switch mi.piece {
	case WhitePawn:
		if mi.expectCapture {
			for _, incs := range []increments{{-1, -1}, {1, -1}} {
				addr := AddressPlus(mi.dstAddr, incs.incX, incs.incY)
				if isSrc(mi, addr, g) {
					return addr, nil
				}
			}
			return "", fmt.Errorf("expected capture, but no capture possible to %s", mi.dstAddr)
		}

		addr := AddressPlus(mi.dstAddr, 0, -1)
		if isSrc(mi, addr, g) {
			return addr, nil
		}

		// first White pawn move can be 2 squares starting from rank 2
		addr = NewAddress(AddressFile(mi.dstAddr), '2')
		if AddressRank(mi.dstAddr) == '4' && isSrc(mi, addr, g) {
			return addr, nil
		}
	case BlackPawn:
		if mi.expectCapture {
			for _, incs := range []increments{{1, 1}, {-1, 1}} {
				addr := AddressPlus(mi.dstAddr, incs.incX, incs.incY)
				if isSrc(mi, addr, g) {
					return addr, nil
				}
			}
			return "", fmt.Errorf("expected capture, but no capture possible to %s", mi.dstAddr)
		}

		addr := AddressPlus(mi.dstAddr, 0, 1)
		if isSrc(mi, addr, g) {
			return addr, nil
		}

		// first Black pawn move can be 2 squares starting from Board.MaxRank() - 1
		addr = NewAddress(AddressFile(mi.dstAddr), byte(g.Board.MaxRank())-1)
		if AddressRank(mi.dstAddr) == byte(g.Board.MaxRank()-3) && isSrc(mi, addr, g) {
			return addr, nil
		}
	}

	return "", fmt.Errorf("couldn't find Pawn src for %q", mi.dstAddr)
}

func findDiagonalSrc(mi MoveInfo, g Game) (string, error) {
	for _, diag := range []increments{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}} {
		incX, incY := diag.incX, diag.incY
		for addr := AddressPlus(mi.dstAddr, incX, incY); g.Board.InBounds(addr); addr = AddressPlus(addr, incX, incY) {
			if isSrc(mi, addr, g) {
				return addr, nil
			}
		}
	}

	return "", fmt.Errorf("couldn't find a diagonal source for %s", mi.dstAddr)

}

func findHorizontalSrc(mi MoveInfo, g Game) (string, error) {
	for _, horiz := range []increments{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		incX, incY := horiz.incX, horiz.incY
		for addr := AddressPlus(mi.dstAddr, incX, incY); g.Board.InBounds(addr); addr = AddressPlus(addr, incX, incY) {
			if isSrc(mi, addr, g) {
				return addr, nil
			}
		}
	}

	return "", fmt.Errorf("couldn't find a horizontal source for %s", mi.dstAddr)
}

func findKingSrc(mi MoveInfo, g Game) (string, error) {
	for _, offs := range []increments{{-1, -1}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}} {
		addr := AddressPlus(mi.dstAddr, offs.incX, offs.incY)
		if g.Board.InBounds(addr) && isSrc(mi, addr, g) {
			return addr, nil
		}
	}

	return "", fmt.Errorf("couldn't find a King source for %s", mi.dstAddr)
}

func findKnightSrc(mi MoveInfo, g Game) (string, error) {
	for _, offs := range []increments{{-1, -2}, {-2, -1}, {1, -2}, {2, -1}, {1, 2}, {2, 1}, {-1, 2}, {-2, 1}} {
		addr := AddressPlus(mi.dstAddr, offs.incX, offs.incY)
		if g.Board.InBounds(addr) && isSrc(mi, addr, g) {
			return addr, nil
		}
	}

	return "", fmt.Errorf("couldn't find a Knight source for %s", mi.dstAddr)
}
