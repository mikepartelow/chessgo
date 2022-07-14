package chessgo_test

import (
	"fmt"
	"log"
	"math"
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
			board:     StubBoard{squares: []byte("P   ")},
			move:      "a2",
			wantBoard: StubBoard{squares: []byte("  P ")},
		},
		{
			board:     StubBoard{squares: []byte(" P  ")},
			move:      "b2",
			wantBoard: StubBoard{squares: []byte("   P")},
		},
		{
			board:     StubBoard{squares: []byte("B   ")},
			move:      "Bb2",
			wantBoard: StubBoard{squares: []byte("   B")},
		},
		{
			board:     StubBoard{squares: []byte("b   ")},
			turn:      chessgo.Black,
			move:      "Bb2",
			wantBoard: StubBoard{squares: []byte("   b")},
		},
		{
			board:     StubBoard{squares: []byte("  p ")},
			turn:      chessgo.Black,
			move:      "a1",
			wantBoard: StubBoard{squares: []byte("p   ")},
		},
		{
			board:        StubBoard{squares: []byte("b  N")},
			turn:         chessgo.Black,
			move:         "Bxb2",
			wantBoard:    StubBoard{squares: []byte("   b")},
			wantCaptured: chessgo.WhiteKnight,
		},
		{
			board:     StubBoard{squares: []byte("b  n")},
			turn:      chessgo.Black,
			move:      "Bxb2",
			wantBoard: StubBoard{squares: []byte("b  n")},
			wantErr:   chessgo.ErrorFriendlyFire{},
		},
		{
			board:     StubBoard{squares: []byte("    P           ")},
			move:      "a4",
			wantBoard: StubBoard{squares: []byte("            P   ")},
		},
		{
			board:     StubBoard{squares: []byte("        p       ")},
			turn:      chessgo.Black,
			move:      "a1",
			wantBoard: StubBoard{squares: []byte("p               ")},
		},
		{
			board:        StubBoard{squares: []byte("P  q")},
			move:         "axb2",
			wantBoard:    StubBoard{squares: []byte("   P")},
			wantCaptured: chessgo.BlackQueen,
		},
		{
			board:        StubBoard{squares: []byte(" Pq ")},
			move:         "bxa2",
			wantBoard:    StubBoard{squares: []byte("  P ")},
			wantCaptured: chessgo.BlackQueen,
		},
		// todo: disambiguate the move (left diagonal capture) between two possible pawns
		// todo: disambiguate the move (right diagonal capture) between two possible pawns
		{
			board:        StubBoard{squares: []byte("Q  p")},
			turn:         chessgo.Black,
			move:         "bxa1",
			wantBoard:    StubBoard{squares: []byte("p   ")},
			wantCaptured: chessgo.WhiteQueen,
		},
		{
			board:        StubBoard{squares: []byte(" Qp ")},
			turn:         chessgo.Black,
			move:         "axb1",
			wantBoard:    StubBoard{squares: []byte(" p  ")},
			wantCaptured: chessgo.WhiteQueen,
		},
		// todo: disambiguate the move (left diagonal capture) between two possible pawns
		// todo: disambiguate the move (right diagonal capture) between two possible pawns

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
				t.Errorf("got %q, wanted %q after move %q", g.Board, &tC.wantBoard, tC.move)
			}

			if tC.wantCaptured != chessgo.Piece(0) && captured != tC.wantCaptured {
				t.Errorf("got %q captured, wanted %q", captured, tC.wantCaptured)
			}
		})
	}

	t.Run("turn toggles on valid move", func(t *testing.T) {
		g := chessgo.Game{Board: &StubBoard{squares: []byte("P   ")}}
		if g.Turn != chessgo.White {
			t.Errorf("expected Turn = White")
		}
		_, err := g.Move("a2")

		if err != nil {
			t.Errorf("expected nil err")
		}

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
	log.Printf("Moving %q (%d) to %q (%d)", srcAddr, b.getIndex(srcAddr), dstAddr, b.getIndex(dstAddr))
	b.SetSquare(dstAddr, b.GetSquare(srcAddr))
	b.SetSquare(srcAddr, chessgo.NoPiece)
	return captured
}

func (b *StubBoard) String() string {
	return string(b.squares)
}

func (b *StubBoard) MaxFile() rune {
	m := rune(uint8('a') + uint8(math.Sqrt(float64(len(b.squares)))))
	log.Printf("MaxFile: %c", m)
	return m
}

func (b *StubBoard) MaxRank() rune {
	log.Printf("MaxRank: len(squares): %d", len(b.squares))
	m := rune(uint8('0') + uint8(math.Sqrt(float64(len(b.squares)))))
	log.Printf("MaxRank: %c", m)
	return m
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
	case 'c':
		x = 2
	case 'd':
		x = 3
	}

	switch rank {
	case '1':
		y = 0
	case '2':
		y = 1
	case '3':
		y = 2
	case '4':
		y = 3
	}

	idx := y*uint8(math.Sqrt(float64(len(b.squares)))) + x
	log.Printf("%q => x=%d, y=%d, idx=%d, len(b.squares)=%d", addr, x, y, idx, len(b.squares))
	return idx
}
