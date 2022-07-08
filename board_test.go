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

func TestInBounds(t *testing.T) {
	b := chessgo.NewBoard()

	testCases := []struct {
		addr string
		want bool
	}{
		{
			addr: "b8",
			want: true,
		},
		{
			addr: "b9",
			want: false,
		},
		{
			addr: "a99",
			want: false,
		},
		{
			addr: "9",
			want: false,
		},
		{
			addr: "a",
			want: false,
		},
		{
			addr: "",
			want: false,
		},
		{
			addr: fmt.Sprintf("%c1", rune(uint8('a')-1)),
			want: false,
		},
	}
	for _, tC := range testCases {
		desc := fmt.Sprintf("%v at %s", tC.want, tC.addr)
		t.Run(desc, func(t *testing.T) {
			got := b.InBounds(tC.addr)

			if got != tC.want {
				t.Errorf("got %v wanted %v for %q", got, tC.want, tC.addr)
			}
		})
	}
}

func TestGetSquare(t *testing.T) {
	b := chessgo.NewBoard()

	testCases := []struct {
		addr string
		want chessgo.Piece
	}{
		{
			addr: "b8",
			want: chessgo.BlackKnight,
		},
	}
	for _, tC := range testCases {
		desc := fmt.Sprintf("%c at %s", tC.want, tC.addr)
		t.Run(desc, func(t *testing.T) {
			got := b.GetSquare(tC.addr)

			if got != tC.want {
				t.Errorf("got %q wanted %q", got, tC.want)
			}
		})
	}

	t.Run("panics on out of bounds", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Panic")
			}
		}()
		b.GetSquare("h9")
	})
}

func TestSetSquare(t *testing.T) {
	b := chessgo.NewBoard()

	t.Run("sets square contents", func(t *testing.T) {
		addr := "d5"
		want := chessgo.BlackKnight
		b.SetSquare(addr, want)

		assertSquare(t, b, addr, want)
	})

	t.Run("panics on out of bounds(a)", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Panic")
			}
		}()
		b.SetSquare("-1", chessgo.BlackBishop)
	})

	t.Run("panics on out of bounds(b)", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Panic")
			}
		}()
		b.SetSquare("a99", chessgo.BlackBishop)
	})

	t.Run("panics on invalid piece", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Panic")
			}
		}()
		b.SetSquare("a1", 'X')
	})
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

	t.Run("panics on out of bounds(a)", func(t *testing.T) {
		b := chessgo.NewBoard()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Panic")
			}
		}()
		b.Move("a1", "a99")
	})

	t.Run("panics on out of bounds(b)", func(t *testing.T) {
		b := chessgo.NewBoard()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Panic")
			}
		}()
		b.Move("z99", "a1")
	})
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
