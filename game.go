package chessgo

import "log"

type Game struct {
	Board Board
	Turn  Color
}

func (g *Game) Move(move string) (Piece, error) {
	mv := parseMove(move, g)

	log.Printf("mv.srcAddr: %s mv.dstAddr: %s", mv.srcAddr, mv.dstAddr)

	if mv.error != nil {
		return NoPiece, mv.error
	}

	g.Board.Move(mv.srcAddr, mv.dstAddr)
	g.Turn = ToggleColor(g.Turn)

	// todo:
	// if mv.capture && mv.replaced == NoPiece {
	// 	return NoPiece, errors.New("Expected capture but didn't.")
	// }

	return mv.replaced, nil
}
