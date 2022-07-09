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
	dstFile  byte
	dstRank  byte
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

func (g *Game) parseMove(moveStr string) (mv move) { // todo: refactor this even more
	mv.piece = Pawn(g.Turn)
	mv.replaced = NoPiece

	if err := g.parseDst(moveStr, &mv); err != nil {
		return
	}

	log.Printf("Moving %c to %s", mv.piece, mv.dstAddr)

	switch mv.piece {
	case WhitePawn, BlackPawn:
		g.findPawnSrc(&mv)
		return
	case WhiteBishop, BlackBishop:
		g.findBishopSrc(&mv)
		return
	default:
		panic(fmt.Sprintf("Unhandled Piece: %c", mv.piece))
	}
}

func (g *Game) parseDst(moveStr string, mv *move) error {
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

	mv.dstFile, mv.dstRank = dstAddrBuf.Bytes()[1], dstAddrBuf.Bytes()[0]

	mv.dstAddr = fmt.Sprintf("%c%c", mv.dstFile, mv.dstRank)
	mv.replaced = g.Board.GetSquare(mv.dstAddr)

	if mv.replaced != NoPiece && ColorOf(mv.replaced) == g.Turn {
		mv.error = ErrorFriendlyFire{}
	}

	return mv.error
}

func (g *Game) isSrc(mv *move, file, rank byte) bool {
	srcAddr := fmt.Sprintf("%c%c", file, rank)
	if g.Board.InBounds(srcAddr) && g.Board.GetSquare(srcAddr) == mv.piece {
		mv.srcAddr = srcAddr
		return true
	}
	return false
}

func (g *Game) findPawnSrc(mv *move) {
	switch mv.piece {
	case WhitePawn:
		if g.isSrc(mv, mv.dstFile, byte(uint8(mv.dstRank)-1)) {
			return
		}
	case BlackPawn:
		if g.isSrc(mv, mv.dstFile, byte(uint8(mv.dstRank)+1)) {
			return
		}
	default:
		panic("WTF")
	}
}

func (g *Game) findBishopSrc(mv *move) {
	if g.isSrc(mv, byte(uint8(mv.dstFile)-1), byte(uint8(mv.dstRank)-1)) {
		return
	}
	if g.isSrc(mv, byte(uint8(mv.dstFile)-1), byte(uint8(mv.dstRank)+1)) {
		return
	}
	if g.isSrc(mv, byte(uint8(mv.dstFile)+1), byte(uint8(mv.dstRank)-1)) {
		return
	}
	if g.isSrc(mv, byte(uint8(mv.dstFile)+1), byte(uint8(mv.dstRank)+1)) {
		return
	}
}
