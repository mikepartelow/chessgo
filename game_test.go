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
		board        StubBoard
		turn         chessgo.Color
		move         string
		wantBoard    StubBoard
		wantCaptured chessgo.Piece
		wantErr      chessgo.Error
	}{
		{
			board:        StubBoard{squares: []byte("P   ")},
			move:         "a2",
			wantBoard:    StubBoard{squares: []byte("  P ")},
			wantCaptured: chessgo.NoPiece,
		},
		{
			board:        StubBoard{squares: []byte(" P  ")},
			move:         "b2",
			wantBoard:    StubBoard{squares: []byte("   P")},
			wantCaptured: chessgo.NoPiece,
		},
		{
			board:        StubBoard{squares: []byte("B   ")},
			move:         "Bb2",
			wantBoard:    StubBoard{squares: []byte("   B")},
			wantCaptured: chessgo.NoPiece,
		},
		{
			board:        StubBoard{squares: []byte("b   ")},
			turn:         chessgo.Black,
			move:         "Bb2",
			wantBoard:    StubBoard{squares: []byte("   b")},
			wantCaptured: chessgo.NoPiece,
		},
		{
			board:        StubBoard{squares: []byte("  p ")},
			turn:         chessgo.Black,
			move:         "a1",
			wantBoard:    StubBoard{squares: []byte("p   ")},
			wantCaptured: chessgo.NoPiece,
		},
		{
			board:        StubBoard{squares: []byte("b  N")},
			turn:         chessgo.Black,
			move:         "Bxb2",
			wantBoard:    StubBoard{squares: []byte("   b")},
			wantCaptured: chessgo.WhiteKnight,
		},
		{
			board:        StubBoard{squares: []byte("b  n")},
			turn:         chessgo.Black,
			move:         "Bxb2",
			wantBoard:    StubBoard{squares: []byte("b  n")},
			wantCaptured: chessgo.NoPiece,
			wantErr:      chessgo.ErrorFriendlyFire{},
		},
	}
	for _, tC := range testCases {
		desc := fmt.Sprintf("move %s", tC.move)
		t.Run(desc, func(t *testing.T) {
			g := chessgo.Game{Board: &tC.board, Turn: tC.turn}
			captured, err := g.Move(tC.move)

			if err != tC.wantErr {
				t.Errorf("got err %v wanted %v", err, tC.wantErr)
			}

			if !reflect.DeepEqual(g.Board, &tC.wantBoard) {
				t.Errorf("got %+v, wanted %+v after move %q", g.Board, &tC.wantBoard, tC.move)
			}

			if captured != tC.wantCaptured {
				t.Errorf("got %q captured, wanted %q", captured, tC.wantCaptured)
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
	captured := b.GetSquare(dstAddr)
	log.Printf("Moving %q to %q", srcAddr, dstAddr)
	b.SetSquare(dstAddr, b.GetSquare(srcAddr))
	b.SetSquare(srcAddr, chessgo.NoPiece)
	return captured
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
