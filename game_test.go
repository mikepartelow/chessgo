package chessgo_test

import (
	"fmt"
	"math"
	"mp/chessgo"
	"strings"
	"testing"
)

type MockBoard struct {
	squares []byte
}

func TestGameMove(t *testing.T) {
	testCases := []struct {
		board        MockBoard
		turn         chessgo.Color
		move         string
		wantBoard    MockBoard
		wantCaptured chessgo.Piece
		wantErr      string
	}{
		{
			board:     MockBoard{squares: []byte("P   ")},
			move:      "a2",
			wantBoard: MockBoard{squares: []byte("  P ")},
		},
		{
			board:     MockBoard{squares: []byte(" P  ")},
			move:      "b2",
			wantBoard: MockBoard{squares: []byte("   P")},
		},
		{
			board:     MockBoard{squares: []byte("B   ")},
			move:      "Bb2",
			wantBoard: MockBoard{squares: []byte("   B")},
		},
		{
			board:     MockBoard{squares: []byte("b   ")},
			turn:      chessgo.Black,
			move:      "Bb2",
			wantBoard: MockBoard{squares: []byte("   b")},
		},
		{
			board:     MockBoard{squares: []byte("  p ")},
			turn:      chessgo.Black,
			move:      "a1",
			wantBoard: MockBoard{squares: []byte("p   ")},
		},
		{
			board:        MockBoard{squares: []byte("b  N")},
			turn:         chessgo.Black,
			move:         "Bxb2",
			wantBoard:    MockBoard{squares: []byte("   b")},
			wantCaptured: chessgo.WhiteKnight{},
		},
		{
			board:     MockBoard{squares: []byte("b  n")},
			turn:      chessgo.Black,
			move:      "Bxb2",
			wantBoard: MockBoard{squares: []byte("b  n")},
			wantErr:   "attempt to capture own piece",
		},
		{
			board:     MockBoard{squares: []byte("    P           ")},
			move:      "a4",
			wantBoard: MockBoard{squares: []byte("            P   ")},
		},
		{
			board:     MockBoard{squares: []byte("        p       ")},
			turn:      chessgo.Black,
			move:      "a1",
			wantBoard: MockBoard{squares: []byte("p               ")},
		},
		{
			board:        MockBoard{squares: []byte("P  q")},
			move:         "axb2",
			wantBoard:    MockBoard{squares: []byte("   P")},
			wantCaptured: chessgo.BlackQueen{},
		},
		{
			board:        MockBoard{squares: []byte(" Pq ")},
			move:         "bxa2",
			wantBoard:    MockBoard{squares: []byte("  P ")},
			wantCaptured: chessgo.BlackQueen{},
		},
		// todo: disambiguate the move (left diagonal capture) between two possible pawns
		// todo: disambiguate the move (right diagonal capture) between two possible pawns
		{
			board:        MockBoard{squares: []byte("Q  p")},
			turn:         chessgo.Black,
			move:         "bxa1",
			wantBoard:    MockBoard{squares: []byte("p   ")},
			wantCaptured: chessgo.WhiteQueen{},
		},
		{
			board:        MockBoard{squares: []byte(" Qp ")},
			turn:         chessgo.Black,
			move:         "axb1",
			wantBoard:    MockBoard{squares: []byte(" p  ")},
			wantCaptured: chessgo.WhiteQueen{},
		},
		// todo: disambiguate the move (left diagonal capture) between two possible pawns
		// todo: disambiguate the move (right diagonal capture) between two possible pawns
		{
			board:     MockBoard{squares: []byte("B               ")},
			move:      "Bd4",
			wantBoard: MockBoard{squares: []byte("               B")},
		},
		{
			board:     MockBoard{squares: []byte("               B")},
			move:      "Ba1",
			wantBoard: MockBoard{squares: []byte("B               ")},
		},
		{
			board:     MockBoard{squares: []byte("   B            ")},
			move:      "Ba4",
			wantBoard: MockBoard{squares: []byte("            B   ")},
		},
		{
			board:     MockBoard{squares: []byte("            B   ")},
			move:      "Bd1",
			wantBoard: MockBoard{squares: []byte("   B            ")},
		},
		{
			board:     MockBoard{squares: []byte("Q      k ")},
			move:      "Qc3+",
			wantBoard: MockBoard{squares: []byte("       kQ")},
		},
		{
			board:     MockBoard{squares: []byte("K   ")},
			move:      "Kb1",
			wantBoard: MockBoard{squares: []byte(" K  ")},
		},
		{
			board:     MockBoard{squares: []byte("   k")},
			turn:      chessgo.Black,
			move:      "Kb1",
			wantBoard: MockBoard{squares: []byte(" k  ")},
		},
		{
			board:     MockBoard{squares: []byte("K   ")},
			move:      "Kb2",
			wantBoard: MockBoard{squares: []byte("   K")},
		},
		{
			board:     MockBoard{squares: []byte("N        ")},
			move:      "Nb3",
			wantBoard: MockBoard{squares: []byte("       N ")},
		},
		{
			board:     MockBoard{squares: []byte("n        ")},
			turn:      chessgo.Black,
			move:      "Nc2",
			wantBoard: MockBoard{squares: []byte("     n   ")},
		},
		{
			board:     MockBoard{squares: []byte("       N ")},
			move:      "Na1",
			wantBoard: MockBoard{squares: []byte("N        ")},
		},
		{
			board:     MockBoard{squares: []byte("     n   ")},
			turn:      chessgo.Black,
			move:      "Na1",
			wantBoard: MockBoard{squares: []byte("n        ")},
		},

		{
			board:     MockBoard{squares: []byte("Q        ")},
			move:      "Qa3",
			wantBoard: MockBoard{squares: []byte("      Q  ")},
		},
		{
			board:     MockBoard{squares: []byte("   q     ")},
			turn:      chessgo.Black,
			move:      "Qc2",
			wantBoard: MockBoard{squares: []byte("     q   ")},
		},
		{
			board:     MockBoard{squares: []byte("R        ")},
			move:      "Ra3",
			wantBoard: MockBoard{squares: []byte("      R  ")},
		},
		{
			board:     MockBoard{squares: []byte("RP R             ")},
			move:      "Rc1",
			wantBoard: MockBoard{squares: []byte("RPR              ")},
		},
		{
			board:     MockBoard{squares: []byte("RP R             ")},
			move:      "Rc1",
			wantBoard: MockBoard{squares: []byte("RPR              ")},
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

			assertBoard(t, g.Board.String(), tC.wantBoard.String())
			assertCaptured(t, captured, tC.wantCaptured)
		})
	}

	t.Run("turn toggles on valid move", func(t *testing.T) {
		g := chessgo.Game{Board: &MockBoard{squares: []byte("P   ")}}
		if g.Turn != chessgo.White {
			t.Errorf("expected Turn = White")
		}
		_, err := g.Move("a2")

		assertNoError(t, err)

		if g.Turn != chessgo.Black {
			t.Errorf("expected Turn = Black")
		}
	})

	// todo:
	// t.Run("turn does not toggle on invalid move")

}

