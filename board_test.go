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

	want := 'n'
	got := b.GetSquare("b8")

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}

	// todo: assert a0, h9, i1, ("%c%c", 'a'-1, 10) panics
}

func TestSetSquare(t *testing.T) {
	b := chessgo.NewBoard()

	addr := "d5"
	want := 'n'
	b.SetSquare(addr, want)

	assertSquare(t, b, addr, want)

	// todo: assert invalid piece panics
	// todo: assert a0, h9, i1, ("%c%c", 'a'-1, 10) panics

}

func TestMove(t *testing.T) {
	testCases := []struct {
		fromAddr string
		toAddr   string
		replaced rune
	}{
		{
			fromAddr: "a2",
			toAddr:   "a3",
			replaced: chessgo.EmptySquare,
		},
		{
			fromAddr: "b7",
			toAddr:   "b5",
			replaced: chessgo.EmptySquare,
		},
		{
			fromAddr: "c7",
			toAddr:   "c8",
			replaced: 'b',
		},
	}

	for _, tC := range testCases {
		b := chessgo.NewBoard()

		desc := fmt.Sprintf("%s to %s", tC.fromAddr, tC.toAddr)
		t.Run(desc, func(t *testing.T) {
			want := b.GetSquare(tC.fromAddr)
			replaced := b.Move(tC.fromAddr, tC.toAddr)
			assertSquare(t, b, tC.toAddr, want)
			assertSquare(t, b, tC.fromAddr, chessgo.EmptySquare)
			if replaced != tC.replaced {
				t.Errorf("wanted to replace %c, replaced %c instead", tC.replaced, replaced)
			}
		})
	}

	// todo: inalid fromAddr panic
	// todo: inalid toAddr panic
}

func assertSquare(t testing.TB, board *chessgo.Board, addr string, want rune) {
	t.Helper()
	got := board.GetSquare(addr)
	if got != want {
		t.Errorf("got %q wanted %q at %q", got, want, addr)
	}
}

func defaultPiece(rank, file rune) rune {
	switch rank {
	case '1':
		switch file {
		case 'a', 'h':
			return 'R'
		case 'b', 'g':
			return 'N'
		case 'c', 'f':
			return 'B'
		case 'd':
			return 'Q'
		case 'e':
			return 'K'
		}
	case '8':
		switch file {
		case 'a', 'h':
			return 'r'
		case 'b', 'g':
			return 'n'
		case 'c', 'f':
			return 'b'
		case 'd':
			return 'q'
		case 'e':
			return 'k'
		}
	case '2':
		return 'P'
	case '7':
		return 'p'
	}

	return chessgo.EmptySquare
}
