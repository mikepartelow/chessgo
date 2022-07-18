package chessgo

import "fmt"

type Game struct {
	Board Board
	Turn  Color
}

func (g *Game) Move(move string) (Piece, error) {
	// todo: *yuck
	mv, err := parseMove(move, *g)
	if err != nil {
		return NoPiece, fmt.Errorf("error parsing move %q: %v", move, err)
	}

	g.Board.Move(mv.srcAddr, mv.dstAddr)
	g.Turn = ToggleColor(g.Turn)

	// todo: tdd
	// if mv.capture && mv.replaced == NoPiece {
	// 	return NoPiece, errors.New("Expected capture but didn't.")
	// }

	return mv.captured, nil
}
