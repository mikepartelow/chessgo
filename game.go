package chessgo

import (
	"bytes"
	"fmt"
	"log"
)

type Game struct {
	Board Board
	Turn  Color
}

type move struct {
	srcAddr  string
	dstAddr  string
	piece    Piece
	replaced Piece
	capture  bool
	error    error
}

func (g *Game) Move(move string) (Piece, error) {
	mv := g.parseMove(move)

	if mv.error != nil {
		return NoPiece, mv.error
	}

	log.Printf("%+v\n", mv)
	g.Board.Move(mv.srcAddr, mv.dstAddr)
	g.Turn = ToggleColor(g.Turn)

	// todo:
	// if mv.capture && mv.replaced == NoPiece {
	// 	return NoPiece, errors.New("Expected capture but didn't.")
	// }

	return mv.replaced, nil
}

func (g *Game) parseMove(moveStr string) (mv move) {
	// todo: refactor this even more
	mv.piece = Pawn(g.Turn)
	mv.replaced = NoPiece

	dstAddrBuf := bytes.Buffer{}

	for i := len(moveStr) - 1; i >= 0; i-- {
		if dstAddrBuf.Len() < 2 {
			dstAddrBuf.WriteByte(moveStr[i])
		} else {
			switch moveStr[i] {
			case 'x':
				mv.capture = true
			case 'B':
				mv.piece = Bishop(g.Turn)
			}
		}
	}

	dstFile := dstAddrBuf.Bytes()[1]
	dstRank := dstAddrBuf.Bytes()[0]
	mv.dstAddr = fmt.Sprintf("%c%c", dstFile, dstRank)
	mv.replaced = g.Board.GetSquare(mv.dstAddr)

	if mv.replaced != NoPiece && ColorOf(mv.replaced) == g.Turn {
		mv.error = ErrorFriendlyFire{}
	}

	log.Printf("Moving %c to %s", mv.piece, mv.dstAddr)

	switch mv.piece {
	case WhitePawn:
		if g.isSrc(&mv, dstFile, byte(uint8(dstRank)-1)) {
			return
		}
	case BlackPawn:
		if g.isSrc(&mv, dstFile, byte(uint8(dstRank)+1)) {
			return
		}
	case WhiteBishop, BlackBishop:
		if g.isSrc(&mv, byte(uint8(dstFile)-1), byte(uint8(dstRank)-1)) {
			return
		}
		if g.isSrc(&mv, byte(uint8(dstFile)-1), byte(uint8(dstRank)+1)) {
			return
		}
		if g.isSrc(&mv, byte(uint8(dstFile)+1), byte(uint8(dstRank)-1)) {
			return
		}
		if g.isSrc(&mv, byte(uint8(dstFile)+1), byte(uint8(dstRank)+1)) {
			return
		}
	default:
		panic(fmt.Sprintf("Unhandled Piece: %c", mv.piece))
	}

	panic(fmt.Sprintf("Failed to parse move string: %s", moveStr))
}

func (g *Game) isSrc(mv *move, file, rank byte) bool {
	mv.srcAddr = fmt.Sprintf("%c%c", file, rank)
	return g.Board.InBounds(mv.srcAddr) && g.Board.GetSquare(mv.srcAddr) == mv.piece
}
