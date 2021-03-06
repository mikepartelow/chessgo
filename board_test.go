package chessgo_test

import (
	"fmt"
	"mp/chessgo"
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := chessgo.NewBoard()

	files := []byte("abcdefgh")
	ranks := []byte("12345678")

	for _, file := range files {
		for _, rank := range ranks {
			addr := chessgo.NewAddress(byte(file), byte(rank))
			want := defaultPiece(rank, file)
			assertSquare(t, b, addr, want)
		}
	}
}

func TestInBounds(t *testing.T) {
	b := chessgo.NewBoard()

	testCases := []struct {
		addr chessgo.Address
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
			addr: chessgo.NewAddress(uint8('a')-1, 1),
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
		addr chessgo.Address
		want chessgo.Piece
	}{
		{
			addr: "b8",
			want: chessgo.BlackKnight{},
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
		addr := chessgo.NewAddress('d', '5')
		want := chessgo.BlackKnight{}
		b.SetSquare(addr, want)

		assertSquare(t, b, addr, want)
	})

	t.Run("panics on out of bounds(a)", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Panic")
			}
		}()
		b.SetSquare("-1", chessgo.BlackBishop{})
	})

	t.Run("panics on out of bounds(b)", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Panic")
			}
		}()
		b.SetSquare("a99", chessgo.BlackBishop{})
	})
}

func TestBoardMove(t *testing.T) {
	testCases := []struct {
		srcAddr  chessgo.Address
		dstAddr  chessgo.Address
		replaced chessgo.Piece
	}{
		{
			srcAddr: "a2",
			dstAddr: "a3",
		},
		{
			srcAddr: "b7",
			dstAddr: "b5",
		},
		{
			srcAddr:  "c7",
			dstAddr:  "c8",
			replaced: chessgo.BlackBishop{},
		},
	}

	for _, tC := range testCases {
		b := chessgo.NewBoard()

		desc := fmt.Sprintf("%s to %s", tC.srcAddr, tC.dstAddr)
		t.Run(desc, func(t *testing.T) {
			want := b.GetSquare(tC.srcAddr)
			replaced := b.Move(tC.srcAddr, tC.dstAddr)
			assertSquare(t, b, tC.dstAddr, want)
			assertSquare(t, b, tC.srcAddr, nil)
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

func TestBoardString(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		board := chessgo.NewBoard()
		want := "RNBQKBNRPPPPPPPP                                pppppppprnbqkbnr"
		got := board.String()

		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	})

	t.Run("after a move", func(t *testing.T) {
		board := chessgo.NewBoard()
		board.Move("e2", "e4")
		want := "RNBQKBNRPPPP PPP            P                   pppppppprnbqkbnr"
		got := board.String()

		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	})
}

func TestCheck(t *testing.T) {
	testCases := []struct {
		boardString string
		color       chessgo.Color
		want        bool
	}{
		{
			boardString: "RNBQKBNRPPPPPPPP                                pppppppprnbqkbnr",
			color:       chessgo.Black,
			want:        false,
		},
		{
			boardString: "RNBQKBNRPPPPPPPP                                pppppppprnbqkQnr",
			color:       chessgo.Black,
			want:        true,
		},
		{
			boardString: "RNBQKBNRPPPPPPPP                                pppppQpprnbqkbnr",
			color:       chessgo.Black,
			want:        true,
		},
		{
			boardString: "q     b P P K      P Q        P  p NP  Pn  B    p  p pNpr b k nr",
			color:       chessgo.Black,
			want:        true,
		},
		{
			boardString: "q     b P P K      P          P  p NP  Pn    n  p  pBpNpr bk   r",
			color:       chessgo.Black,
			want:        true,
		},
	}

	for idx, tC := range testCases {
		desc := fmt.Sprintf("Case %d want %v", idx, tC.want)
		t.Run(desc, func(t *testing.T) {
			board := chessgo.NewBoardFromString(tC.boardString)
			got := board.InCheck(tC.color)
			if got != tC.want {
				t.Errorf("got %v wanted %v", got, tC.want)
			}
		})
	}
}

func assertSquare(t testing.TB, board chessgo.Board, addr chessgo.Address, want chessgo.Piece) {
	t.Helper()
	got := board.GetSquare(addr)
	if got != want {
		t.Errorf("got %q wanted %q at %q", got, want, addr)
	}
}

func defaultPiece(rank, file byte) chessgo.Piece {
	switch rank {
	case '1':
		switch file {
		case 'a', 'h':
			return chessgo.WhiteRook{}
		case 'b', 'g':
			return chessgo.WhiteKnight{}
		case 'c', 'f':
			return chessgo.WhiteBishop{}
		case 'd':
			return chessgo.WhiteQueen{}
		case 'e':
			return chessgo.WhiteKing{}
		}
	case '8':
		switch file {
		case 'a', 'h':
			return chessgo.BlackRook{}
		case 'b', 'g':
			return chessgo.BlackKnight{}
		case 'c', 'f':
			return chessgo.BlackBishop{}
		case 'd':
			return chessgo.BlackQueen{}
		case 'e':
			return chessgo.BlackKing{}
		}
	case '2':
		return chessgo.WhitePawn{}
	case '7':
		return chessgo.BlackPawn{}
	}

	return nil
}
