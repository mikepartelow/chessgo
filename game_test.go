package chessgo_test

import (
	"fmt"
	"log"
	"mp/chessgo"
	"reflect"
	"testing"
)

// todo: this is an integration test!

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
		turn  chessgo.Color
		move  string
		want  StubBoard
	}{
		{
			board: StubBoard{squares: []byte("P   ")},
			move:  "a2",
			want:  StubBoard{squares: []byte("  P ")},
		},
		{
			board: StubBoard{squares: []byte(" P  ")},
			move:  "b2",
			want:  StubBoard{squares: []byte("   P")},
		},
		{
			board: StubBoard{squares: []byte("B   ")},
			move:  "Bb2",
			want:  StubBoard{squares: []byte("   B")},
		},
		{
			board: StubBoard{squares: []byte("b   ")},
			turn:  chessgo.Black,
			move:  "Bb2",
			want:  StubBoard{squares: []byte("   b")},
		},
		{
			board: StubBoard{squares: []byte("  p ")},
			turn:  chessgo.Black,
			move:  "a1",
			want:  StubBoard{squares: []byte("p   ")},
		},
	}
	for _, tC := range testCases {
		desc := fmt.Sprintf("move %s", tC.move)
		t.Run(desc, func(t *testing.T) {
			g := chessgo.Game{Board: &tC.board, Turn: tC.turn}
			g.Move(tC.move)
			got := g.Board

			if !reflect.DeepEqual(got, &tC.want) {
				t.Errorf("got %+v, wanted %+v after move %q", got, &tC.want, tC.move)
			}
		})
	}

	t.Run("turn toggles on valid move", func(t *testing.T) {
		g := chessgo.Game{Board: &StubBoard{squares: []byte("P   ")}}
		if g.Turn != chessgo.White {
			t.Errorf("expected Turn = White")
		}
		g.Move("a2")
		if g.Turn != chessgo.Black {
			t.Errorf("expected Turn = Black")
		}
	})

	// t.Run("turn does not toggle on invalid move", func(t *testing.T) {

	// })

}

func (b *StubBoard) InBounds(addr string) bool {
	return b.getIndex(addr) < uint8(len(b.squares))
}

func (b *StubBoard) GetSquare(addr string) chessgo.Piece {
	return chessgo.Piece(b.squares[b.getIndex(addr)])
}

func (b *StubBoard) SetSquare(addr string, piece chessgo.Piece) {
	b.squares[b.getIndex(addr)] = byte(piece)
}

func (b *StubBoard) Move(srcAddr, dstAddr string) chessgo.Piece {
	replaced := b.GetSquare(dstAddr)
	log.Printf("Moving %q to %q", srcAddr, dstAddr)
	b.SetSquare(dstAddr, b.GetSquare(srcAddr))
	b.SetSquare(srcAddr, chessgo.NoPiece)
	return replaced
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
