package chess_test

import (
	"fmt"
	chess "mp/chessgo"
	"testing"
)

func TestBoard(t *testing.T) {
	b := &chess.Board{}

	t.Run("ranks", func(t *testing.T) {
		wantRanks := uint8(8)
		gotRanks := b.Ranks()

		if gotRanks != wantRanks {
			t.Errorf("got %v wanted %v", gotRanks, wantRanks)
		}
	})

	t.Run("files", func(t *testing.T) {
		wantFiles := uint8(8)
		gotFiles := b.Files()

		if gotFiles != wantFiles {
			t.Errorf("got %v wanted %v", gotFiles, wantFiles)
		}
	})

	t.Run("bounds", func(t *testing.T) {
		testCases := []struct {
			square chess.Square
			want   bool
		}{
			{
				chess.Square{0, 0},
				true,
			},
			{
				chess.Square{0, 8},
				false,
			},
			{
				chess.Square{8, 8},
				false,
			},
			{
				chess.Square{8, 0},
				false,
			},
			{
				chess.Square{2, 2},
				true,
			},
			{
				chess.Square{7, 7},
				true,
			},
		}
		for _, tC := range testCases {
			t.Run(fmt.Sprintf("InBounds(%v)", tC.square), func(t *testing.T) {
				got := b.InBounds(tC.square)
				if got != tC.want {
					t.Errorf("got %v want %v", got, tC.want)
				}
			})
		}
	})
}

func TestNewBoard(t *testing.T) {
	b := chess.NewBoard()

	t.Run("white pieces", func(t *testing.T) {
		testCases := []struct {
			square chess.Square
			want   rune
		}{
			{
				square: chess.Square{0, 7},
				want:   'R',
			},
			{
				square: chess.Square{1, 7},
				want:   'N',
			},
			{
				square: chess.Square{2, 7},
				want:   'B',
			},
		}
		for _, tC := range testCases {
			t.Run(fmt.Sprintf("%c at %v", tC.want, tC.square), func(t *testing.T) {
				got := b.GetPieceAt(tC.square)
				if got != tC.want {
					t.Errorf("got %q at %v, wanted %q", got, tC.square, tC.want)
				}
			})
		}
	})

	// t.Run("black pieces", func(t *testing.T) {

	// }

	// t.Run("empties", func(t *testing.T) {

	// }
}

func TestGetPieceAt(t *testing.T) {
	b := chess.NewBoard()

	var want = ' '
	got := b.GetPieceAt(chess.Square{0, 0})

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
