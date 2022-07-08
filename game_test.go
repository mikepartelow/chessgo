package chessgo_test

import (
	"fmt"
	"math"
	"mp/chessgo"
	"testing"
)

// func TestGame(t *testing.T) {
// 	g := chessgo.NewGame()
// 	assertSquare(t, g.Board(), "e2", 'P')

// 	// g.Move("e4")
// 	// assertSquare(t, g.Board().Square("e4"), 'P')
// }

type StubBoard struct {
	squares []byte
}

func TestSourcesForDest(t *testing.T) {
	testCases := []struct {
		board    StubBoard
		destAddr string
		want     string
	}{
		{
			board:    StubBoard{squares: []byte("P   ")},
			destAddr: "a2",
			want:     "a1",
		},
		{
			board:    StubBoard{squares: []byte(" P  ")},
			destAddr: "b2",
			want:     "b1",
		},
	}
	for _, tC := range testCases {
		desc := fmt.Sprintf("%s from %s", tC.destAddr, tC.want)
		t.Run(desc, func(t *testing.T) {
			g := chessgo.Game{Board: &tC.board}
			got := g.SourceForDest(tC.destAddr)

			if got != tC.want {
				t.Errorf("got %q, wanted %q as source for dest %q", got, tC.want, tC.destAddr)
			}
		})
	}
}

func (b *StubBoard) InBounds(addr string) bool {
	return false
}

func (b *StubBoard) GetSquare(addr string) chessgo.Piece {
	return chessgo.Piece(b.squares[b.getIndex(addr)])
}

func (b *StubBoard) SetSquare(addr string, piece chessgo.Piece) {
}

func (b *StubBoard) Move(srcAddr, dstAddr string) chessgo.Piece {
	return chessgo.NoPiece
}

func (b *StubBoard) getIndex(addr string) uint8 {
	n := uint8(math.Sqrt(float64(len(b.squares))))

	maxFile := uint8('a') + n
	maxRank := uint8('1') + n

	file := 7 - (maxFile - addr[0])
	rank := 7 - (maxRank - uint8(addr[1]))

	return file + rank*8
}
