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
	if g.Board.InBounds(srcAddr) && g.Board.GetSquare(srcAddr) == mv.piece {
		mv.srcAddr = srcAddr
		return true
	}
	return false
}

func findPawnSrc(mv *move, g *Game) {
	switch mv.piece {
	case WhitePawn:
		if mv.capture {
			if isSrc(mv, bytePlus(mv.dstFile, -1), bytePlus(mv.dstRank, -1), g) {
				return
			}
			if isSrc(mv, bytePlus(mv.dstFile, 1), bytePlus(mv.dstRank, -1), g) {
				return
			}
			panic("Illegal move.")
		}

		if isSrc(mv, mv.dstFile, bytePlus(mv.dstRank, -1), g) {
			return
		}
		homeRank := byte('2') // where the white pawn starts the game
		log.Printf("homeRank: %c, other: %c, other: %c", homeRank, bytePlus(homeRank, 2), mv.dstFile)
		if mv.dstRank == bytePlus(homeRank, 2) && isSrc(mv, mv.dstFile, homeRank, g) {
			return
		}
	case BlackPawn:
		if mv.capture {
			if isSrc(mv, bytePlus(mv.dstFile, 1), bytePlus(mv.dstRank, 1), g) {
				return
			}
			if isSrc(mv, bytePlus(mv.dstFile, -1), bytePlus(mv.dstRank, 1), g) {
				return
			}
			panic("Illegal move.")
		}

		if isSrc(mv, mv.dstFile, bytePlus(mv.dstRank, 1), g) {
			return
		}
		homeRank := bytePlus(byte(g.Board.MaxRank()), -1) // where the black pawn starts the game
		log.Printf("homeRank: %c, other: %c, other: %c", homeRank, bytePlus(homeRank, -2), mv.dstFile)
		if mv.dstRank == bytePlus(homeRank, -2) && isSrc(mv, mv.dstFile, homeRank, g) {
			return
		}

	default:
		panic("WTF")
	}
}

func bytePlus(b byte, n int8) byte {
	return byte(int8(b) + n)
}

func addrPlus(addr string, incX, incY int8) string {
	return fmt.Sprintf("%c%c", bytePlus(addr[0], incX), bytePlus(addr[1], incY))
}

func findBishopSrc(mv *move, g *Game) {
	log.Printf("findBishopSrc(%s)", mv.dstAddr)
	diagonals := []struct {
		incX int8
		incY int8
	}{
		{
			-1, -1,
		},
		{
			1, 1,
		},
		{
			1, -1,
		},
		{
			-1, 1,
		},
	}

	for _, diag := range diagonals {
		incX, incY := diag.incX, diag.incY
		log.Printf(" incX/incY=%d/%d", incX, incY)
		for addr := addrPlus(mv.dstAddr, incX, incY); g.Board.InBounds(addr); addr = addrPlus(addr, incX, incY) {
			if isSrc(mv, addr[0], addr[1], g) {
				log.Printf("Found at %q", addr)
				return
			}
		}
	}
	// other cases not tested
}
