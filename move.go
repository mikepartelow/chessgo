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
			case 'R':
				mi.piece = Rook(g.Turn)
			}
		}
	}

	mi.dstAddr = NewAddress(dstAddrBuf.Bytes()[1], dstAddrBuf.Bytes()[0])
	mi.captured = g.Board.GetSquare(mi.dstAddr)

	if mi.captured != NoPiece && ColorOf(mi.captured) == g.Turn {
		return nil, fmt.Errorf("attempt to capture own piece at %q with %c", mi.dstAddr, mi.piece)
	}

	return &mi, nil
}

func findSrc(mi MoveInfo, g Game) (string, error) {
	var srcAddr string

	switch mi.piece {
	case WhitePawn, BlackPawn:
		srcAddr, _ = findPawnSrc(mi, g)
	case WhiteBishop, BlackBishop:
		srcAddr = findDiagonalSrc(mi.dstAddr, mi.piece, g.Board)
	case WhiteQueen, BlackQueen:
		srcAddr = findDiagonalSrc(mi.dstAddr, mi.piece, g.Board)
		if srcAddr == "" {
			srcAddr = findHorizontalSrc(mi.dstAddr, mi.piece, g.Board)
		}
	case WhiteKing, BlackKing:
		srcAddr = findKingSrc(mi.dstAddr, mi.piece, g.Board)
	case WhiteKnight, BlackKnight:
		srcAddr = findKnightSrc(mi.dstAddr, mi.piece, g.Board)
	case WhiteRook, BlackRook:
		srcAddr = findHorizontalSrc(mi.dstAddr, mi.piece, g.Board)
	default:
		return "", fmt.Errorf("unhandled Piece: %c", mi.piece)
	}

	if srcAddr == "" {
		return "", fmt.Errorf("could not find source %v for dest %s", mi.piece, mi.dstAddr)
	}

	return srcAddr, nil
}

func findPawnSrc(mi MoveInfo, g Game) (string, error) {
	isSrc := func(mi MoveInfo, addr string, g Game) bool {
		return g.Board.InBounds(addr) && g.Board.GetSquare(addr) == mi.piece
	}

	switch mi.piece {
	case WhitePawn:
		if mi.expectCapture {
			for _, incs := range []increments{{-1, -1}, {1, -1}} {
				addr := addressPlus(mi.dstAddr, incs.incX, incs.incY)
				if isSrc(mi, addr, g) {
					return addr, nil
				}
			}
			return "", fmt.Errorf("expected capture, but no capture possible to %s", mi.dstAddr)
		}

		addr := addressPlus(mi.dstAddr, 0, -1)
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
				addr := addressPlus(mi.dstAddr, incs.incX, incs.incY)
				if isSrc(mi, addr, g) {
					return addr, nil
				}
			}
			return "", fmt.Errorf("expected capture, but no capture possible to %s", mi.dstAddr)
		}

		addr := addressPlus(mi.dstAddr, 0, 1)
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
