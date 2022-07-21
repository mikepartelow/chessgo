package chessgo

import (
	"bytes"
	"fmt"
)

type MoveInfo struct {
	piece          Piece
	dstAddr        Address
	srcAddr        Address
	captured       Piece
	expectCapture  bool
	check          bool
	mate           bool
	kingSideCastle bool
}

func parseMove(move string, g Game) (*MoveInfo, error) {
	if move == "O-O" { // todo: 0-0 (zero instead of letter O) may be acceptable. confirm and test.
		return &MoveInfo{kingSideCastle: true}, nil
	}

	mi, err := parseDst(move, g)
	if err != nil {
		return nil, fmt.Errorf("could not parse dst from %q: %v", move, err)
	}

	srcAddr := mi.piece.SourceForDest(mi.dstAddr, g.Board)
	if srcAddr == "" {
		return nil, fmt.Errorf("couldn't find src for %q: %v", move, err)
	}

	mi.srcAddr = srcAddr

	return mi, nil
}

func parseDst(move string, g Game) (*MoveInfo, error) {
	mi := MoveInfo{piece: Pawn(g.Turn), captured: nil}

	dstAddrBuf := bytes.Buffer{}

	if move[len(move)-1:] == "+" {
		mi.check = true
		move = move[:len(move)-1]
	}

	if move[len(move)-1:] == "#" {
		mi.check = true
		mi.mate = true
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

	if mi.captured != nil && mi.captured.Color() == g.Turn {
		return nil, fmt.Errorf("attempt to capture own piece at %q with %c", mi.dstAddr, mi.piece)
	}

	return &mi, nil
}

type increments struct {
	incX, incY int8
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
		case nil:
			continue
		default:
			// move is blocked by some piece
			return ""
		}
	}
	return ""
}
