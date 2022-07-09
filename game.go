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
	srcAddr string
	dstAddr string
	piece   Piece
}

func (g *Game) Move(move string) {
	mv := g.parseMove(move)
	log.Printf("%+v\n", mv)
	g.Board.Move(mv.srcAddr, mv.dstAddr)
	g.Turn = ToggleColor(g.Turn)
}

func (g *Game) parseMove(moveStr string) (mv move) {
	// todo: refactor this even more
	mv.piece = Pawn(g.Turn)

	dstAddrBuf := bytes.Buffer{}

	for i := len(moveStr) - 1; i >= 0; i-- {
		if dstAddrBuf.Len() < 2 {
			dstAddrBuf.WriteByte(moveStr[i])
		} else {
			switch moveStr[i] {
			case 'B':
				mv.piece = Bishop(g.Turn)
			}
		}
	}

	dstFile := dstAddrBuf.Bytes()[1]
	dstRank := dstAddrBuf.Bytes()[0]
	mv.dstAddr = fmt.Sprintf("%c%c", dstFile, dstRank)

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