func TestGameCastle(t *testing.T) {
	t.Run("king side", func(t *testing.T) {
		board := &MockBoard{squares: []byte("R BQKB RPP  PPPP  N  N    PP                 np ppppppbprnbqk  r")}
		wantBoard := "R BQKB RPP  PPPP  N  N    PP                 np ppppppbprnbq rk "
		game := chessgo.Game{Board: board, Turn: chessgo.Black}

		captured, err := game.Move("O-O")
		assertNoError(t, err)
		assertCaptured(t, captured, nil)
		assertBoard(t, board.String(), wantBoard)
	})

	// todo:
	// t.Run("queen side")
	// t.Run("moved rook, can't castle")
	// t.Run("moved king, can't castle")
	// t.Run("can't castle through check")
}

func (b *MockBoard) InBounds(addr chessgo.Address) bool {
	file := addr[0]
	rank := addr[1]

	r := file >= 'a' && rank >= '1' && file <= byte(b.MaxFile()) && rank <= byte(b.MaxRank())
	return r
}

func (b *MockBoard) GetSquare(addr chessgo.Address) chessgo.Piece {
	return chessgo.PieceFromByte(b.squares[b.getIndex(addr)])
}

func (b *MockBoard) SetSquare(addr chessgo.Address, piece chessgo.Piece) {
	if piece != nil {
		b.squares[b.getIndex(addr)] = piece.Byte()
	} else {
		b.squares[b.getIndex(addr)] = ' '
	}
}

func (b *MockBoard) Move(srcAddr, dstAddr chessgo.Address) chessgo.Piece {
	captured := b.GetSquare(dstAddr)
	b.SetSquare(dstAddr, b.GetSquare(srcAddr))
	b.SetSquare(srcAddr, nil)
	return captured
}

func (b *MockBoard) String() string {
	return string(b.squares)
}

func (b *MockBoard) MaxFile() byte {
	m := byte(uint8('a')+uint8(math.Sqrt(float64(len(b.squares))))) - 1
	return m
}

func (b *MockBoard) MaxRank() byte {
	m := byte(uint8('0') + uint8(math.Sqrt(float64(len(b.squares)))))
	return m
}

func (b *MockBoard) InCheck(color chessgo.Color) bool {
	return false
}

func (b *MockBoard) Mated(color chessgo.Color) bool {
	return false
}

func (b *MockBoard) getIndex(addr chessgo.Address) uint8 {
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
