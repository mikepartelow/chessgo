package chessgo_test

import (
	"fmt"
	"math"
	"mp/chessgo"
	"reflect"
	"strings"
	"testing"
)

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
		wantErr      string
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
			wantCaptured: chessgo.WhiteKnight{},
		},
		{
			board:     StubBoard{squares: []byte("b  n")},
			turn:      chessgo.Black,
			move:      "Bxb2",
			wantBoard: StubBoard{squares: []byte("b  n")},
			wantErr:   "attempt to capture own piece",
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
			wantCaptured: chessgo.BlackQueen{},
		},
		{
			board:        StubBoard{squares: []byte(" Pq ")},
			move:         "bxa2",
			wantBoard:    StubBoard{squares: []byte("  P ")},
			wantCaptured: chessgo.BlackQueen{},
		},
		// todo: disambiguate the move (left diagonal capture) between two possible pawns
		// todo: disambiguate the move (right diagonal capture) between two possible pawns
		{
			board:        StubBoard{squares: []byte("Q  p")},
			turn:         chessgo.Black,
			move:         "bxa1",
			wantBoard:    StubBoard{squares: []byte("p   ")},
			wantCaptured: chessgo.WhiteQueen{},
		},
		{
			board:        StubBoard{squares: []byte(" Qp ")},
			turn:         chessgo.Black,
			move:         "axb1",
			wantBoard:    StubBoard{squares: []byte(" p  ")},
			wantCaptured: chessgo.WhiteQueen{},
		},
		// todo: disambiguate the move (left diagonal capture) between two possible pawns
		// todo: disambiguate the move (right diagonal capture) between two possible pawns
		{
			board:     StubBoard{squares: []byte("B               ")},
			move:      "Bd4",
			wantBoard: StubBoard{squares: []byte("               B")},
		},
		{
			board:     StubBoard{squares: []byte("               B")},
			move:      "Ba1",
			wantBoard: StubBoard{squares: []byte("B               ")},
		},
		{
			board:     StubBoard{squares: []byte("   B            ")},
			move:      "Ba4",
			wantBoard: StubBoard{squares: []byte("            B   ")},
		},
		{
			board:     StubBoard{squares: []byte("            B   ")},
			move:      "Bd1",
			wantBoard: StubBoard{squares: []byte("   B            ")},
		},
		{
			board:     StubBoard{squares: []byte("Q      k ")},
			move:      "Qc3+",
			wantBoard: StubBoard{squares: []byte("       kQ")},
		},
		{
			board:     StubBoard{squares: []byte("K   ")},
			move:      "Kb1",
			wantBoard: StubBoard{squares: []byte(" K  ")},
		},
		{
			board:     StubBoard{squares: []byte("   k")},
			turn:      chessgo.Black,
			move:      "Kb1",
			wantBoard: StubBoard{squares: []byte(" k  ")},
		},
		{
			board:     StubBoard{squares: []byte("K   ")},
			move:      "Kb2",
			wantBoard: StubBoard{squares: []byte("   K")},
		},
		{
			board:     StubBoard{squares: []byte("N        ")},
			move:      "Nb3",
			wantBoard: StubBoard{squares: []byte("       N ")},
		},
		{
			board:     StubBoard{squares: []byte("n        ")},
			turn:      chessgo.Black,
			move:      "Nc2",
			wantBoard: StubBoard{squares: []byte("     n   ")},
		},
		{
			board:     StubBoard{squares: []byte("       N ")},
			move:      "Na1",
			wantBoard: StubBoard{squares: []byte("N        ")},
		},
		{
			board:     StubBoard{squares: []byte("     n   ")},
			turn:      chessgo.Black,
			move:      "Na1",
			wantBoard: StubBoard{squares: []byte("n        ")},
		},

		{
			board:     StubBoard{squares: []byte("Q        ")},
			move:      "Qa3",
			wantBoard: StubBoard{squares: []byte("      Q  ")},
		},
		{
			board:     StubBoard{squares: []byte("   q     ")},
			turn:      chessgo.Black,
			move:      "Qc2",
			wantBoard: StubBoard{squares: []byte("     q   ")},
		},
		{
			board:     StubBoard{squares: []byte("R        ")},
			move:      "Ra3",
			wantBoard: StubBoard{squares: []byte("      R  ")},
		},
		{
			board:     StubBoard{squares: []byte("RP R             ")},
			move:      "Rc1",
			wantBoard: StubBoard{squares: []byte("RPR              ")},
		},
	}

	for _, tC := range testCases {
		desc := fmt.Sprintf("move %s / %s", tC.move, tC.turn.String())
		t.Run(desc, func(t *testing.T) {
			g := chessgo.Game{Board: &tC.board, Turn: tC.turn}
			captured, err := g.Move(tC.move)

			if err != nil && !strings.Contains(err.Error(), tC.wantErr) {
				t.Errorf("got err %T/%v wanted %v", err, err, tC.wantErr)
			}

			if !reflect.DeepEqual(g.Board, &tC.wantBoard) {
				t.Errorf("got %q, wanted %q after move %q", g.Board, &tC.wantBoard, tC.move)
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

func (b *StubBoard) InBounds(addr chessgo.Address) bool {
	file := addr[0]
	rank := addr[1]

	r := file >= 'a' && rank >= '1' && file <= byte(b.MaxFile()) && rank <= byte(b.MaxRank())
	return r
}

func (b *StubBoard) GetSquare(addr chessgo.Address) chessgo.Piece {
	return chessgo.PieceFromByte(b.squares[b.getIndex(addr)])
}

func (b *StubBoard) SetSquare(addr chessgo.Address, piece chessgo.Piece) {
	if piece != nil {
		b.squares[b.getIndex(addr)] = piece.Byte()
	} else {
		b.squares[b.getIndex(addr)] = ' '
	}
}

func (b *StubBoard) Move(srcAddr, dstAddr chessgo.Address) chessgo.Piece {
	captured := b.GetSquare(dstAddr)
	b.SetSquare(dstAddr, b.GetSquare(srcAddr))
	b.SetSquare(srcAddr, nil)
	return captured
}

func (b *StubBoard) String() string {
	return string(b.squares)
}

func (b *StubBoard) MaxFile() byte {
	m := byte(uint8('a')+uint8(math.Sqrt(float64(len(b.squares))))) - 1
	return m
}

func (b *StubBoard) MaxRank() byte {
	m := byte(uint8('0') + uint8(math.Sqrt(float64(len(b.squares)))))
	return m
}

func (b *StubBoard) InCheck(color chessgo.Color) bool {
	return false
}

func (b *StubBoard) Mated(color chessgo.Color) bool {
	return false
}

func (b *StubBoard) getIndex(addr chessgo.Address) uint8 {
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
	if idx >= uint8(len(b.squares)) {
		panic("idx >= len(b.squares)")
	}
	return idx
}
