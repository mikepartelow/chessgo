package chessgo

import (
	"bytes"
	"fmt"
	"log"
)

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

func parseMove(moveStr string, g *Game) (mv move) {
	mv.piece = Pawn(g.Turn)
	mv.replaced = NoPiece

	if err := parseDst(moveStr, &mv, g); err != nil {
		return
	}

	// log.Printf("Moving %c to %s", mv.piece, mv.dstAddr)

	switch mv.piece {
	case WhitePawn, BlackPawn:
		findPawnSrc(&mv, g)
		return
	case WhiteBishop, BlackBishop:
		findBishopSrc(&mv, g)
		return
	default:
		panic(fmt.Sprintf("Unhandled Piece: %c", mv.piece))
	}
}

func parseDst(moveStr string, mv *move, g *Game) error {
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

func isSrc(mv *move, file, rank byte, g *Game) bool {
	srcAddr := fmt.Sprintf("%c%c", file, rank)
	log.Printf("%q g.Board.InBounds(srcAddr): %v g.Board.GetSquare(srcAddr): %q mv.piece: %q", srcAddr, g.Board.InBounds(srcAddr), g.Board.GetSquare(srcAddr), mv.piece)
	if g.Board.InBounds(srcAddr) && g.Board.GetSquare(srcAddr) == mv.piece {
		mv.srcAddr = srcAddr
		return true
	}
	return false
}

func findPawnSrc(mv *move, g *Game) {
	switch mv.piece {
	case WhitePawn:
		if isSrc(mv, mv.dstFile, byte(uint8(mv.dstRank)-1), g) {
			return
		}
		homeRank := byte('2') // where the white pawn starts the game
		log.Printf("homeRank: %c, other: %c, other: %c", homeRank, byte(uint8(homeRank)+2), mv.dstFile)
		if mv.dstRank == byte(uint8(homeRank)+2) && isSrc(mv, mv.dstFile, homeRank, g) {
			return
		}
	case BlackPawn:
		if isSrc(mv, mv.dstFile, byte(uint8(mv.dstRank)+1), g) {
			return
		}
		homeRank := byte(uint8(g.Board.MaxRank()) - uint8(1)) // where the black pawn starts the game
		log.Printf("homeRank: %c, other: %c, other: %c", homeRank, byte(uint8(homeRank)-2), mv.dstFile)
		if mv.dstRank == byte(uint8(homeRank)-2) && isSrc(mv, mv.dstFile, homeRank, g) {
			return
		}
	default:
		panic("WTF")
	}
}

func findBishopSrc(mv *move, g *Game) {
	if isSrc(mv, byte(uint8(mv.dstFile)-1), byte(uint8(mv.dstRank)-1), g) {
		return
	}
	if isSrc(mv, byte(uint8(mv.dstFile)-1), byte(uint8(mv.dstRank)+1), g) {
		return
	}
	if isSrc(mv, byte(uint8(mv.dstFile)+1), byte(uint8(mv.dstRank)-1), g) {
		return
	}
	if isSrc(mv, byte(uint8(mv.dstFile)+1), byte(uint8(mv.dstRank)+1), g) {
		return
	}
}
