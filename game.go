package chessgo

import (
	"fmt"
	"log"
)

type Game struct {
	Board Board
	Turn  Color
}

func (g *Game) Move(move string) (Piece, error) {
	// todo: *yuck
	mi, err := parseMove(move, *g)
	if err != nil {
		return nil, fmt.Errorf("error parsing move %q: %v", move, err)
	}

	if mi.kingSideCastle {
		var kingSrc, rookSrc Address
		if g.Turn == White {
			kingSrc, rookSrc = Address("e1"), Address("h1")
		} else {
			kingSrc, rookSrc = Address("e8"), Address("h8")
		}

		kingDst := kingSrc.Plus(2, 0)
		rookDst := rookSrc.Plus(-2, 0)

		log.Printf(" king side castle for %s -> %s, %s -> %s", kingSrc, kingDst, rookSrc, rookDst)

		g.Board.Move(kingSrc, kingDst)
		g.Board.Move(rookSrc, rookDst)
	} else {
		g.Board.Move(mi.srcAddr, mi.dstAddr)
	}

	g.Turn = g.Turn.Opponent()

	// todo: tdd
	// if mi.capture && mi.replaced == NoPiece {
	// 	return NoPiece, errors.New("Expected capture but didn't.")
	// }

	return mi.captured, nil
}
