package chessgo_test

import (
	"fmt"
	"mp/chessgo"
	"reflect"
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

func TestGameMove(t *testing.T) {
	testCases := []struct {
		board StubBoard
		move  string
		want  StubBoard
	}{
		{
			board: StubBoard{squares: []byte("P   ")},
			move:  "a2",
			want:  StubBoard{squares: []byte("  P ")},
		},
		{
			board: StubBoard{squares: []byte(" P   ")},
			move:  "a2",
			want:  StubBoard{squares: []byte("   P")},
		},
	}
	for _, tC := range testCases {
		desc := fmt.Sprintf("move %s ", tC.move)
		t.Run(desc, func(t *testing.T) {
			g := chessgo.Game{Board: &tC.board}
			g.Move(tC.move)
			got := g.Board

			if !reflect.DeepEqual(got, &tC.want) {
				t.Errorf("got %+v, wanted %+v after move %q", got, &tC.want, tC.move)
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
	b.squares[b.getIndex(addr)] = byte(piece)
}

func (b *StubBoard) Move(srcAddr, dstAddr string) chessgo.Piece {
	return chessgo.NoPiece
}

func (b *StubBoard) getIndex(addr string) uint8 {
	file := addr[0]
	rank := addr[1]

	var x, y uint8

	switch file {
	case 'a':
		x = 0
	case 'b':
		x = 1
	}

	switch rank {
	case '1':
		y = 0
	case '2':
		y = 1
	}

	return y*2 + x
}
