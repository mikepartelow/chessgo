package chessgo_test

import (
	"fmt"
	"mp/chessgo"
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := chessgo.NewBoard()

	files := "abcdefgh"
	ranks := "12345678"

	for _, file := range files {
		for _, rank := range ranks {
			addr := fmt.Sprintf("%c%c", file, rank)
			want := defaultPiece(rank, file)
			assertSquare(t, b, addr, want)
		}
	}
}

func TestGetSquare(t *testing.T) {
	b := chessgo.NewBoard()

	want := chessgo.BlackKnight
	got := b.GetSquare("b8")

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}

	// todo: assert a0, h9, i1, ("%c%c", 'a'-1, 10) panics
}

func TestSetSquare(t *testing.T) {
	b := chessgo.NewBoard()

	addr := "d5"
	want := chessgo.BlackKnight
	b.SetSquare(addr, want)

	assertSquare(t, b, addr, want)

	// todo: assert invalid piece panics
	// todo: assert a0, h9, i1, ("%c%c", 'a'-1, 10) panics

}

func TestMove(t *testing.T) {
	testCases := []struct {
		fromAddr string
		toAddr   string
		replaced chessgo.Piece
	}{
		{
			fromAddr: "a2",
			toAddr:   "a3",
			replaced: chessgo.NoPiece,
		},
		{
			fromAddr: "b7",
			toAddr:   "b5",
			replaced: chessgo.NoPiece,
		},
		{
			fromAddr: "c7",
			toAddr:   "c8",
			replaced: chessgo.BlackBishop,
		},
	}

	for _, tC := range testCases {
		b := chessgo.NewBoard()

		desc := fmt.Sprintf("%s to %s", tC.fromAddr, tC.toAddr)
		t.Run(desc, func(t *testing.T) {
			want := b.GetSquare(tC.fromAddr)
			replaced := b.Move(tC.fromAddr, tC.toAddr)
			assertSquare(t, b, tC.toAddr, want)
			assertSquare(t, b, tC.fromAddr, chessgo.NoPiece)
			if replaced != tC.replaced {
				t.Errorf("wanted to replace %c, replaced %c instead", tC.replaced, replaced)
			}
		})
	}

	// todo: inalid fromAddr panic
	// todo: inalid toAddr panic
}

func assertSquare(t testing.TB, board *chessgo.Board, addr string, want chessgo.Piece) {
	t.Helper()
	got := board.GetSquare(addr)
	if got != want {
		t.Errorf("got %q wanted %q at %q", got, want, addr)
	}
}

func defaultPiece(rank, file rune) chessgo.Piece {
	switch rank {
	case '1':
		switch file {
		case 'a', 'h':
			return chessgo.WhiteRook
		case 'b', 'g':
			return chessgo.WhiteKnight
		case 'c', 'f':
			return chessgo.WhiteBishop
		case 'd':
			return chessgo.WhiteQueen
		case 'e':
			return chessgo.WhiteKing
		}
	case '8':
		switch file {
		case 'a', 'h':
			return chessgo.BlackRook
		case 'b', 'g':
			return chessgo.BlackKnight
		case 'c', 'f':
			return chessgo.BlackBishop
		case 'd':
			return chessgo.BlackQueen
		case 'e':
			return chessgo.BlackKing
		}
	case '2':
		return chessgo.WhitePawn
	case '7':
		return chessgo.BlackPawn
	}

	return chessgo.NoPiece
}
