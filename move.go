package chessgo

import (
	"bytes"
	"fmt"
)

type move struct {
	srcAddr  string
	dstFile  byte
	dstRank  byte
	dstAddr  string
	piece    Piece
	replaced Piece
	capture  bool
	check    bool
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
	case WhiteBishop, BlackBishop, WhiteQueen, BlackQueen:
		findDiagonalSrc(&mv, g)
		return
	default:
		panic(fmt.Sprintf("Unhandled Piece: %c", mv.piece))
	}
}

func parseDst(moveStr string, mv *move, g *Game) error {
	dstAddrBuf := bytes.Buffer{}

	if moveStr[len(moveStr)-1:] == "+" {
		mv.check = true
		moveStr = moveStr[:len(moveStr)-1]
	}

	for i := len(moveStr) - 1; i >= 0; i-- {
		if dstAddrBuf.Len() < 2 {
			dstAddrBuf.WriteByte(moveStr[i])
		} else {
			switch moveStr[i] {
			case 'x':
				mv.capture = true
			case 'B':
				mv.piece = Bishop(g.Turn)
			case 'Q':
				mv.piece = Queen(g.Turn)
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
		// log.Printf("homeRank: %c, other: %c, other: %c", homeRank, bytePlus(homeRank, 2), mv.dstFile)
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
		// log.Printf("homeRank: %c, other: %c, other: %c", homeRank, bytePlus(homeRank, -2), mv.dstFile)
		if mv.dstRank == bytePlus(homeRank, -2) && isSrc(mv, mv.dstFile, homeRank, g) {
			return
		}

	default:
		panic("WTF")
	}
}

func findDiagonalSrc(mv *move, g *Game) {
	// log.Printf("findDiagonalSrc(%s)", mv.dstAddr)
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

	// log.Printf("findDiagonalSrc(%c)", mv.piece)

	for _, diag := range diagonals {
		incX, incY := diag.incX, diag.incY
		// log.Printf(" incX/incY=%d/%d", incX, incY)
		for addr := AddressPlus(mv.dstAddr, incX, incY); g.Board.InBounds(addr); addr = AddressPlus(addr, incX, incY) {
			// log.Printf("  Checking at %q: %c", addr, g.Board.GetSquare(addr))
			if isSrc(mv, addr[0], addr[1], g) {
				// log.Printf("Found at %q", addr)
				return
			}
		}
	}
	// other cases not tested
}